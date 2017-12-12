package main

import (
	"reflect"
	"testing"
)

func TestReverseEven(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expect := []int{6, 5, 4, 3, 2, 1}

	result := reverse(input)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestReverseOdd(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expect := []int{5, 4, 3, 2, 1}

	result := reverse(input)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestPlainSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expect := []int{3, 4, 5}

	result := slice(2, 5, input)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestCircularSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expect := []int{4, 5, 1, 2}

	result := slice(3, 2, input)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestSingleTwist(t *testing.T) {
	list := []int{0, 1, 2, 3, 4}
	input := "3"
	expect := 2

	result := round(list, input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestCircularTwist(t *testing.T) {
	list := []int{0, 1, 2, 3, 4}
	input := "3"
	expect := 2

	result := round(list, input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestTwoTwists(t *testing.T) {
	list := []int{0, 1, 2, 3, 4}
	input := "3,4"
	expect := 12

	result := round(list, input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}

func TestMultipleTwists(t *testing.T) {
	list := []int{0, 1, 2, 3, 4}
	input := "3,4,1,5"
	expect := 12

	result := round(list, input)

	if result != expect {
		t.Error("Expected", expect, "Got", result)
	}
}
