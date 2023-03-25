package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input/day10
var d10 string

type operation struct {
	op string
	v  int
}

func day10(input string) (int, int) {
	rows := strings.Split(input, "\n")
	instructions := make([]operation, 0, len(rows))
	for _, row := range rows {
		if row == "" {
			continue
		}
		r := strings.Split(row, " ")
		if len(r) == 2 {
			instructions = append(instructions, operation{op: r[0], v: mustAtoi(r[1])})
		} else {
			instructions = append(instructions, operation{op: r[0], v: 0})
		}
	}

	return d10p1(instructions), d10p2(instructions)
}

func d10p2(instructions []operation) int {
	cycles := 0
	x := 1
	for i := 0; i < len(instructions); i++ {
		printCRT(cycles, x)
		cycles++

		if instructions[i].op == "addx" {
			printCRT(cycles, x)
			cycles++
			x += instructions[i].v
		}
	}

	fmt.Println()
	return 0
}

func d10p1(instructions []operation) int {
	var (
		cycles    int
		pc        int
		tmpCycle  int
		strengths int
	)
	x := 1
	for pc < len(instructions) {
		printCRT(cycles, x)
		cycles++
		switch cycles {
		case 20, 60, 100, 140, 180, 220:
			strengths += x * cycles
		}

		op := instructions[pc]
		switch op.op {
		case "addx":
			tmpCycle++
			if tmpCycle == 2 {
				printCRT(cycles, x)
				tmpCycle = 0
				x += op.v
				pc++
			}
		case "noop":
			tmpCycle = 0
			pc++
		}
	}

	return strengths
}

func printCRT(cycles, x int) {
	if cycles%40 == 0 && cycles <= 220 {
		fmt.Println()
	}
	if x-1 == cycles%40 || x == (cycles%40) || x+1 == (cycles%40) {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

func mustAtoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
