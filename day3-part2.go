package main

import (
	"fmt"
	"os"
	"strconv"
)

var grid map[string]int
var x, y, val int = 0, 0, 1
var inputFound bool
var input int

func main() {
	in, _ := strconv.Atoi(os.Args[1])
	input = in
	problemOne()
}

func problemOne() {
	grid = make(map[string]int, 0)
	sidelength := 1

	set(x, y, val)

	for !inputFound {

		// Right
		for i := 1; i <= sidelength; i++ {
			x += 1
			val = calculatePoint(x, y)
			set(x, y, val)
		}

		// Up
		for i := 1; i <= sidelength; i++ {
			y += 1
			val = calculatePoint(x, y)
			set(x, y, val)
		}

		sidelength += 1

		// Left
		for i := 1; i <= sidelength; i++ {
			x -= 1
			val = calculatePoint(x, y)
			set(x, y, val)
		}

		// Down
		for i := 1; i <= sidelength; i++ {
			y -= 1
			val = calculatePoint(x, y)
			set(x, y, val)
		}

		sidelength += 1
	}
}

func calculatePoint(x int, y int) int {
	val := get(x-1, y-1) +
		get(x-1, y) +
		get(x-1, y+1) +
		get(x, y-1) +
		get(x, y+1) +
		get(x+1, y-1) +
		get(x+1, y) +
		get(x+1, y+1)

	// Lazy..
	if val > input {
		inputFound = true
		fmt.Printf("Found %d", val)
		os.Exit(0)
	}

	return val
}

func get(x int, y int) int {
	return grid[strconv.Itoa(x)+", "+strconv.Itoa(y)]
}

func set(x int, y int, val int) {
	grid[strconv.Itoa(x)+", "+strconv.Itoa(y)] = val
}
