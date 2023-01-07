package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input/day4
var d4 string

func day4(input string) (int, int) {
	lines := strings.Split(input, "\n")
	fullContains := 0
	overlaps := 0

	for _, lll := range lines {
		var s1, e1, s2, e2 int
		parsed, err := fmt.Sscanf(lll, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)
		if err != nil || parsed != 4 {
			return fullContains, overlaps
		}

		if (s1 <= s2 && e1 >= e2) || (s2 <= s1 && e2 >= e1) {
			fullContains++
		}

		if s2 <= e1 && e2 >= s1 {
			overlaps++
		}

	}

	return fullContains, overlaps
}
