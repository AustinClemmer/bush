package main

import (
	"os"
	"testing"
)

func TestPrintedOutputList(t *testing.T) {
	status := executor("ls")

	if status != nil {
		t.Error("Failed to exit cleanly")
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
