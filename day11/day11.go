package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := parseInput()

	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	finalSteps, maxSteps := solve(input)
	fmt.Println("Problem One result:", finalSteps)
	fmt.Println("Problem Two result:", maxSteps)
}

func solve(route []string) (int, int) {
	x, y, z := 0, 0, 0
	distanceHistory := make([]int, 0)

	stepMap := map[string]func(){
		"n":  func() { y += 1; z -= 1 },
		"ne": func() { x += 1; z -= 1 },
		"se": func() { x += 1; y -= 1 },
		"s":  func() { y -= 1; z += 1 },
		"sw": func() { x -= 1; z += 1 },
		"nw": func() { x -= 1; y += 1 },
	}

	for _, val := range route {
		stepMap[val]()
		distanceHistory = append(distanceHistory, getDistance(x, y, z))
	}

	return getDistance(x, y, z), max(distanceHistory...)
}

func getDistance(x int, y int, z int) int {
	return max(abs(x), abs(y), abs(z))
}

func max(nums ...int) int {
	biggest := 0
	for _, i := range nums {
		if i > biggest {
			biggest = i
		}
	}

	return biggest
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func parseInput() ([]string, error) {
	file, err := readInput()

	if err != nil {
		return nil, err
	}

	return strings.Split(file, ","), nil
}

func readInput() (string, error) {
	file, err := ioutil.ReadFile("input.txt")

	return string(file), err
}
