package main

import "testing"

func TestProcessInput(t *testing.T) {
	got := processInputStr("1,0,0,0,99")
	if got != "2,0,0,0,99" {
		t.Errorf("processInput = %s; want 2,0,0,0,99", got)
	}

	got = processInputStr("2,3,0,3,99")
	if got != "2,3,0,6,99" {
		t.Errorf("processInput = %s; want 2,3,0,6,99", got)
	}

	got = processInputStr("2,4,4,5,99,0")
	if got != "2,4,4,5,99,9801" {
		t.Errorf("processInput = %s; want 2,4,4,5,99,9801", got)
	}

	got = processInputStr("1,1,1,4,99,5,6,0,99")
	if got != "30,1,1,4,2,5,6,0,99" {
		t.Errorf("processInput = %s; want 30,1,1,4,2,5,6,0,99", got)
	}

	got = processInputStr("1,9,10,3,2,3,11,0,99,30,40,50")
	if got != "3500,9,10,70,2,3,11,0,99,30,40,50" {
		t.Errorf("processInput = %s; want 3500,9,10,70,2,3,11,0,99,30,40,50", got)
	}
}
