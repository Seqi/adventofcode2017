package main

import (
	"testing"
)

func TestStreamOne(t *testing.T) {
	in := "{}"
	expected := 1
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestStreamTwo(t *testing.T) {
	in := "{{{}}}"
	expected := 6
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestStreamThree(t *testing.T) {
	in := "{{},{}}"
	expected := 5
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestStreamFour(t *testing.T) {
	in := "{{{},{},{{}}}}"
	expected := 16
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestStreamFive(t *testing.T) {
	in := "{<a>,<a>,<a>,<a>}"
	expected := 1
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbage1(t *testing.T) {
	in := "{{<!!>},{<!!>},{<!!>},{<!!>}}"
	expected := 9
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbage2(t *testing.T) {
	in := "{{<ab>},{<ab>},{<ab>},{<ab>}}"
	expected := 9
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbage3(t *testing.T) {
	in := "{!{<ab>!}}"
	expected := 1
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbage4(t *testing.T) {
	in := "{{<a!>},{<a!>},{<a!>},{<ab>}}"
	expected := 3
	result, _ := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbageScore1(t *testing.T) {
	in := "{<>}"
	expected := 0
	_, result := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbageScore2(t *testing.T) {
	in := "{<random characters>}"
	expected := 17
	_, result := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbageScore3(t *testing.T) {
	in := "{<<<<>}"
	expected := 3
	_, result := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbageScore4(t *testing.T) {
	in := "{<{!>}>}"
	expected := 2
	_, result := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbageScore5(t *testing.T) {
	in := "{<!!>}"
	expected := 0
	_, result := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbageScore6(t *testing.T) {
	in := "{<!!!>>}"
	expected := 0
	_, result := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestGarbageScore7(t *testing.T) {
	in := "{<{o\"i!a,<{i<a>}"
	expected := 10
	_, result := solve(in)

	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}
