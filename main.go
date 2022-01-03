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

type dirT []string

const usage = "=====================\nUSAGE\n" +
	"# Allows many paths with the -dir flag\n" +
	"> $ go run main.go -dir=/full/path1... -dir=/full/path2... " +
	"-dir=/full/pathN...\n" +
	"MORE EXAMPLES\n" +
	"# With single path\n" +
	"> $ go run main.go -dir=/home/user/go/src/github.com/user\n" +
	"# With author\n" +
	"> $ go run main.go -dir=/home/user/go/src/github.com/user -author='User.*'"

var (
	tots   map[int]int
	dirS   dirT
	author string
)

func (i *dirT) String() string {
	return "A string"
}

func (i *dirT) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func getMap() map[int]int {
	t := make(map[int]int)
	for i := 0; i < 24; i++ {
		t[i] = 0
	}
	return t
}

func init() {
	tots = getMap()
	flag.Var(&dirS, "dir", "directories")
	flag.StringVar(&author, "author", "", "a string")
	flag.Parse()
	if len(dirS) < 1 {
		fmt.Println(usage)
		os.Exit(0)
	}
	// add author to the git command if present
	if author != "" {
		author = "--author='" + author + "'"
	}
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func logPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func addFolderCommits(outs *bytes.Buffer) map[int]int {
	scanner := bufio.NewScanner(outs)
	r := getMap()
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		logPanic(err)
		r[v]++
	}
	return r
}

func printOut(ti map[int]int) {
	var keys []int
	for k := range ti {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	max := float64(0)
	total := 0
	for _, k := range keys {
		if float64(ti[k]) > max {
			max = float64(ti[k])
		}
		total += ti[k]
	}
	fmt.Println("MAX", max, "TOTAL", total)
	for _, k := range keys {
		line := fmt.Sprintf("%3v %7v ", k, ti[k])
		for n := 0; float64(n) < math.Abs(float64(ti[k])/float64(max)*80); n++ {
			line += "*"
		}
		fmt.Println(line)
	}
}

func checkDir(dir string, c chan map[int]int) {
	cmd := exec.Command("sh", "-c", "git --git-dir="+dir+"/.git log "+author+" --format='%ad' --date='format:%H'")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	printCommand(cmd)
	err := cmd.Run()
	printError(err)
	t := addFolderCommits(cmdOutput)
	c <- t
}

func main() {
	start := time.Now()
	// Folders that will be added
	folders := make(map[string][]string)
	projects := 0
	// For each directory
	for _, v := range dirS {
		gitFolders, err := ioutil.ReadDir(v)
		logPanic(err)
		for _, f := range gitFolders {
			// Don't add this project
			if f.Name() == "git-counter" {
				continue
			}
			folders[v] = append(folders[v], v+"/"+f.Name())
			projects++
		}
	}
	c := make(chan map[int]int)
	for _, folder := range folders {
		for _, dir := range folder {
			go checkDir(dir, c)
		}
	}
	completed := 0
	for t := range c {
		for v, k := range t {
			tots[v] += k
		}
		completed++
		if completed == projects {
			break
		}
	}
	printOut(tots)
	log.Printf("%s", time.Since(start))
}
