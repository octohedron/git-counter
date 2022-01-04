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
	totalCols     = 60
	padding       = 10
	colsNoPadding = totalCols - (padding * 2)
)

var (
	flagDirectories allDirectories
	author          string
)

type allDirectories []string

type directory struct {
	gitCommand    string
	dayilyCommits map[int]int
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
		gitCommand: fmt.Sprintf("git -C %s log ", path) +
			author + ` --all --format='%ad' --date='format:%H'`,
	}
}

func showResults(results map[int]int) {
	var keys []int
	for k := range results {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	maxCommits := 0
	totalCommits := 0
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
	projects := 0
	c := make(chan map[int]int)
	// For each directory
	for _, v := range flagDirectories {
		gitFolders, err := ioutil.ReadDir(v)
		logPanic(err)
		for _, f := range gitFolders {
			dir := getDir(v + "/" + f.Name())
			go dir.parseDir(c)
			projects++
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
