package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string
var banksHistory []string

func main() {
	input = readInput()

	result := 0
	banksHistory = make([]string, 0)

	// Create a container for the current state of the banks
	banks := parseInput()
	banksHistory = append(banksHistory, banksToString(banks))

	// Copy the initial value of banks into the bank state ()
	var initialState = make([]int, len(banks))
	copy(initialState, banks)
	combinations := [][]int{initialState}

	for areCombinationsUnique() {
		banks = reallocate(banks)
		banksHistory = append(banksHistory, banksToString(banks))

		// Copy the new state into a bank state to preserve pointer
		var thisState = make([]int, len(banks))
		copy(thisState, banks)
		combinations = append(combinations, thisState)

		result += 1
	}

	// Find the first result that matches the duplicate
	dupeIndex, duplicate := 0, banksHistory[len(banksHistory)-1]

	for i, val := range banksHistory {
		if duplicate == val {
			dupeIndex = i
			break
		}
	}

	fmt.Printf("Problem One result: %d\n", result)
	fmt.Printf("Problem Two result: %d\n", len(banksHistory)-(dupeIndex+1))
}

func areCombinationsUnique() bool {
	// Only need to search the last added
	lastAddedCombination := banksHistory[len(banksHistory)-1]

	for i := 0; i < len(banksHistory)-1; i++ {
		if lastAddedCombination == banksHistory[i] {
			return false
		}
	}

	return true
}

func banksToString(bank []int) string {
	str := make([]string, 0)
	for _, val := range bank {
		str = append(str, strconv.Itoa(val))
	}

	return strings.Join(str, ",")
}

func reallocate(banks []int) []int {
	max, index := getMaxValueIndex(banks)

	// Reset
	banks[index] = 0

	// Distribute
	for i := 1; i <= max; i++ {
		currIdx := (index + i) % len(banks)
		banks[currIdx] += 1
	}

	return banks
}

func getMaxValueIndex(banks []int) (int, int) {
	index, max := 0, 0
	for i, val := range banks {
		if val > max {
			max = val
			index = i
		}
	}
	return max, index
}

func parseInput() []int {
	out := make([]int, 0)
	for _, in := range strings.Split(input, "\t") {
		val, _ := strconv.Atoi(in)
		out = append(out, val)
	}

	return out
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading in input file", err)
	}

	return string(b)
}
