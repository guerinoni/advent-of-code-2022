package main

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed input/day18
var d18 string

type cube struct {
	x int
	y int
	z int
}

func day18(input string) (int, int) {
	lines := strings.Split(input, "\n")
	cubes := map[cube]bool{}
	var xs, ys, zs []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		sides := strings.Split(line, ",")
		x := mustAtoi(sides[0])
		y := mustAtoi(sides[1])
		z := mustAtoi(sides[2])

		xs = append(xs, x)
		ys = append(ys, y)
		zs = append(zs, z)

		cubes[cube{x, y, z}] = true
	}

	sort.Ints(xs)
	sort.Ints(ys)
	sort.Ints(zs)

	return d18p1(cubes), d18p2(cubes, xs, ys, zs)
}

var neighbours = [][3]int{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}

func d18p1(cubes map[cube]bool) int {
	sides := 0

	for c := range cubes {
		sides += len(neighbours)
		for _, offset := range neighbours {
			possibleNeighbour := cube{c.x + offset[0], c.y + offset[1], c.z + offset[2]}
			if _, ok := cubes[possibleNeighbour]; ok {
				sides--
			}
		}
	}

	return sides
}

type cub struct {
	valid [][][]bool
	start [3]int
}

func d18p2(cubes map[cube]bool, xs, ys, zs []int) int {
	sides := 0

	startX, startY, startZ := xs[0]-1, ys[0]-1, zs[0]-1

	c := cub{
		start: [3]int{startX, startY, startZ},
	}
	c.valid = make([][][]bool, xs[len(xs)-1]-startX+4)
	for x := range c.valid {
		c.valid[x] = make([][]bool, ys[len(ys)-1]-startY+4)
		for y := range c.valid[x] {
			c.valid[x][y] = make([]bool, zs[len(zs)-1]-startZ+4)
		}
	}
	sides = countSides(c, cubes, c.start[0], c.start[1], c.start[2])
	return sides
}

func countSides(c cub, cubes map[cube]bool, x, y, z int) int {
	c.valid[x-c.start[0]][y-c.start[1]][z-c.start[2]] = true
	sum := 0

	for _, offset := range neighbours {
		cx, cy, cz := x+offset[0]-c.start[0], y+offset[1]-c.start[1], z+offset[2]-c.start[2]
		if cx < 0 || cx >= len(c.valid) || cy < 0 || cy >= len(c.valid[cx]) || cz < 0 || cz >= len(c.valid[cx][cy]) {
			continue
		} else if ok := c.valid[cx][cy][cz]; ok {
			continue
		}

		newX := x + offset[0]
		newY := y + offset[1]
		newZ := z + offset[2]
		if cubes[cube{newX, newY, newZ}] {
			sum++
		} else {
			sum += countSides(c, cubes, x+offset[0], y+offset[1], z+offset[2])
		}
	}
	return sum
}
