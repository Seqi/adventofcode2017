package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := readInput()
	list := makeList()

	if err != nil {
		fmt.Println("Could not read input file")
		return
	}

	result := round(list, input)
	fmt.Println("Problem One result:", result)
}

func makeList() []int {
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}

	return list
}

func readInput() (string, error) {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
