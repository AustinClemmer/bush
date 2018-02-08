package main

import (
        "io"
	"bufio"
	"fmt"
	"github.com/pkg/term"
	"os"
	"os/exec"
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
        			return 1
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
	cmd := exec.Command(string(comm))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func printPrompt(stdout io.ReadWriteCloser) {
	fmt.Fprintf(stdout, ">>> ")
}
