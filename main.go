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
	HistoryFile = filepath.Join(os.Getenv("HOME"), ".bush_history")
	rl, err     = readline.New(">>> ")
)

func main() {
	fmt.Println("Welcome to bush- the barely usable shell")
	defer rl.Close()
	readline.SetHistoryPath(HistoryFile)
	for {
		line, err := rl.Readline()
		errorCheck(err)
		err = executor(line)
		errorCheck(err)
		if f, err := os.Open(HistoryFile); err == nil {
			readline.AddHistory(line)
			f.Close()
		}
	}
}

func executor(s string) (e error) {
	e = nil
	var args []string
	parsed := strings.Fields(s)
	if len(parsed) == 0 {
		return
	}
	for _, val := range parsed[1:] {
		if val[0] == '$' {
			args = append(args, os.Getenv(val[1:]))
		} else {
			args = append(args, val)
		}
	}
	if last := s[len(s)-1:]; last == "&" {
		s = strings.TrimSuffix(s, "&")
		for _, val := range parsed[1:] {
			if val[0] == '$' {
				args = append(args, os.Getenv(val[1:]))
			} else {
				args = append(args, val)
			}
		}
		cmd := exec.Command(parsed[0], args...)
		err := cmd.Start()
		errorCheck(err)
		return
	}
	switch parsed[0] {
	case "exit":
		quitHandler()
	case "cd":
		e = directoryChangeHandler(args)
		return
	case "ls":
		args = listHandler(args)
	case "jobs":
		parsed[0] = "ps"
	}

	cmd := exec.Command(parsed[0], args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errorCheck(err)
	return
}

func errorCheck(the error) {
	if the != nil {
		fmt.Fprintln(rl, the)
	}
}

func quitHandler() {
	fmt.Println("Bye!")
	os.Exit(0)
	return
}

func directoryChangeHandler(args []string) (e error) {
	if len(args) == 0 {
		os.Chdir(os.Getenv("HOME"))
		return
	} else {
		if _, cdError := os.Stat(args[0]); cdError == nil {
			os.Chdir(args[0])
		} else {
			e = fmt.Errorf("CD: No such directory exists")
			return
		}
		return
	}
}

func listHandler(args []string) []string {
	if runtime.GOOS == "darwin" {
		return append(args, "-G")
	} else {
		return append(args, "--color")
	}
}
