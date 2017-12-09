package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

func main() {
	input = readInput()
	inputs := ParseInput()

	//  Find all programs that are owned by another program
	children := make([]string, 0)
	for _, program := range inputs {
		children = append(children, program.programs...)
	}

	// The program that is not owned by another program must be the bottom
	for _, program := range inputs {
		hasMatch := false
		for _, child := range children {
			if program.name == child {
				hasMatch = true
			}
		}

		if !hasMatch {
			fmt.Println("Problem One result: ", program.name)
			break
		}
	}

}

func ParseInput() []Program {
	lines := strings.Split(input, "\r\n")
	result := make([]Program, len(lines))

	for i, line := range lines {
		words := strings.Split(line, " ")
		weight, _ := strconv.Atoi(words[1][1:3])

		program := new(Program)
		program.name = words[0]
		program.weight = weight

		if len(words) > 2 {
			// Trim the comma/spaces
			program.programs = strings.Split(strings.Join(words[3:], ""), ",")
		}

		result[i] = *program
	}

	return result
}

type Program struct {
	name     string
	weight   int
	programs []string
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading in input file", err)
	}

	return string(b)
}
