package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	totalCols     = 60
	padding       = 10
	colsNoPadding = totalCols - (padding * 2)
)

type allDirectories []string

type directory struct {
	path          string
	gitCommand    string
	dayilyCommits map[int]int
}

func getPad() string {
	return strings.Repeat("#", padding)
}

func printDoc(contents string, center bool) {
	if center {
		contents = fmt.Sprintf(
			"%"+strconv.Itoa((padding*2)+len(contents)/2)+"s", contents)
	}
	fmt.Printf(
		"%10s%-"+strconv.Itoa(colsNoPadding)+"s%s\n", getPad(), contents, getPad())
}

func printUsage() {
	printDoc(strings.Repeat("#", colsNoPadding), false)
	printDoc("USAGE", true)
	fmt.Println(strings.Repeat("#", totalCols) + "\n")
	printDoc("Allows many paths with the -dir flag", true)
	fmt.Print("> $ go build && ./git-counter -dir=/full/path1... -dir=/full/path2... " +
		"-dir=/full/pathN...\n\n")
	printDoc(strings.Repeat("#", colsNoPadding), false)
	printDoc("EXAMPLES", true)
	printDoc(strings.Repeat("#", colsNoPadding), false)
	printDoc("With single path", true)
	fmt.Print("> $ go build && ./git-counter -dir=/home/user/go/src/github.com/user\n\n")
	printDoc("With autor", true)
	fmt.Print("> $ go build && ./git-counter -dir=/home/user/go/src/github.com/user" +
		" -author='User.*'\n\n")
}

var (
	flagDirectories allDirectories
	author          string
)

func (dirs *allDirectories) String() string {
	return "string"
}

func (dirs *allDirectories) Set(value string) error {
	*dirs = append(*dirs, value)
	return nil
}

// Initialize a map if time of day to amount of commits
func get24HourMap() map[int]int {
	t := make(map[int]int)
	for i := 0; i < 24; i++ {
		t[i] = 0
	}
	return t
}

func init() {
	flag.Var(&flagDirectories, "dir", "flagDirectories")
	flag.StringVar(&author, "author", "", "a string")
	flag.Parse()
	if len(flagDirectories) < 1 {
		printUsage()
		os.Exit(0)
	}
	// add author to the git command if present
	if author != "" {
		author = fmt.Sprintf("--author='%s'", author)
	}
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func logError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func logPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (d *directory) addDirectoryCommits(outs *bytes.Buffer) {
	scanner := bufio.NewScanner(outs)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		logPanic(err)
		d.dayilyCommits[v]++
	}
}

func getDir(path string) directory {
	return directory{
		dayilyCommits: get24HourMap(),
		path: fmt.Sprintf("git --git-dir=%s/.git log ", path) +
			author + ` --format='%ad' --date='format:%H'`,
	}
}

func showResults(results map[int]int) {
	var keys []int
	for k := range results {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	maxCommits := float64(0)
	totalCommits := 0
	for _, k := range keys {
		if float64(results[k]) > maxCommits {
			maxCommits = float64(results[k])
		}
		totalCommits += results[k]
	}
	fmt.Println("MAX", maxCommits, "TOTAL", totalCommits)
	for _, k := range keys {
		line := fmt.Sprintf("%3v %7v ", k, results[k])
		for n := 0; float64(n) < math.Abs(
			float64(results[k])/float64(maxCommits)*80); n++ {
			line += "*"
		}
		fmt.Println(line)
	}
}

func (d *directory) parseDir(c chan map[int]int) {
	cmd := exec.Command("sh", "-c", d.gitCommand)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	printCommand(cmd)
	err := cmd.Run()
	logError(err)
	d.addDirectoryCommits(cmdOutput)
	c <- d.dayilyCommits
}

func main() {
	start := time.Now()
	// Folders that will be added
	folders := make(map[string][]directory)
	projects := 0
	// For each directory
	for _, v := range flagDirectories {
		gitFolders, err := ioutil.ReadDir(v)
		logPanic(err)
		for _, f := range gitFolders {
			folders[v] = append(folders[v], getDir(v+"/"+f.Name()))
			projects++
		}
	}
	c := make(chan map[int]int)
	for _, directory := range folders {
		for _, dir := range directory {
			go dir.parseDir(c)
		}
	}
	completed := 0
	total := get24HourMap()
	for t := range c {
		for v, k := range t {
			total[v] += k
		}
		completed++
		if completed == projects {
			break
		}
	}
	showResults(total)
	log.Println(time.Since(start))
}
