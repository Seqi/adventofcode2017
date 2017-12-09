package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var input string
var inputs []Program

func main() {
	input = readInput()
	inputs = ParseInput()

	bottomProgram := getBottomProgram()

	r := findChildOfOffendingTower(bottomProgram)
	offendingTower := getParentProgram(r.name)

	// Find the weights of the offending tower
	weights := make([]int, len(offendingTower.programs))
	for i, p := range offendingTower.programs {
		weight := getProgram(p).calculateWeight()
		weights[i] = weight
	}

	// Get the diff of the offending tower
	_, i := isStackEven(weights)
	badWeight := weights[i]

	var goodIdx int
	if i == 0 {
		goodIdx = 1
	} else {
		goodIdx = 0
	}

	goodWeight := weights[goodIdx]
	diff := goodWeight - badWeight

	// Add the difference to the offending weight to get what it should be
	requiredWeight := r.weight + diff

	fmt.Println("Problem Two solution:", requiredWeight)
}

func findChildOfOffendingTower(program Program) Program {

	// If there's no programs we're at the top of the tower
	if len(program.programs) == 0 {
		return Program{}
	}

	// Calculate the weight of all programs above this one
	weights := make([]int, len(program.programs))
	for i, p := range program.programs {
		weight := getProgram(p).calculateWeight()
		weights[i] = weight
	}

	isStackEven, offendingIndex := isStackEven(weights)

	if isStackEven {
		return program
	}

	p := getProgram(program.programs[offendingIndex])
	return findChildOfOffendingTower(p)
}

func isStackEven(weights []int) (bool, int) {
	freq := make(map[int]int, 0)
	for i := 1; i < len(weights); i++ {
		freq[weights[i]] += 1
	}

	// Find the uneven tower weight
	isEven := true
	unevenWeight := -1
	for key, val := range freq {
		if val == 1 && len(freq) > 1 {
			isEven = false
			unevenWeight = key
		}
	}

	if !isEven {
		for i, val := range weights {
			if val == unevenWeight {
				return false, i
			}
		}
	}

	return true, -1
}

func (program Program) calculateWeight() int {
	total := program.weight

	for _, p := range program.programs {
		next := getProgram(p)
		total += next.calculateWeight()
	}

	return total
}

func getProgram(name string) Program {
	var result Program
	for _, program := range inputs {
		if program.name == name {
			result = program
			break
		}
	}

	return result
}

func getParentProgram(name string) Program {
	var result Program
	for _, program := range inputs {
		for _, child := range program.programs {
			if child == name {
				return program
			}
		}
	}

	return result
}

func getBottomProgram() Program {
	//  Find all programs that are owned by another program
	children := make([]string, 0)
	for _, program := range inputs {
		children = append(children, program.programs...)
	}

	// The program that is not owned by another program must be the bottom
	var result Program
	for _, program := range inputs {
		hasMatch := false
		for _, child := range children {
			if program.name == child {
				hasMatch = true
			}
		}

		if !hasMatch {
			result = program
			break
		}
	}

	return result
}

func ParseInput() []Program {
	lines := strings.Split(input, "\r\n")
	result := make([]Program, len(lines))
	rgx := regexp.MustCompile("(\\w+) \\((\\d+)\\)(?: -> (.+))?")

	for i, line := range lines {
		words := rgx.FindStringSubmatch(line)

		weight, _ := strconv.Atoi(words[2])

		program := new(Program)
		program.name = words[1]
		program.weight = weight

		if len(words) > 2 {
			// Trim the comma/spaces
			program.programs = strings.Split(words[3], ", ")
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
