package main

import (
	"os"
	"testing"
)

func TestPrintedOutputList(t *testing.T) {
	status := executor("ls")

	if status != nil {
		t.Error("LS failed to cleanly exit")
	}
}

func TestHome(t *testing.T) {
	status := executor("cd")
	hd := os.Getenv("HOME")
	wd, _ := os.Getwd()

	if status != nil {
		t.Error("Failed to exit cleanly")
	}
	if hd != wd {
		t.Error("Directories mismatched, failed to return home")
	}
}

func TestDirectoryNotExist(t *testing.T) {
	status := executor("cd bogusLongNameDirectory")

	if status != nil {
		t.Log("intentional failure, error caught")
	}
}

func TestExitChecker(t *testing.T) {
	statusExit := executor("exit")

	if statusExit != nil {
		t.Error("exit not clean")
	}
}
