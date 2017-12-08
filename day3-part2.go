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
		fmt.Println(val)
		if val > input {
			return val
		}
	}

	return 0
}

func buildGrid(maxVal int) map[string]int {

	grid := make(map[string]int, 0)
	x, y, val, dir, sidelength := 0, 0, 1, 1, 1
	set := func(x int, y int, val int) {
		grid[strconv.Itoa(x)+", "+strconv.Itoa(y)] = val
	}

	set(x, y, val)

	searching := true
	for searching {

		// X
		for i := 1; i <= sidelength; i++ {
			x += dir
			val = calculatePoint(x, y, grid)
			set(x, y, val)

			// Lazy..
			if val > maxVal {
				searching = false
				break
			}

		}

		if (!searching) {
			break
		}

		// Y
		for i := 1; i <= sidelength; i++ {
			y += dir
			val = calculatePoint(x, y, grid)
			set(x, y, val)

			// Lazy..
			if val > maxVal {
				searching = false
				break
			}

		}

		sidelength += 1
		dir *= -1
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
