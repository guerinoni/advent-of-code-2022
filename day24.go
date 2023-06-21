package main

import (
	_ "embed"
	"math"
	"strings"

	"github.com/oleiade/lane/v2"
)

//go:embed input/day24
var d24 string

type Blizzard [4][][]bool

func newBlizzard(minX, maxX, minY, maxY int) *Blizzard {
	var b Blizzard
	lenX := maxX - minX - 1
	lenY := maxY - minY - 1
	b[0] = make([][]bool, lenY)
	b[1] = make([][]bool, lenY)
	b[2] = make([][]bool, lenX)
	b[3] = make([][]bool, lenX)
	for j := 0; j < lenY; j++ {
		b[0][j] = make([]bool, lenX)
		b[1][j] = make([]bool, lenX)
	}
	for i := 0; i < lenX; i++ {
		b[2][i] = make([]bool, lenY)
		b[3][i] = make([]bool, lenY)
	}
	return &b
}

func (b *Blizzard) add(x, y int, dir byte) {
	switch dir {
	case '^':
		b[0][y][x] = true
	case 'v':
		b[1][y][x] = true
	case '<':
		b[2][x][y] = true
	case '>':
		b[3][x][y] = true
	}
}

func clone(s [][]bool) [][]bool {
	dup := make([][]bool, len(s))
	for i := range s {
		dup[i] = make([]bool, len(s[i]))
		copy(dup[i], s[i])
	}
	return dup
}

func (b *Blizzard) step() *Blizzard {
	newB := *b
	for i := 0; i < 4; i++ {
		newB[i] = clone(b[i])
	}
	newB[0] = append(newB[0][1:], newB[0][0])
	newB[1] = append(newB[1][len(newB[1])-1:], newB[1][:len(newB[1])-1]...)
	newB[2] = append(newB[2][1:], newB[2][0])
	newB[3] = append(newB[3][len(newB[3])-1:], newB[3][:len(newB[3])-1]...)
	return &newB
}

type State struct {
	pos     point
	minutes int
}

func neighbors(s State, blizzards []Blizzard) []State {
	i, j := s.pos.x, s.pos.y
	explore := []point{{i, j}, {i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	var res []State
	newTime := s.minutes + 1
	b := blizzards[newTime%len(blizzards)]

	exit1 := point{0, -1}
	exit2 := point{x: len(b[2]) - 1, y: len(b[0]) - 1 + 1}

	for _, p := range explore {
		if p == exit1 || p == exit2 {
			res = append(res, State{p, newTime})
			continue
		}
		if p.x < 0 || p.y < 0 || p.x >= len(b[0][0]) || p.y >= len(b[0]) {
			continue
		}
		if b[0][p.y][p.x] || b[1][p.y][p.x] || b[2][p.x][p.y] || b[3][p.x][p.y] {
			continue
		}
		res = append(res, State{p, newTime})
	}
	return res
}

func gdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	res := a * b / gdc(a, b)
	for i := 0; i < len(integers); i++ {
		res = lcm(res, integers[i])
	}
	return res
}

func bounds(grid map[point]uint8) (minX, maxX, minY, maxY int) {
	minX, minY, maxX, maxY = math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for p := range grid {
		minX = Min(p.x, minX)
		minY = Min(p.y, minY)
		maxX = Max(p.x, maxX)
		maxY = Max(p.y, maxY)
	}
	return minX, maxX, minY, maxY
}

type (
	heuristicFunction[T comparable] func(from T) int
	goalFunction[T comparable]      func(from T) bool
	costFunction[T comparable]      func(from, to T) int
)

func aStart(start State, goal goalFunction[State], cost costFunction[State], heuristic heuristicFunction[State], blizzards []Blizzard) (path []State, distance int) {
	queue := lane.NewMinPriorityQueue[State, int]()
	queue.Push(start, 0)

	cameFrom := map[State]State{start: start}
	costSoFar := map[State]int{start: 0}

	for {
		if queue.Size() == 0 {
			// There's no path, return found false.
			return
		}
		current, _, _ := queue.Pop()
		if goal(current) {
			// Found a path to the goal.
			var path []State
			curr := current
			for curr != start {
				path = append(path, curr)
				curr = cameFrom[curr]
			}
			return path, costSoFar[current]
		}

		for _, neighbor := range neighbors(current, blizzards) {
			newCost := costSoFar[current] + cost(current, neighbor)
			if _, ok := costSoFar[neighbor]; !ok || newCost < costSoFar[neighbor] {
				costSoFar[neighbor] = newCost
				priority := newCost + heuristic(neighbor)
				queue.Push(neighbor, priority)
				cameFrom[neighbor] = current
			}
		}
	}
}

func day24(input string) (int, int) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	grid := make(map[point]uint8)
	for j, l := range lines {
		for i, c := range l {
			grid[point{i, j}] = uint8(c)
		}
	}

	minX, maxX, minY, maxY := bounds(grid)
	b := newBlizzard(minX, maxX, minY, maxY)

	for p, v := range grid {
		if v == '^' || v == 'v' || v == '<' || v == '>' {
			b.add(p.x-1, p.y-1, v)
		}
	}

	blizzards := []Blizzard{*b}
	lenX := maxX - minX - 1
	lenY := maxY - minY - 1

	lcm := lcm(lenX, lenY)
	for i := 1; i < lcm; i++ {
		b = b.step()
		blizzards = append(blizzards, *b)
	}

	start := point{0, -1}
	goal := point{maxX - minX - 2, maxY - minY - 1}

	costF := func(from, to State) int { return 1 }
	goalF := func(s State) bool { return s.pos == goal }
	heuristicF := func(s State) int { return manhattanDistance(s.pos, goal) }

	_, cost := aStart(State{start, 0}, goalF, costF, heuristicF, blizzards)

	goal1 := func(s State) bool { return s.pos == goal }
	goal2 := func(s State) bool { return s.pos == start }

	heuristic2 := func(s State) int { return manhattanDistance(s.pos, start) }

	_, cost1 := aStart(State{start, 0}, goal1, costF, heuristicF, blizzards)
	_, cost2 := aStart(State{goal, cost1}, goal2, costF, heuristic2, blizzards)
	_, cost3 := aStart(State{start, cost1 + cost2}, goal1, costF, heuristicF, blizzards)

	return cost, cost1 + cost2 + cost3
}
