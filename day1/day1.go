package main

import (
	"fmt"
	"os"
	"strconv"
)

var input string = os.Args[1]

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	total := solve(1)
	fmt.Printf("Problem One Result: %d\n", total)
}

func problemTwo() {
	total := solve(len(input) / 2)
	fmt.Printf("Problem Two Result: %d\n", total)
}

func solve(stepSize int) int {
	strLen := len(input)

	var total int = 0
	for i := 0; i < strLen; i++ {
		curr := string(input[i])
		next := string(input[(i+stepSize)%strLen])

		if curr == next {
			num, _ := strconv.Atoi(curr)
			total += num
		}
	}
	return total
}
