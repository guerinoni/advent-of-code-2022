package main

import (
	_ "embed"
	"strings"
	"unicode"
)

//go:embed input/day3
var d3 string

func day3(input string) (int, int) {
	lines := strings.Split(input, "\n")

	return d3p1(lines), d3p2(lines)
}

func common(s1 string, s2 string) []rune {
	m1 := make(map[rune]int, 26)
	m2 := make(map[rune]int, 26)

	for _, s := range s1 {
		m1[s]++
	}
	for _, s := range s2 {
		m2[s]++
	}

	com := []rune{}
	for m := range m1 {
		if _, ok := m2[m]; ok {
			com = append(com, m)
		}
	}

	return com
}

func d3p1(lines []string) int {
	priority := 0
	for _, line := range lines {
		left := line[:len(line)/2]
		right := line[len(line)/2:]

		com := common(left, right)
		priority += calculatePriority(com)

	}
	return priority
}

func calculatePriority(com []rune) int {
	s := 0
	for _, a := range com {
		if unicode.IsUpper(a) {
			s += int(a) - 38
		}

		if unicode.IsLower(a) {
			s += int(a) - 96
		}
	}

	return s
}

func d3p2(lines []string) int {
	sum := 0
	group := []string{}
	for idx, line := range lines {
		group = append(group, line)

		if (idx+1)%3 == 0 {
			g1 := common(group[0], group[1])
			com := common(group[2], string(g1))
			sum += calculatePriority(com)
			group = []string{}
		}
	}

	return sum
}
