package main

import (
	_ "embed"
	"strings"
)

//go:embed input/day21
var d21 string

type monkey21 struct {
	name   string
	op     string
	number int
	wait   []string
}

func day21(input string) (int, int) {
	lines := strings.Split(input, "\n")
	mapOfMonkeys := make(map[string]monkey21)
	for _, line := range lines {
		if line == "" {
			continue
		}
		var m monkey21
		first := strings.Split(line, ": ")
		m.name = first[0]
		data := strings.Split(first[1], " ")
		if len(data) == 3 {
			m.wait = append(m.wait, data[0])
			m.wait = append(m.wait, data[2])
			m.op = data[1]
		} else {
			m.number = mustAtoi(data[0])
		}

		mapOfMonkeys[m.name] = m
	}

	p1 := calculateYall("root", mapOfMonkeys)

	root := mapOfMonkeys["root"]
	m1 := mapOfMonkeys[root.wait[0]]
	m2 := mapOfMonkeys[root.wait[1]]
	p2 := calculateYallWanted(m1, m2.number, mapOfMonkeys)

	return p1, p2
}

func calculateYall(name string, mapOfMonkeys map[string]monkey21) int {
	m := mapOfMonkeys[name]
	if m.number != 0 {
		return m.number
	}

	left := calculateYall(m.wait[0], mapOfMonkeys)
	right := calculateYall(m.wait[1], mapOfMonkeys)

	mL := mapOfMonkeys[m.wait[0]]
	mL.number = left
	mapOfMonkeys[m.wait[0]] = mL

	mR := mapOfMonkeys[m.wait[1]]
	mR.number = right
	mapOfMonkeys[m.wait[1]] = mR

	switch m.op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}

	panic("unknown op")
}

func calculateYallWanted(m monkey21, wanted int, mapOfMonkeys map[string]monkey21) int {
	if m.name == "humn" {
		return wanted
	}

	if m.wait == nil {
		return 0
	}

	m1 := mapOfMonkeys[m.wait[0]]
	m2 := mapOfMonkeys[m.wait[1]]

	var newWanted int
	switch m.op {
	case "+":
		newWanted = wanted - m2.number
	case "-":
		newWanted = wanted + m2.number
	case "*":
		newWanted = wanted / m2.number
	case "/":
		newWanted = wanted * m2.number
	}

	if v := calculateYallWanted(m1, newWanted, mapOfMonkeys); v > 0 {
		return v
	}

	switch m.op {
	case "+":
		newWanted = wanted - m1.number
	case "-":
		newWanted = -wanted + m1.number
	case "*":
		newWanted = wanted / m1.number
	case "/":
		newWanted = wanted * m1.number
	}

	return calculateYallWanted(m2, newWanted, mapOfMonkeys)
}
