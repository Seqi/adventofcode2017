package main

import (
  "fmt"
  "strings"
  "strconv"
)

var input = `5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6`
var banksHistory map[string]int

func main() {
  problemOne()
}

func problemOne() {
  result := 0
  banksHistory = make(map[string]int, 0)

  // Create a container for the current state of the banks
  banks := parseInput()

  // Copy the initial value of banks into the bank state
  var initialState = make([]int, len(banks))
  copy(initialState, banks)
  combinations := [][]int { initialState }

  for areCombinationsUnique() {
    banks = reallocate(banks)
    banksHistory[banksToString(banks)] += 1

    // Copy the new state into a bank state to preserve pointer
    var thisState = make([]int, len(banks))
    copy(thisState, banks)
    combinations = append(combinations, thisState)

    result += 1
  }

  fmt.Printf("Problem One result: %d", result)
}

func areCombinationsUnique() bool {
  for _, val := range banksHistory {
    if val > 1{
      return false
    }
  }

  return true
}

func banksToString(bank []int) string {
  str := make([]string, 0)
  for _, val := range bank {
    str = append(str, strconv.Itoa(val))
  }

  return strings.Join(str, ",")
}

func reallocate(banks []int) []int {
  max, index := getMaxValueIndex(banks)

  // Reset
  banks[index] = 0

  // Distribute
  for i := 1; i <= max; i++ {
    currIdx := (index + i) % len(banks)
    banks[currIdx] += 1
  }

  return banks
}

func getMaxValueIndex(banks []int) (int, int) {
  index, max := 0, 0
  for i, val := range banks {
    if val > max {
      max = val
      index = i
    }
  }
  return max, index
}

func parseInput() []int {
  out := make([]int, 0)
  for _, in := range strings.Split(input, "\t") {
    val, _ := strconv.Atoi(in)
    out = append(out, val)
  }

  return out
}
