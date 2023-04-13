package main

import (
	_ "embed"
	"strings"
)

//go:embed input/day14
var d14 string

type path struct {
	points []point
}

func day14(input string) (int, int) {
	data := strings.Split(input, "\n")
	var paths []path

	for _, line := range data {
		if line == "" {
			continue
		}
		points := strings.Split(line, " -> ")
		var onePath path
		for _, pp := range points {
			p := strings.Split(pp, ",")
			ppp := point{x: mustAtoi(p[0]), y: mustAtoi(p[1])}
			onePath.points = append(onePath.points, ppp)
		}
		paths = append(paths, onePath)
	}

	field := make(map[point]bool)

	minX := 500
	maxX := 500
	maxY := 9 // Because is only 1 digit after the comma
	for _, path := range paths {
		for i := 0; i < len(path.points)-1; i++ {
			start := path.points[i]
			end := path.points[i+1]
			if start.x < minX {
				minX = start.x
			}
			if start.x > maxX {
				maxX = start.x
			}
			if start.y > maxY {
				maxY = start.y
			}
			for _, p := range PointsBetween(start, end) {
				field[p] = true
			}
		}
	}

	fieldCopy := make(map[point]bool, len(field))
	for k, v := range field {
		fieldCopy[k] = v
	}

	// Add floor of rocks at maxY + 2
	for x := minX - 300; x <= maxX+300; x++ { // I used 300 and infinity
		fieldCopy[point{x, maxY + 2}] = true
	}

	return simulateFallSand(field, maxY), simulateFallSandWithoutVoid(fieldCopy)
}

func simulateFallSand(field map[point]bool, maxY int) int {
	intoVoid := false
	sand := 0
	for !intoVoid {
		newSand := point{500, 0}
		for {
			if newSand.y+1 > maxY {
				intoVoid = true
				break
			}
			if !field[point{newSand.x, newSand.y + 1}] { // down
				newSand.y++
			} else if !field[point{newSand.x - 1, newSand.y + 1}] { // (down) left
				newSand.y++
				newSand.x--
			} else if !field[point{newSand.x + 1, newSand.y + 1}] { // (down) right
				newSand.y++
				newSand.x++
			} else {
				field[newSand] = true
				sand++
				break
			}
		}
	}

	return sand
}

func simulateFallSandWithoutVoid(field map[point]bool) int {
	sand := 0
	for {
		newSand := point{500, 0}
		if field[newSand] {
			break
		}
		for {
			if !field[point{newSand.x, newSand.y + 1}] {
				newSand.y++
			} else if !field[point{newSand.x - 1, newSand.y + 1}] {
				newSand.y++
				newSand.x--
			} else if !field[point{newSand.x + 1, newSand.y + 1}] {
				newSand.y++
				newSand.x++
			} else {
				field[newSand] = true
				sand++
				break
			}
		}
	}

	return sand
}

// PointsBetween returns all points between two points, including the start and end points.
func PointsBetween(start, end point) []point {
	var points []point
	if start.x == end.x {
		if start.y > end.y {
			start, end = end, start
		}
		for y := start.y; y <= end.y; y++ {
			points = append(points, point{x: start.x, y: y})
		}
	} else if start.y == end.y {
		if start.x > end.x {
			start, end = end, start
		}
		for x := start.x; x <= end.x; x++ {
			points = append(points, point{x: x, y: start.y})
		}
	}
	return points
}
