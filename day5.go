package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input/day5
var d5 string

func day5(input string) (string, string) {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(strings.Split(parts[0], "\n"))
	stacksCopy := make([][]string, len(stacks))
	copy(stacksCopy, stacks)

	return d5impl(parts, stacks, true), d5impl(parts, stacksCopy, false)
}

func d5impl(parts []string, stacks [][]string, moveOneAtTime bool) string {
	for _, cmd := range strings.Split(parts[1], "\n") {
		var quantity, from, to int
		fmt.Sscanf(cmd, "move %d from %d to %d", &quantity, &from, &to)

		if moveOneAtTime {
			for i := 0; i < quantity; i++ {
				stacks[to-1] = append([]string{stacks[from-1][0]}, stacks[to-1]...)
				stacks[from-1] = append(stacks[from-1][1:])
			}
		} else {
			if cmd == "" {
				break
			}
			v := make([]string, quantity)
			copy(v, stacks[from-1][:quantity])
			stacks[to-1] = append(v, stacks[to-1]...)
			stacks[from-1] = append(stacks[from-1][quantity:])
		}
	}

	return composeSolution(stacks)
}

func composeSolution(stacks [][]string) string {
	var s string
	for _, st := range stacks {
		s += st[0]
	}
	return s
}

func parseStacks(startDiagram []string) [][]string {
	stacks := make([][]string, len(startDiagram[len(startDiagram)-1])/4+1)

	for _, line := range startDiagram {
		for idx, c := range line {
			if c == int32('[') {
				stacks[idx/4] = append(stacks[idx/4], string(line[idx+1]))
			}
		}
	}

	return stacks
}
