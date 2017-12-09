package main

import "strings"
import "io/ioutil"
import "fmt"

const open string = "{"
const close string = "}"
const garbageOpen string = "<"
const garbageClose string = ">"
const skip string = "!"

func main() {
	input := readInput()

	score, garbageScore := solve(input)

	fmt.Println("Problem One result", score)
	fmt.Println("Problem Two result", garbageScore)
}

func solve(input string) (int, int) {
	score := 0
	garbageScore := 0
	chars := strings.Split(input, "")

	inGarbage := false
	currStreamScore := 0

	for i := 0; i < len(chars); i++ {

		if chars[i] == skip { // Skip if !
			i += 1
		} else if inGarbage { // In garbage, listen for garbage close
			if chars[i] == garbageClose {
				inGarbage = false
			} else {
				garbageScore += 1
			}
		} else {
			if chars[i] == garbageOpen { // Record if we're now in garbage
				inGarbage = true
			} else if chars[i] == open { // Increment worth on {
				currStreamScore += 1
			} else if chars[i] == close { // Stream closing, add stream score and decrement
				score += currStreamScore
				currStreamScore -= 1
			}
		}
	}

	return score, garbageScore
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading in input file")
	}

	return string(b)
}
