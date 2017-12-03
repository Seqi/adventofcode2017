package main

import (
	"fmt"
	"math"
)

const input = 14

func main() {
	problemOne()
}

// 37   36   35   34   33    32   31
// 38   17   16   15   14    13   30
// 39   18    5    4    3    12   29
// 40   19    6    1    2    11   28
// 41   20    7    8    9    10   27
// 42   21   22   23   24    25   26  
// 43   44   45   46   47    48   49

func problemOne() {
	length := getMinGridLength(input)
	area := math.Pow(float64(length), 2)
	fmt.Println("Length of outer grid ", length)
	fmt.Println("Area of required grid ", area)

	radius := math.Pow(float64(length), 2) - math.Pow(float64(length-2), 2)
	fmt.Println("Radius  ", radius)

	squareStartPoint := area - radius
	fmt.Println("Starting point ", squareStartPoint)

	index := input - squareStartPoint
	fmt.Println("Index  ", index)

	side := math.Ceil(index / 4)
	fmt.Println("Side", side)

	lengthCenter := math.Ceil(float64(length / 2))
	fmt.Println("Length center", lengthCenter)

	
}

func getMinGridLength(num float64) int {
	sqrt := int(math.Ceil(math.Sqrt(num)))

	if sqrt%2 == 0 {
		return sqrt + 1
	} else {
		return sqrt
	}
}
