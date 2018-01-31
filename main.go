package main

import (
	"bufio"
	"fmt"
	//	"github.com/pkg/term"
	"os"
)

func main() {

	fmt.Printf("bush >")

	reader := bufio.NewReader(os.Stdin)
	for {
		character, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		fmt.Println(character)
	}
}
