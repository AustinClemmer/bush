package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "cd", Description: "change the directory"},
		{Text: "ls", Description: "list the contents of current directory"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

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
	if parsed[0] == "cd" {
		if len(args) == 0 {
			os.Chdir(os.Getenv("HOME"))
			return
		}
		os.Chdir(args[0])
		return
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
	t := prompt.New(executor, completer, prompt.OptionPrefix(">>> "))
	t.Run()
}
