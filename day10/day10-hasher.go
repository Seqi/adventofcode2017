package main

import (
	"strconv"
	"strings"
)

func round(list []int, input string) int {
	stepsStr := strings.Split(input, ",")
	steps := make([]int, len(stepsStr))

	for i, val := range stepsStr {
		steps[i], _ = strconv.Atoi(val)
	}

	startIndex, skipSize := 0, 0
	for _, step := range steps {
		// If there's no step, do nothing
		if step > 0 {
			// Take the slice this step requires
			endIndex := (startIndex + step) % len(list)
			slice := slice(startIndex, endIndex, list)

			// Reverse the slice
			reverse := reverse(slice)

			// Check if we've overflowed
			hasOverflowed := startIndex+step > len(list)

			// We're good to just append
			if !hasOverflowed {
				// Append untouched start, the reversed array, then the rest of it
				list = append(list[0:startIndex], append(reverse, list[endIndex:]...)...)
			} else {
				// Find how many elements we've overflowed by
				overflowLen := (startIndex + step) % len(list)
				overflowRange := len(reverse) - overflowLen

				// List is strung together by:
				// End of reverse (len: overflow len), rest of list, start of reverse(len: overflow range))
				tail := append(list[overflowLen:len(list)-overflowRange], reverse[:overflowRange]...)
				list = append(reverse[overflowRange:], tail...)
			}
		}

		startIndex = (step + skipSize + startIndex) % len(list)
		skipSize += 1
	}

	return list[0] * list[1]
}

func slice(start int, end int, list []int) []int {
	if end > start {
		return list[start:end]
	}

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
