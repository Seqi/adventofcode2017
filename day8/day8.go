package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var registers map[string]int

var operatorMap = map[string]func(int, int) bool{
	"==": func(left int, right int) bool { return left == right },
	"!=": func(left int, right int) bool { return left != right },
	">":  func(left int, right int) bool { return left > right },
	"<":  func(left int, right int) bool { return left < right },
	">=": func(left int, right int) bool { return left >= right },
	"<=": func(left int, right int) bool { return left <= right },
}

func main() {
	input := readInput()
	registers = make(map[string]int, 0)

	historyMax := 0
	for _, str := range strings.Split(input, "\n") {
		row := parseRow(str)
		newVal := processInstruction(row)

		if newVal > historyMax {
			historyMax = newVal
		}
	}

	postMax := 0
	for _, val := range registers {
		if val > postMax {
			postMax = val
		}
	}

	fmt.Println("Problem One result:", postMax)
	fmt.Println("Problem Two result:", historyMax)
}

func processInstruction(instruction Instruction) int {
	register := registers[instruction.condition.target]

	if operatorMap[instruction.condition.operator](register, instruction.condition.value) {
		if instruction.operation == "inc" {
			registers[instruction.target] += instruction.value
		} else {
			registers[instruction.target] -= instruction.value
		}
	}

	return registers[instruction.target]
}

func parseRow(row string) Instruction {
	rgx := regexp.MustCompile("([a-z]+) ([a-z]{3}) (-?\\d+) if ([a-z]+) ([><!=]+) (-?\\d+)")
	match := rgx.FindStringSubmatch(row)

	instruction := new(Instruction)
	instruction.target = match[1]
	instruction.operation = match[2]

	val1, _ := strconv.Atoi(match[3])
	instruction.value = val1

	instruction.condition = *new(Condition)
	instruction.condition.target = match[4]
	instruction.condition.operator = match[5]

	val2, _ := strconv.Atoi(match[6])
	instruction.condition.value = val2

	return *instruction
}

type Instruction struct {
	target    string
	operation string
	value     int
	condition Condition
}

type Condition struct {
	target   string
	operator string
	value    int
}

func readInput() string {
	b, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading in input file")
	}

	return string(b)
}
