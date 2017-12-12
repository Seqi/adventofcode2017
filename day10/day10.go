package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()
	list := makeList()

	if err != nil {
		fmt.Println("Could not read input file")
		return
	}

	result := solve(list, input)
	fmt.Println("Problem One result:", result)
}

func solve(list []int, input string) int {
	stepsStr := strings.Split(input, ",")
	steps := make([]int, len(stepsStr))

	for i, val := range stepsStr {
		steps[i], _ = strconv.Atoi(val)
	}

	startIndex, skipSize := 0, 0
	for _, step := range steps {

		fmt.Println("Length: ", len(list))

		// If there's no step, increment everything and continue
		if step == 0 {
			startIndex = (step + skipSize + startIndex) % len(list)
			skipSize += 1
			continue
		}

		fmt.Println("Start index:", startIndex)

		fmt.Println("Skip size:", skipSize)
		fmt.Println("Step:", step)

		// Take the slice this step requires
		endIndex := (startIndex + step) % len(list)
		fmt.Println("End index:", endIndex)
		slice := slice(startIndex, endIndex, list)
		fmt.Println("Slice:", slice)

		// Reverse the slice
		reverse := reverse(slice)
		fmt.Println("Reversed", reverse)

		// Check if we've overflowed
		fmt.Println("test!", startIndex+step)
		hasOverflowed := startIndex+step > len(list)

		// We're good to just append
		if !hasOverflowed {
			fmt.Println("We haven't overflowed")
			// Append untouched start, the reversed array, then the rest of it
			list = append(list[0:startIndex], append(reverse, list[endIndex:]...)...)
		} else {
			fmt.Println("Overflowed")

			// Find how many elements we've overflowed by
			overflowLen := (startIndex + step) % len(list)
			overflowRange := len(reverse) - overflowLen

			fmt.Println("Overflow length:", overflowLen)
			fmt.Println("Overflow range:", overflowRange)

			// List is strung together by:
			// End of reverse (len: overflow len), rest of list, start of reverse(len: overflow range))
			tail := append(list[overflowLen:len(list)-overflowRange], reverse[:overflowRange]...)

			list = append(reverse[overflowRange:], tail...)
		}

		startIndex = (step + skipSize + startIndex) % len(list)
		skipSize += 1

		fmt.Println("new list", list)
		fmt.Println("New Start index:", startIndex)
		fmt.Println("\n---------\n")

	}

	return list[0] * list[1]
}

func slice(start int, end int, list []int) []int {
	if end > start {
		return list[start:end]
	}

	fmt.Println("Wrapping around..", start, end)

	return append(list[start:], list[:end]...)
}

func reverse(list []int) []int {
	result := make([]int, len(list))

	l := len(list) - 1

	for i := 0; i <= l; i++ {
		result[l-i] = list[i]
	}

	return result
}

func makeList() []int {
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}

	return list
}

func readInput() (string, error) {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
