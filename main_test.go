package main

import (
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

	if status != nil {
		t.Error("Failed to exit cleanly")
	}
}
