package main

import (
  "fmt"
)

var registers map[string]int

func main() {
  registers = make(map[string]int, 0)
}

type Instruction struct {
  target string
  operation string
  value int
  condition Condition
}

type Condition struct {
  target string
  operator string
  value int
}

func parseInput() []Instruction {
  regex := regex.Compile("([a-z]+)\s([a-z]{3})\s(-?\d+)\sif\s([a-z]+)\s([><!=]+)\s(-?\d+)")

}

var input string = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
