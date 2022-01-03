package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Initialize a map if time of day to amount of commits
func get24HourMap() map[int]int {
	t := make(map[int]int)
	for i := 0; i < 24; i++ {
		t[i] = 0
	}
	return t
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

func (dirs *allDirectories) String() string {
	return "string"
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
