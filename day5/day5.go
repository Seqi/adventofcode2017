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

	p1Result := problemOne()
	p2Result := problemTwo()

	fmt.Println("Problem One result: ", p1Result)
	fmt.Println("Problem Two result: ", p2Result)
}

func problemOne() int {
	escaped := false
	index, count := 0, 0
	in := parseInput()

	for escaped == false {
		num, err := tryGetValue(in, index)

		if err {
			escaped = true
		} else {
			in[index] = num + 1
			count += 1
			index += num
		}
	}

	return count
}

func problemTwo() int {
	escaped := false
	index, count := 0, 0
	in := parseInput()

	for escaped == false {
		num, err := tryGetValue(in, index)

		if err {
			escaped = true
		} else {
			if num >= 3 {
				in[index] = num - 1
			} else {
				in[index] = num + 1
			}

			count += 1
			index += num
		}
	}

	return count
}

func tryGetValue(arr []int, index int) (int, bool) {
	len := len(arr)
	if index < 0 || index >= len {
		return 0, true
	}

	return arr[index], false
}

func parseInput() []int {
	lines := strings.Split(input, "\r\n")

	steps := make([]int, 0, len(lines))
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		steps = append(steps, i)
	}

	return steps
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading in input file", err)
	}

	return string(b)
}
