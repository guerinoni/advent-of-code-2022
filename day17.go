package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input/day17
var d17 string

type rock []point

var rocks = []rock{
	// |..@@@@.|
	{{2, 0}, {3, 0}, {4, 0}, {5, 0}},

	// |...@...|
	// |..@@@..|
	// |...@...|
	{{3, 0}, {2, 1}, {3, 1}, {4, 1}, {3, 2}},

	// |....@..|
	// |....@..|
	// |..@@@..|
	{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},

	// |..@....|
	// |..@....|
	// |..@....|
	// |..@....|
	{{2, 0}, {2, 1}, {2, 2}, {2, 3}},

	// |..@@...|
	// |..@@...|
	{{2, 0}, {3, 0}, {2, 1}, {3, 1}},
}

var directions = map[byte]point{
	'<': {-1, 0},
	'>': {1, 0},
	'v': {0, -1},
}

func day17(input string) (int, int) {
	input = strings.TrimSuffix(input, "\n")
	height := 0
	busyPosition := map[point]bool{}
	directionIdx := 0
	numRocks := 1
	cache := map[string]struct {
		numRocks int
		height   int
	}{}

	p1 := 0
	p2 := 0

	for {
		r := rocks[(numRocks-1)%len(rocks)]

		newRock := make(rock, len(r))
		for i := 0; i < len(r); i++ {
			newRock[i].x = r[i].x
			newRock[i].y = r[i].y + height + 3
		}

		for {
			if directionIdx >= len(input) {
				directionIdx = 0
			}
			direction := input[directionIdx]
			directionIdx++

			var rockTemp rock
			for _, dir := range []point{directions[direction], directions['v']} {
				rockTemp = make(rock, len(newRock))
				for i, p := range newRock {
					p.x += dir.x
					p.y += dir.y

					if p.x < 0 || p.x > 6 || p.y < 0 {
						rockTemp = nil
						break
					}

					if _, ok := busyPosition[p]; ok {
						rockTemp = nil
						break
					}

					rockTemp[i] = p
				}

				if rockTemp != nil {
					newRock = rockTemp
				}
			}

			if rockTemp != nil {
				continue
			}

			for _, p := range newRock {
				busyPosition[p] = true
				if p.y+1 > height {
					height = p.y + 1
				}
			}

			if numRocks == 2022 {
				p1 = height
			}

			key := fmt.Sprintf("%d%d", (numRocks-1)%len(rocks), directionIdx-1)
			if val, ok := cache[key]; ok {
				quotient := (1000000000000 - numRocks) / (numRocks - val.numRocks)
				remainder := (1000000000000 - numRocks) % (numRocks - val.numRocks)

				if remainder == 0 {
					p2 = height + (height-val.height)*quotient
				}
			} else {
				cache[key] = struct {
					numRocks int
					height   int
				}{numRocks, height}
			}

			break
		}

		numRocks++
		if p1 != 0 && p2 != 0 {
			break
		}
	}

	return p1, p2
}
