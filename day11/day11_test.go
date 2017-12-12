package main

import "testing"

func TestOne(t *testing.T) {
	input := []string{"ne", "ne", "ne"}
	expect := 3

	result := solve(input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestTwo(t *testing.T) {
	input := []string{"ne", "ne", "sw", "sw"}
	expect := 0

	result := solve(input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestThree(t *testing.T) {
	input := []string{"ne", "ne", "s", "s"}
	expect := 2

	result := solve(input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestFour(t *testing.T) {
	input := []string{"se", "sw", "se", "sw", "sw"}
	expect := 3

	result := solve(input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}
