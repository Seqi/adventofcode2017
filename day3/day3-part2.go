package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, _ := strconv.Atoi(os.Args[1])

	grid := buildGrid(in)
	val := findMax(in, grid)
	fmt.Printf("\nProblem Two result: %d\n", val)
}

func findMax(input int, grid map[string]int) int {
	for _, val := range grid {
		if val > input {
			return val
		}
	}

	return 0
}

func buildGrid(maxVal int) map[string]int {

	grid := make(map[string]int, 0)
	x, y, direction, sidelength, val := 0, 0, 0, 1, 0
	set := func(x int, y int, val int) {
		grid[strconv.Itoa(x)+", "+strconv.Itoa(y)] = val
	}

	set(x, y, 1)

	for val < maxVal {

		isX := true
		sign := 1

		// Calculate axis and direction
		if direction%2 == 1 {
			isX = false
		}

		if direction > 1 {
			sign *= -1
		}

		for i := 1; i <= sidelength; i++ {

			if isX {
				x += sign
			} else {
				y += sign
			}

			val = calculatePoint(x, y, grid)
			set(x, y, val)

			if val > maxVal {
				break
			}
		}

		if direction%2 == 1 {
			sidelength += 1
		}

		direction = (direction + 1) % 4
	}

	return grid
}

func calculatePoint(x int, y int, grid map[string]int) int {
	val := get(x-1, y-1, grid) +
		get(x-1, y, grid) +
		get(x-1, y+1, grid) +
		get(x, y-1, grid) +
		get(x, y+1, grid) +
		get(x+1, y-1, grid) +
		get(x+1, y, grid) +
		get(x+1, y+1, grid)

	return val
}

func get(x int, y int, grid map[string]int) int {
	return grid[strconv.Itoa(x)+", "+strconv.Itoa(y)]
}
