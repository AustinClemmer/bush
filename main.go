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
)

func executor(s string) (e error) {
	e = nil
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
	if parsed[0] == "cd" {
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
	if parsed[0] == "ls" {
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
		fmt.Printf("%s\n", err.Error())
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
		err = executor(line)
		if err != nil {
			fmt.Fprintln(rl, err)
		}
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
