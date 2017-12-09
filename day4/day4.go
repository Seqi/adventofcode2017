package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var input string

func main() {
	input = readInput()

	p1Result := problemOne()
	p2Result := problemTwo()

	fmt.Printf("Problem One result: %d\n", p1Result)
	fmt.Printf("Problem Two result: %d\n", p2Result)
}

func problemOne() int {
	count := 0
	for _, pwd := range strings.Split(input, "\n") {
		if areAllUnique(pwd) {
			count++
		}
	}

	return count
}

func problemTwo() int {
	count := 0
	for _, pwd := range strings.Split(input, "\n") {
		if areAllUniquePermutations(pwd) {
			count++
		}
	}

	return count
}

func areAllUniquePermutations(input string) bool {
	// Sort all entries before building freq map
	sortedWords := make([]string, 0)
	for _, word := range strings.Split(input, " ") {
		sortedWords = append(sortedWords, sortWord(word))
	}

	sortedPassword := strings.Join(sortedWords, " ")

	return areAllUnique(sortedPassword)
}

func sortWord(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func areAllUnique(input string) bool {
	freq := buildFreqMap(input)

	for _, val := range freq {
		if val > 1 {
			return false
		}
	}

	return true
}

func buildFreqMap(input string) map[string]int {
	freq := make(map[string]int)

	for _, word := range strings.Split(input, " ") {
		freq[word]++
	}

	return freq
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading in input file", err)
	}

	return string(b)
}
