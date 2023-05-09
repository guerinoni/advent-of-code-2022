package main

import (
	_ "embed"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

//go:embed input/day20
var d20 string

func day20(input string) (int64, int64) {
	lines := strings.Split(input, "\n")

	var numbers []int64
	var numbersWithFactor []int64
	var indexes []int
	factor := int64(811589153)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		num, _ := strconv.ParseInt(lines[i], 10, 64)
		numbers = append(numbers, num)
		numbersWithFactor = append(numbersWithFactor, num*factor)
		indexes = append(indexes, i)
	}

	copyIndexes := make([]int, len(indexes))
	copy(copyIndexes, indexes)
	return d20p1(numbers, indexes), d20p2(numbersWithFactor, copyIndexes)
}

func d20p1(numbers []int64, indexes []int) int64 {
	l := int64(len(indexes) - 1)
	moveNumbers(numbers, indexes, l)
	zero := slices.Index(indexes, slices.Index(numbers, 0))
	length := len(numbers)
	v1 := numbers[indexes[modularIndex((zero+1000)%length, length)]]
	v2 := numbers[indexes[modularIndex((zero+2000)%length, length)]]
	v3 := numbers[indexes[modularIndex((zero+3000)%length, length)]]

	return v1 + v2 + v3
}

func d20p2(numbers []int64, indexes []int) int64 {
	l := int64(len(indexes) - 1)
	for n := 0; n < 10; n++ {
		indexes = moveNumbers(numbers, indexes, l)
	}

	zero1 := slices.Index(indexes, slices.Index(numbers, 0))
	v1 := numbers[indexes[modularIndex((zero1+1000)%len(numbers), len(numbers))]]
	v2 := numbers[indexes[modularIndex((zero1+2000)%len(numbers), len(numbers))]]
	v3 := numbers[indexes[modularIndex((zero1+3000)%len(numbers), len(numbers))]]

	return v1 + v2 + v3
}

func moveNumbers(numbers []int64, indexes []int, l int64) []int {
	for i := range numbers {
		j := slices.Index(indexes, i)
		indexes = slices.Delete(indexes, j, j+1)
		newPosition := int((int64(j) + numbers[i]) % l)
		idx := modularIndex(newPosition, int(l))
		indexes = slices.Insert(indexes, idx, i)
	}
	return indexes
}

func modularIndex(i, mod int) int {
	if i < 0 {
		return i + mod
	}
	return i
}
