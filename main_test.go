package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPrintedOutput(t *testing.T) {
	term, _ := ioutil.TempFile("", "term")
	stdin, _ := ioutil.TempFile("", "in")
	stdout, _ := ioutil.TempFile("", "out")
	stderr, _ := ioutil.TempFile("", "err")

	fmt.Fprintln(stdin, "exit")

	status := Process(term, stdin, stdout, stderr)

	if status != 0 {
		t.Error("Failed to exit cleanly")
	}
}
