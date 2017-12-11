package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := readInput()

	if err != nil {
		fmt.Println("Could not read input file")
		return
	}

	result := solve(input)
	fmt.Println("Problem One result:", result)
}

func solve(input string) int {
	return 0
}

func readInput() (string, error) {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
