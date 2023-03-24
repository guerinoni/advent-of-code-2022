package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input/day9
var d9 string

type instruction struct {
	dir  string
	dist int
}

type point struct {
	x int
	y int
}

func day9(input string) (int, int) {
	rows := strings.Split(input, "\n")
	instructions := make([]instruction, 0, len(rows))
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		v, err := strconv.Atoi(row[2:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{dir: row[:1], dist: v})
	}

	return impl(instructions, 2), impl(instructions, 10)
}

func impl(instructions []instruction, lenOfSnake int) int {
	pointsMarked := make(map[point]bool)
	snake := make([]point, lenOfSnake)

	for _, i := range instructions {
		for j := 0; j < i.dist; j++ {
			head := &snake[0]
			switch i.dir {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			for i := 1; i < lenOfSnake; i++ {
				followHead(snake[i-1], &snake[i])
			}

			pointsMarked[snake[lenOfSnake-1]] = true
		}
	}

	return len(pointsMarked)
}

func followHead(head point, tail *point) {
	p := point{head.x - tail.x, head.y - tail.y}

	yForward := p == point{-2, 1} || p == point{-1, 2} || p == point{0, 2} || p == point{1, 2} || p == point{2, 1} || p == point{2, 2} || p == point{-2, 2}
	xForward := p == point{1, 2} || p == point{2, 1} || p == point{2, 0} || p == point{2, -1} || p == point{1, -2} || p == point{2, 2} || p == point{2, -2}
	yBackward := p == point{2, -1} || p == point{1, -2} || p == point{0, -2} || p == point{-1, -2} || p == point{-2, -1} || p == point{2, -2} || p == point{-2, -2}
	xBackward := p == point{-1, -2} || p == point{-2, -1} || p == point{-2, -0} || p == point{-2, 1} || p == point{-1, 2} || p == point{-2, 2} || p == point{-2, -2}

	if yForward {
		tail.y++
	}
	if xForward {
		tail.x++
	}
	if yBackward {
		tail.y--
	}
	if xBackward {
		tail.x--
	}
}
