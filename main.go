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
	// Used for printing the instructions
	totalCols     = 60
	padding       = 10
	colsNoPadding = totalCols - (padding * 2)
)

var (
	// Used for parsing the directory flags
	flagDirectories allDirectories
	author          string
	// Colors for the terminal
	colors = []string{
		"\033[0m",  // Reset
		"\033[90m", // DarkGray
		"\033[34m", // Blue
		"\033[37m", // LightGray
		"\033[94m", // LightBlue
		"\033[36m", // Cyan
		"\033[96m", // LightCyan
		"\033[92m", // LightGreen
		"\033[32m", // Green
		"\033[35m", // Magenta
		"\033[95m", // LightMagenta
		"\033[33m", // Yellow
		"\033[93m", // LightYellow
		"\033[91m", // LightRed
		"\033[31m", // Red
	}
)

// Used for parsing the directory flags
type allDirectories []string

type directory struct {
	gitCommand    string
	hourlyCommits map[int]int
	path          string
	maxCommits    int
	totalCommits  int
}

type gitCounter struct {
	directories  []directory
	totalCommits int
	maxCommits   int
	results      map[int]int
}

type counter interface {
	setResults()
	setMaxCommits()
	setTotalCommits()
	printResults()
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

func (d *directory) setMaxCommits() {
	maxCommits := 0
	for _, hourlyCommits := range d.hourlyCommits {
		if hourlyCommits > maxCommits {
			d.maxCommits = hourlyCommits
			maxCommits = hourlyCommits
		}
	}
}

func (d *directory) setTotalCommits() {
	for _, hourlyCommits := range d.hourlyCommits {
		d.totalCommits += hourlyCommits
	}
}

// getDir initializes a directory and returns a pointer to it
func getDir(path string) *directory {
	return &directory{
		hourlyCommits: get24HourMap(),
		gitCommand: fmt.Sprintf("git -C %s log ", path) +
			author + ` --all --format='%ad' --date='format:%H'`,
		path: path,
	}
}

func (c *gitCounter) setResults() {
	c.results = get24HourMap()
	for _, dir := range c.directories {
		for hour, hourlyCommits := range dir.hourlyCommits {
			if hourlyCommits > c.maxCommits {
				c.maxCommits = hourlyCommits
			}
			c.totalCommits += hourlyCommits
			c.results[hour] += hourlyCommits
		}
	}
}

func (c *gitCounter) setMaxCommits() {
	for _, d := range c.directories {
		if c.maxCommits > d.maxCommits {
			c.maxCommits = d.maxCommits
		}
	}
}

func (c *gitCounter) setTotalCommits() {
	for _, d := range c.directories {
		c.totalCommits += d.totalCommits
	}
}

// printResults will print the graph in the terminal
func (c gitCounter) printResults() {
	fmt.Sprintf("MAX", h.Comma(int64(c.maxCommits)),
		"TOTAL", h.Comma(int64(c.totalCommits)))
	// For showing the results starting at 0 to 23h
	var keys []int
	for k, _ := range c.results {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		line := fmt.Sprintf("%3v %7v ", k, c.results[k])
		line += colors[getColorIndex(c.results[k], c.maxCommits)]
		for n := 0; float64(n) < math.Floor(
			float64(c.results[k])/float64(c.maxCommits)*totalCols); n++ {
			line += "█"
		}
		fmt.Println(line)
	}
}

// This will call the git command and parse the commit
func (d directory) parseDir(c chan directory) {
	cmd := exec.Command("sh", "-c", d.gitCommand)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	// printCommand(cmd)
	err := cmd.Run()
	logError(err, &d)
	d.addDirectoryCommits(cmdOutput)
	d.setMaxCommits()
	d.setTotalCommits()
	c <- d
}

func main() {
	start := time.Now()
	projects := 0
	c := make(chan directory)
	count := gitCounter{}
	// For each directory in the flags
	for _, v := range flagDirectories {
		projectFolders, err := ioutil.ReadDir(v)
		logPanic(err)
		for _, f := range projectFolders {
			repoPath := v + "/" + f.Name()
			// Check if it's a git project
			if _, err := os.Stat(repoPath + "/.git"); err == nil {
				dir := getDir(repoPath)
				// Launch goroutine to process folder commits
				go dir.parseDir(c)
				projects++
			}
		}
	}
	// We ensure that each result gets processed one at a time
	// when receiving from the goroutine it will increment the total, then
	// continue with the next one
	for {
		dir := <-c
		count.directories = append(count.directories, dir)
		if len(count.directories) == projects {
			break
		}
	}
	count.setResults()
	count.printResults()
	log.Println(time.Since(start))
}
