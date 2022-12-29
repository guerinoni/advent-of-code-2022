package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input/day1
var d1 string

func day1(input string) (int, int) {
	lines := strings.Split(input, "\n\n")
	top3 := []int{}
	for _, v := range lines {
		currentSum := 0
		numbers := strings.Split(v, "\n")
		for _, vv := range numbers {
			value, err := strconv.Atoi(vv)
			if err != nil {
				break
			}
			currentSum += value
		}

		if len(top3) < 3 {
			top3 = append(top3, currentSum)
			continue
		}

		for idx, v := range top3 {
			if currentSum > v {
				top3[idx] = currentSum
				break
			}
		}

		sort.Slice(top3, func(i, j int) bool {
			return i > j
		})
	}

	s := 0
	for _, t := range top3 {
		s += t
	}

	return top3[0], s
}
