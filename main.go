package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/term"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Command string

func main() {
	//fmt.Printf("bush >")

	//initialize terminal
	terminal, err := term.Open("/dev/tty")
	if err != nil {
		panic(err)
	}
	printPrompt(os.Stdout)

	//set restore point and restore term at end of program
	defer terminal.Restore()
	terminal.SetCbreak()

	status := Process(terminal, os.Stdin, os.Stdout, os.Stderr)
	if status != 0 {
		os.Exit(status)
	}
}

func Process(terminal, stdin, stdout, stderr io.ReadWriteCloser) int {
	//initialize buffer and reader
	reader := bufio.NewReader(terminal)

	//string command variable (global)
	var cmd Command

	//read in the characters/ build the command string
	//panic if errors
	for {
		character, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				//EOF is no worries
				return 0
			}
			panic(err)
		}
		switch character {
		case '\n':
			fmt.Fprintf(stdout, "\n")
			//handle command
			if cmd == "exit" || cmd == "quit" || cmd == "bye" {
				return 0
			} else if cmd == "" {
				printPrompt(stdout)
			} else {
				err := cmd.HandleCmd()
				if err != nil {
					fmt.Fprintf(stderr, "%v\n", err)
				}
				printPrompt(stdout)
			}
			cmd = ""
		case '\u0004':
			if len(cmd) == 0 {
				//		os.Exit(0)
				return 0
			}
		case '\u007f', '\u0008':
			if len(cmd) > 0 {
				cmd = cmd[:len(cmd)-1]
				fmt.Fprintf(stdout, "\u0008 \u0008")
			}
		default:
			fmt.Fprintf(stdout, "%c", character)
			cmd += Command(character)
		}
	}
}

func (comm Command) HandleCmd() error {
	parsed := strings.Fields(string(comm))
	if len(parsed) == 0 {
		printPrompt(os.Stdout)
		return nil
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
			return os.Chdir("/home")
		}
		return os.Chdir(args[0])
	}
	if parsed[0] == "ls" {
		args = append(args, "--color")
	}

	cmd := exec.Command(parsed[0], args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func printPrompt(stdout io.ReadWriteCloser) {
	fmt.Fprintf(stdout, ">>> ")
}
