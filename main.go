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
	"time"

	h "github.com/dustin/go-humanize"
)

const (
	// These are used for printing the instructions
	totalCols     = 60
	padding       = 10
	colsNoPadding = totalCols - (padding * 2)
)

var (
	// Used for parsing the directory flags
	flagDirectories allDirectories
	// Author is a global variable to avoid passing it to the
	// getDir function
	author string
)

// Used for parsing the directory flags
type allDirectories []string

type directory struct {
	gitCommand    string
	hourlyCommits map[int]int
	path          string
}

func init() {
	flag.Var(&flagDirectories, "dir", "flagDirectories")
	flag.StringVar(&author, "author", "", "a string")
	flag.Parse()
	if len(flagDirectories) < 1 {
		printUsage()
		os.Exit(0)
	}
	// Add author to the git command if present
	if author != "" {
		author = fmt.Sprintf("--author='%s'", author)
	}
}

// addDirectoryCommits adds the commits of each github project
// to the directory
func (d *directory) addDirectoryCommits(outs *bytes.Buffer) {
	scanner := bufio.NewScanner(outs)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		logPanic(err)
		d.hourlyCommits[v]++
	}
}

// getDir initializes a directory and returns a pointer
func getDir(path string) *directory {
	return &directory{
		hourlyCommits: get24HourMap(),
		gitCommand: fmt.Sprintf("git -C %s log ", path) +
			author + ` --all --format='%ad' --date='format:%H'`,
		path: path,
	}
}

// This will print the graph in the terminal after collecting the commits
func showResults(results map[int]int) {
	// For showing the results starting at 0 to 23h
	var keys []int
	for k := range results {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	maxCommits, totalCommits := 0, 0
	for _, k := range keys {
		if results[k] > maxCommits {
			maxCommits = results[k]
		}
		totalCommits += results[k]
	}
	fmt.Println("MAX", h.Comma(int64(maxCommits)),
		"TOTAL", h.Comma(int64(totalCommits)))
	for _, k := range keys {
		line := fmt.Sprintf("%3v %7v ", k, results[k])
		for n := 0; float64(n) < math.Abs(
			float64(results[k])/float64(maxCommits)*80); n++ {
			line += "*"
		}
		fmt.Println(line)
	}
}

// This will call the git command and parse the commit
func (d *directory) parseDir(c chan map[int]int) {
	cmd := exec.Command("sh", "-c", d.gitCommand)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	printCommand(cmd)
	err := cmd.Run()
	logError(err, d)
	d.addDirectoryCommits(cmdOutput)
	c <- d.hourlyCommits
}

func main() {
	start := time.Now()
	projects := 0
	c := make(chan map[int]int)
	// For each directory in the flags
	for _, v := range flagDirectories {
		projectFolders, err := ioutil.ReadDir(v)
		logPanic(err)
		for _, f := range projectFolders {
			repoPath := v + "/" + f.Name()
			// Check if it's a git project
			if _, err := os.Stat(repoPath + "/.git"); err == nil {
				dir := getDir(repoPath)
				// Launch goroutine to process each folder
				go dir.parseDir(c)
				projects++
			}
		}
	}
	// We ensure that each result gets processed one at a time
	// when receiving from the goroutine it will increment the total, then
	// continue with the next one
	total := get24HourMap()
	completed := 0
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
