package main

import (
	"fmt"
	"math"
)

const input = 14

func main() {
	problemOne()
}

func problemOne() {
	gridLen := getMinGridLength(input)
	fmt.Println("Area required for input ", gridLen)

	radius := math.Pow(float64(gridLen), 2) - math.Pow(float64(gridLen-2), 2)
	fmt.Println("Radius  ", radius)

	squareStartPoint := int(math.Pow(float64(gridLen-2), 2))

	for j := 0; j < gridLen-1; j++ {
		fmt.Println(squareStartPoint)
	}

	fmt.Println("Starting point ", squareStartPoint)
}

func getMinGridLength(num float64) int {
	sqrt := int(math.Ceil(math.Sqrt(num)))

	if sqrt%2 == 0 {
		return sqrt + 1
	} else {
		return sqrt
	}
}
