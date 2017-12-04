package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var input int

var grid map[string]int
var x, y, val int = 0, 0, 1

func main() {
	in, _ := strconv.Atoi(os.Args[1])
	input = in
	problemOne()
}

func problemOne() {
	grid = make(map[string]int, 0)
	sidelength := 1

	set(x, y, val)

	for val <= input {

		// Right
		for i := 1; i <= sidelength; i++ {
			val += 1
			x += 1
			set(x, y, val)
		}

		// Up
		for i := 1; i <= sidelength; i++ {
			val += 1
			y += 1
			set(x, y, val)
		}

		sidelength += 1

		// Left
		for i := 1; i <= sidelength; i++ {
			val += 1
			x -= 1
			set(x, y, val)
		}

		// Down
		for i := 1; i <= sidelength; i++ {
			val += 1
			y -= 1
			set(x, y, val)
		}

		sidelength += 1
	}

	for coords, val := range grid {
		if val == input {
			crds := strings.Split(coords, ", ")
			x, _ := strconv.Atoi(crds[0])
			y, _ := strconv.Atoi(crds[1])

			xI := int(math.Abs(float64(x)))
			yI := int(math.Abs(float64(y)))
			fmt.Printf("Target found at %d, %d. Manhattan distance: %d\n", xI, yI, xI+yI)
		}
	}
}

func get(x int, y int) int {
	return grid[strconv.Itoa(x)+", "+strconv.Itoa(y)]
}

func set(x int, y int, val int) {
	grid[strconv.Itoa(x)+", "+strconv.Itoa(y)] = val
}
