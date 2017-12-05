package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := strconv.Atoi(os.Args[1])

	grid := buildGrid(in)
	x, y := find(in, grid)
	fmt.Printf("Target found at %d, %d. Manhattan Distance: %d", x, y, x+y)
}

func find(input int, grid map[string]int) (int, int) {
	for coords, val := range grid {
		if val == input {
			crds := strings.Split(coords, ", ")
			x, _ := strconv.Atoi(crds[0])
			y, _ := strconv.Atoi(crds[1])

			xI := int(math.Abs(float64(x)))
			yI := int(math.Abs(float64(y)))
			return xI, yI
		}
	}

	return 0, 0
}

func buildGrid(maxVal int) map[string]int {

	grid := make(map[string]int, 0)
	x, y, val, dir, sidelength := 0, 0, 1, 1, 1
	set := func(x int, y int, val int) {
		grid[strconv.Itoa(x)+", "+strconv.Itoa(y)] = val
	}

	set(x, y, val)

	for val <= maxVal {
		// X
		for i := 1; i <= sidelength; i++ {
			x += dir
			val = calculatePoint(x, y, grid)

			// Lazy..
			if val > maxVal {
				fmt.Printf("Found %d", val)
				os.Exit(0)
			}

			set(x, y, val)
		}

		// Y
		for i := 1; i <= sidelength; i++ {
			y += dir
			val = calculatePoint(x, y, grid)

			// Lazy..
			if val > maxVal {
				fmt.Printf("Found %d", val)
				os.Exit(0)
			}

			set(x, y, val)
		}

		sidelength += 1
		x *= -1
		y *= -1
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
