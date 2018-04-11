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

func executor(s string) (e error) {
	e = nil
	if last := s[len(s)-1:]; last == "&" {
		s = strings.TrimSuffix(s, "&")
		parsed := strings.Fields(s)
		var args []string
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
		//cmd.Wait() this will force the shell to wait until bg job is done
		return
	}
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
	if parsed[0] == "jobs" {
		fmt.Fprintln(rl, "NAME\t", "PID\t")
		fmt.Fprintln(rl, os.Args[0], "\t", os.Getpid())
		return
	}
	cmd := exec.Command(parsed[0], args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	errorCheck(err)
	return
}

func main() {
	fmt.Println("Welcome to bush- the belly up shell")
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

func errorCheck(the error) {
	if the != nil {
		fmt.Fprintln(rl, the)
	}
}
