package main

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed input/day11
var d11 string

type monkey struct {
	operation     string
	startingItems []int
	testDivisible int
	ifTrue        int
	ifFalse       int
}

func day11(input string) (int, int) {
	info := strings.Split(input, "\n\n")
	monkeys := make([]monkey, len(info))
	moderator := 1
	for idx, monkey := range info {
		lines := strings.Split(monkey, "\n")

		numbers := strings.Split(lines[1], ": ")[1]
		ns := strings.Split(numbers, ", ")
		for _, n := range ns {
			monkeys[idx].startingItems = append(monkeys[idx].startingItems, mustAtoi(n))
		}

		monkeys[idx].operation = strings.Split(lines[2], "new = old ")[1]

		test := strings.Split(lines[3], "Test: divisible by ")[1]
		monkeys[idx].testDivisible = mustAtoi(test)
		moderator *= monkeys[idx].testDivisible

		monkeys[idx].ifTrue = mustAtoi(strings.Split(lines[4], "If true: throw to monkey ")[1])
		monkeys[idx].ifFalse = mustAtoi(strings.Split(lines[5], "If false: throw to monkey ")[1])
	}

	monkeysCopy := make([]monkey, len(monkeys))
	copy(monkeysCopy, monkeys)
	return impl11(monkeys, 20, 3, moderator), impl11(monkeysCopy, 10000, 0, moderator)
}

func impl11(monkeys []monkey, turns int, worryLevel int, moderator int) int {
	counts := make([]int, len(monkeys))

	for turn := 0; turn < turns; turn++ {
		for monkeyIdx, currentMonkey := range monkeys {
			for _, item := range currentMonkey.startingItems {
				n := 0
				if strings.Contains(currentMonkey.operation, "old") {
					n = item
				} else {
					n = mustAtoi(currentMonkey.operation[2:])
				}

				if currentMonkey.operation[0] == '+' {
					item += n
				} else {
					item *= n
				}

				if worryLevel > 0 {
					item /= worryLevel
				}

				if moderator > 0 {
					item %= moderator
				}

				toThrow := 0
				if item%currentMonkey.testDivisible == 0 {
					toThrow = currentMonkey.ifTrue
				} else {
					toThrow = currentMonkey.ifFalse
				}
				monkeys[toThrow].startingItems = append(monkeys[toThrow].startingItems, item)
			}
			counts[monkeyIdx] += len(monkeys[monkeyIdx].startingItems)
			monkeys[monkeyIdx].startingItems = []int{}
		}
	}

	sort.Ints(counts)
	return counts[len(counts)-1] * counts[len(counts)-2]
}
