package main

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed input/day12
var d12 string

func day12(input string) (int, int) {
	lines := strings.Split(input, "\n")

	heightmap := make([][]rune, 0)
	var start point
	var target point

	var possibleStarts []point
	for _, l := range lines {
		line := []rune(l)
		for i := 0; i < len(line); i++ {
			if line[i] == 'S' {
				start = point{i, len(heightmap)}
				line[i] = 'a'
			}

			if line[i] == 'a' {
				possibleStarts = append(possibleStarts, point{i, len(heightmap)})
			}

			if line[i] == 'E' {
				target = point{i, len(heightmap)}
				line[i] = 'z'
			}
		}
		heightmap = append(heightmap, line)
	}

	r := distanceFrom(start, target, heightmap)

	var distances []int
	for _, s := range possibleStarts {
		if d := distanceFrom(s, target, heightmap); d > 0 {
			distances = append(distances, d)
		}
	}
	sort.Ints(distances)

	return r, distances[0]
}

func distanceFrom(start, target point, heightmap [][]rune) int {
	visited := map[point]bool{start: true}
	toVisit := []point{start}
	current := toVisit[0]
	distanceFromStart := map[point]int{start: 0}

	for current != target && len(toVisit) > 0 {
		current = toVisit[0]
		visited[current] = true
		toVisit = toVisit[1:]

		dist := distanceFromStart[current] + 1
		for _, neigh := range neighbours() {
			i := neigh[0]
			j := neigh[1]
			nextPoint := point{current.x + j, current.y + i}
			if visited[nextPoint] || !isInMap(nextPoint, heightmap) {
				continue
			}

			if heightmap[nextPoint.y][nextPoint.x]-heightmap[current.y][current.x] <= 1 {
				if distanceFromStart[nextPoint] == 0 {
					toVisit = append(toVisit, nextPoint)
					distanceFromStart[nextPoint] = dist
				}

				if distanceFromStart[nextPoint] > dist {
					distanceFromStart[nextPoint] = dist
				}
			}
		}
	}

	return distanceFromStart[target]
}

func isInMap(p point, heightmap [][]rune) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(heightmap[0]) && p.y < len(heightmap)
}

func neighbours() [][]int {
	return [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
}
