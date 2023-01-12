package main

import (
	_ "embed"
)

//go:embed input/day6
var d6 string

func day6(input string) (int, int) {
	return d6Impl(input, 4), d6Impl(input, 14)
}

func d6Impl(input string, windowSize int8) int {
	last4 := make([]rune, windowSize)
	for idx, c := range input {
		if idx < 4 {
			last4[idx] = c
			continue
		}

		if !hasDuplicated(last4) {
			return idx
		}

		shiftAndAppend(&last4, c)
	}

	return 0
}

func shiftAndAppend(slice *[]rune, c rune) {
	for i := 0; i < len(*slice)-1; i++ {
		(*slice)[i] = (*slice)[i+1]
	}

	(*slice)[len(*slice)-1] = c
}

func hasDuplicated(last4 []rune) bool {
	dup := make([]bool, 256)
	for _, c := range last4 {
		if dup[c] {
			return true
		}
		dup[c] = true
	}

	return false
}
