package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// get24HourMap initializes a map if time of hour of day to amount of commits
func get24HourMap() map[int]int {
	t := make(map[int]int)
	for i := 0; i < 24; i++ {
		t[i] = 0
	}
	return t
}

// getPad returns padding
func getPad(padding int) string {
	return strings.Repeat("#", padding)
}

// printDoc prints a documentation line
func printDoc(contents string, center bool) {
	if center {
		contents = fmt.Sprintf(
			"%"+strconv.Itoa((padding*2)+len(contents)/2)+"s", contents)
	}
	fmt.Printf(
		"%10s%-"+strconv.Itoa(colsNoPadding)+"s%s\n",
		getPad(padding), contents, getPad(padding))
}

// printCommand prints a shell command
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

// logError prints an error on std err if present, with the path of the
// directory
func logError(err error, d *directory) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s %s\n", d.path, err.Error()))
	}
}

// logPanic will panic with the error if it's present
func logPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// String method for the flag parsing
func (dirs *allDirectories) String() string {
	return "string"
}

// getColorIndex will return the index of the color in the graph
// by scaling the maximum to len(colors)
func getColorIndex(value int, maxValue int) int {
	scale := (float64(len(colors)-2) / float64(maxValue))
	return int((float64(value) * scale)) + 1 // this will skip the first value
}

func (dirs *allDirectories) Set(value string) error {
	*dirs = append(*dirs, value)
	return nil
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
