package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	HistoryFile = filepath.Join(os.TempDir(), ".bush_history")
)

func executor(s string) {
	parsed := strings.Fields(s)
	if len(parsed) == 0 {
		return
	} else if parsed[0] == "quit" || parsed[0] == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	var args []string
	for _, val := range parsed[1:] {
		if val[0] == '$' {
			args = append(args, os.Getenv(val[1:]))
		} else {
			args = append(args, val)
		}
	}
	//TODO: WHEN TRYING TO CD DIR WHERE DIR DOESN'T EXIST RETURNS NEW PROMPT, WANT ERROR
	if parsed[0] == "cd" {
		if len(args) == 0 {
			os.Chdir(os.Getenv("HOME"))
			return
		} else {
			cdError := os.Chdir(args[0])
			if cdError != nil {
				fmt.Fprintf(os.Stdout, "Failed to successfully use cd")
				return
			}
			return
		}
	}
	if parsed[0] == "ls" { // THIS IS NAIVE
		if runtime.GOOS == "darwin" {
			args = append(args, "-G")
		} else {
			args = append(args, "--color")
		}
	}

	cmd := exec.Command(parsed[0], args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func main() {
	fmt.Println("Welcome to bush- the belly up shell")
	rl, err := readline.New(">>> ")
	errorCheck(err)
	defer rl.Close()
	readline.SetHistoryPath(HistoryFile)
	for {
		line, err := rl.Readline()
		errorCheck(err)
		executor(line)
		if f, err := os.Open(HistoryFile); err == nil {
			readline.AddHistory(line)
			f.Close()
		}
	}

}

func errorCheck(the error) {
	if the != nil {
		panic(the)
	}
}
