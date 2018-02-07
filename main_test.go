package main

import "testing"

func TestPrintedOutput(t *testing.T) {
	numberOne := 1
	if numberOne == 0 {
		t.Error("Dev needs more coffee")
	}
}
