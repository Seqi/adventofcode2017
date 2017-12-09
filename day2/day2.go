package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

func main() {
	input = readInput()

	problemOne()
	problemTwo()
}

func problemOne() {
	total := 0

	for _, row := range strings.Split(input, "\n") {
		min, max := getMinMaxValueOfRow(row)
		total += max - min
	}

	fmt.Printf("Problem One Result: %d\n", total)
}

func problemTwo() {
	total := 0

	for _, row := range strings.Split(input, "\n") {
		result := getEvenDivisor(row)
		total += result
	}

	fmt.Printf("Problem Two Result: %d\n", total)
}

func getMinMaxValueOfRow(row string) (int, int) {
	min, max := 1<<63-1, 0
	for _, num := range convertRowToInts(row) {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func getEvenDivisor(row string) int {
	nums := convertRowToInts(row)

	for nIdx := 0; nIdx < len(nums); nIdx++ {
		for dIdx := nIdx + 1; dIdx < nIdx+len(nums); dIdx++ {
			if nums[nIdx]%nums[dIdx%len(nums)] == 0 {
				return nums[nIdx] / nums[dIdx%len(nums)]
			}
		}
	}

	return 0
}

func convertRowToInts(row string) []int {
	nums := make([]int, 0)

	for _, value := range strings.Split(row, "\t") {
		num, err := strconv.Atoi(value)
		if err == nil {
			nums = append(nums, num)
		}
	}

	return nums
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading in input file", err)
	}

	return string(b)
}
