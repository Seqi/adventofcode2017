package main

import (
	"testing"
)

func TestSingleTwist(t *testing.T) {
	input := "3"
	expect := 2

	result := solve(input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestTwoTwists(t *testing.T) {
	input := "3,4"
	expect := 12

	result := solve(input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestMultipleTwists(t *testing.T) {
	input := "3,4,1,5"
	expect := 12

	result := solve(input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}
