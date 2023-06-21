package main

import (
	_ "embed"
	"math"
	"strings"
)

//go:embed input/day23
var d23 string

func day23(input string) (int, int) {
	var elves []*point
	for j, line := range strings.Split(input, "\n") {
		for i, c := range line {
			if c == '#' {
				elves = append(elves, &point{x: i, y: j})
			}
		}
	}

	roundNoMoves := math.MaxInt
	after10Round := 0
	i := 0

	for {
		atLeastOneMove := execRound(elves, i)
		i++
		if !atLeastOneMove {
			roundNoMoves = i
			if i > 9 {
				break
			}
		}

		if i == 10 {
			minX, maxX, minY, maxY := 0, 0, 0, 0
			for _, e := range elves {
				minX = Min(minX, e.x)
				maxX = Max(maxX, e.x)
				minY = Min(minY, e.y)
				maxY = Max(maxY, e.y)
			}
			after10Round = (maxX-minX+1)*(maxY-minY+1) - len(elves)
		}
	}

	return after10Round, roundNoMoves
}

// execRound returns true if at least one elf moved
func execRound(elves []*point, round int) bool {
	proposalsCount := map[point]int{}
	proposals := make([]point, len(elves))
	for i, e := range elves {
		p := moveElf(elves, *e, round)
		proposalsCount[p]++
		proposals[i] = p
	}

	atLeastOneMove := false
	for i, e := range elves {
		p := proposals[i]
		if proposalsCount[p] == 1 {
			if !e.equalsPoint(p) {
				atLeastOneMove = true
			}
			*e = p
		}
	}

	return atLeastOneMove
}

// equalsPoint returns true if p and other are equal
func (p point) equalsPoint(other point) bool {
	return p.x == other.x && p.y == other.y
}

var (
	north = point{0, -1}
	south = point{0, 1}
	west  = point{-1, 0}
	east  = point{1, 0}
)

var (
	ne = point{1, -1}
	nw = point{-1, -1}
)

var (
	se = point{1, 1}
	sw = point{-1, 1}
)

var moves = [][]point{
	{north, ne, nw},
	{south, se, sw},
	{west, nw, sw},
	{east, ne, se},
}

func (p point) add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

func moveElf(elves []*point, elf point, round int) point {
	neighbors := 0
	for _, m := range []point{nw, north, ne, west, east, sw, south, se} {
		p := elf.add(m)
		if !available(elves, p) {
			neighbors++
		}
	}

	if neighbors == 0 {
		return elf
	}

outer:
	for i := round; i < round+4; i++ {
		m := moves[i%4]
		for j := 0; j < 3; j++ {
			p := elf.add(m[j])
			if !available(elves, p) {
				continue outer
			}
		}
		return elf.add(m[0])
	}

	return elf
}

func available(elves []*point, p point) bool {
	for _, elf := range elves {
		if elf.x == p.x && elf.y == p.y {
			return false
		}
	}
	return true
}
