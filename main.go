package main

import (
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
	printPrompt()

	//set restore point and restore term at end of program
	defer terminal.Restore()
	terminal.SetCbreak()

	//initialize buffer and reader
	reader := bufio.NewReader(terminal)

	//string command variable (global)
	var cmd Command

	//read in the characters/ build the command string
	//panic if errors
	for {
		character, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		switch character {
		case '\n':
			fmt.Printf("\n")
			//handle command
			if cmd == "exit" || cmd == "quit" || cmd == "bye" {
				os.Exit(0)
			} else if cmd == "" {
				printPrompt()
			} else {
				err := cmd.HandleCmd()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
				}
				printPrompt()
			}
			cmd = ""

		case '\u007f', '\u0008':
			if len(cmd) > 0 {
				cmd = cmd[:len(cmd)-1]
				fmt.Printf("\u0008 \u0008")
			}
		default:
			fmt.Printf("%c", character)
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

func printPrompt() {
	fmt.Printf(">>> ")
}
