package main

import (
	_ "embed"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

//go:embed input/day22
var d22 string

func day22(input string) (int, int) {
	data := strings.TrimSuffix(input, "\n")
	pieces := strings.Split(data, "\n\n")
	s := game{
		grid:  parseGrid(pieces[0]),
		moves: parseMoves(pieces[1]),
		posY:  0,
		dir:   Right,
	}
	for i := 0; i < s.grid.SizeX(); i++ {
		if s.grid.Get(i, 0) == open {
			s.posX = i
			break
		}
	}

	return d22p1(s), d22p2(s)
}

type cell string

const (
	wall   cell = "#"
	open   cell = "."
	closed cell = " "
)

type dir int

const (
	Right dir = 0
	down  dir = 1
	Left  dir = 2
	Up    dir = 3
)

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Grid[T any] struct {
	sizeX, sizeY int
	matrix       [][]T
	empty        T
}

func NewGrid[T any](sizeX, sizeY int, empty T) *Grid[T] {
	matrix := make([][]T, sizeY)
	rows := make([]T, sizeX*sizeY)
	for i := 0; i < sizeX*sizeY; i++ {
		rows[i] = empty
	}

	j := 0
	for i := 0; i < sizeY; i++ {
		matrix[i] = rows[j : j+sizeX : j+sizeX]
		j += sizeX
	}
	return &Grid[T]{
		sizeX:  sizeX,
		sizeY:  sizeY,
		matrix: matrix,
		empty:  empty,
	}
}

func (g *Grid[T]) SizeX() int {
	return g.sizeX
}

func (g *Grid[T]) SizeY() int {
	return g.sizeY
}

func (g *Grid[T]) Get(x, y int) T {
	if x < 0 || x >= g.sizeX {
		return g.empty
	}
	if y < 0 || y >= g.sizeY {
		return g.empty
	}
	return g.matrix[y][x]
}

func (g *Grid[T]) Set(x, y int, v T) {
	if x < 0 || x >= g.sizeX {
		panic("invalid x")
	}
	if y < 0 || y >= g.sizeY {
		panic("invalid y")
	}
	g.matrix[y][x] = v
}

func (g *Grid[T]) StringWithFormatter(formatter func(T, int, int) string) string {
	var r strings.Builder
	for j := 0; j < g.sizeY; j++ {
		for i := 0; i < g.sizeX; i++ {
			_, err := r.WriteString(formatter(g.matrix[j][i], i, j))
			if err != nil {
				panic(err)
			}
		}
		_, err := r.WriteRune('\n')
		if err != nil {
			panic(err)
		}
	}
	return r.String()
}

type move struct {
	Right bool
	Left  bool
	move  int
}

type game struct {
	grid  *Grid[cell]
	moves []move
	posX  int
	posY  int
	dir   dir
}

func d22p1(s game) int {
	for _, m := range s.moves {
		s.playMove(m)
	}

	return 1000*(s.posY+1) + 4*(s.posX+1) + int(s.dir)
}

func d22p2(s game) int {
	for _, m := range s.moves {
		s.playMovePart2(m)
	}

	return 1000*(s.posY+1) + 4*(s.posX+1) + int(s.dir)
}

func parseGrid(input string) *Grid[cell] {
	lines := strings.Split(input, "\n")
	maxX := 0
	for _, line := range lines {
		maxX = Max(maxX, len(line))
	}

	g := NewGrid(maxX, len(lines), closed)
	for j, line := range lines {
		for i, c := range line {
			g.Set(i, j, cell(c))
		}
	}

	return g
}

func parseMoves(input string) []move {
	var r []move
	re := regexp.MustCompile(`([0-9]+)|(L)|(R)`)
	for input != "" {
		pieces := re.FindStringSubmatch(input)
		if pieces[1] != "" {
			m := move{move: mustAtoi(pieces[1])}
			r = append(r, m)
		} else if pieces[2] != "" {
			m := move{Left: true}
			r = append(r, m)
		} else if pieces[3] != "" {
			m := move{Right: true}
			r = append(r, m)
		} else {
			panic("bad input")
		}
		input = input[len(pieces[0]):]
	}
	return r
}

func (s *game) playMove(m move) {
	if m.Left {
		s.dir = (s.dir + 3) % 4
		return
	}
	if m.Right {
		s.dir = (s.dir + 1) % 4
		return
	}
	dx := 0
	dy := 0
	switch s.dir {
	case Right:
		dx = 1
	case down:
		dy = 1
	case Left:
		dx = -1
	case Up:
		dy = -1
	}
	for i := 0; i < m.move; i++ {
		newX := s.posX + dx
		newY := s.posY + dy
		c := s.grid.Get(newX, newY)
		switch c {
		case open:
			s.posX = newX
			s.posY = newY
		case wall:
			return
		case closed:
			// wrap around
			newX, newY, newD := s.wrap(newX, newY, dx, dy)
			c = s.grid.Get(newX, newY)
			switch c {
			case open:
				s.posX = newX
				s.posY = newY
				s.dir = newD
			case wall:
				return
			case closed:
				panic("unreachable")
			}
		}
	}
}

func (s *game) wrap(x, y, dx, dy int) (int, int, dir) {
	for {
		// go in reverse direction until we hit the other edge
		// this assumes the board is chasm-free.
		x = x - dx
		y = y - dy
		c := s.grid.Get(x, y)
		if c == closed {
			break
		}
	}
	return x + dx, y + dy, s.dir
}

func (s *game) playMovePart2(m move) {
	if m.Left {
		s.dir = (s.dir + 3) % 4
		return
	}
	if m.Right {
		s.dir = (s.dir + 1) % 4
		return
	}
	for i := 0; i < m.move; i++ {
		if !s.playSingleMovePart2() {
			break
		}
	}
}

func (s *game) playSingleMovePart2() bool {
	faceId := s.faceId(s.posX, s.posY)
	if faceId == -1 {
		panic("invalid game")
	}
	dx := 0
	dy := 0
	switch s.dir {
	case Right:
		dx = 1
	case down:
		dy = 1
	case Left:
		dx = -1
	case Up:
		dy = -1
	}
	newX := s.posX + dx
	newY := s.posY + dy
	newDir := s.dir
	newFaceId := s.faceId(newX, newY)
	if newFaceId != faceId || newFaceId == -1 {
		newX, newY, newDir = s.wrapPart2(faceId, s.posX, s.posY, s.dir)
	}
	c := s.grid.Get(newX, newY)
	switch c {
	case open:
		s.posX = newX
		s.posY = newY
		s.dir = newDir
		return true
	case wall:
		return false
	default:
		panic("unreachable")
	}
}

func (s *game) wrapPart2(faceId, x, y int, dir dir) (int, int, dir) {
	faces := []point{{50, 0}, {100, 0}, {50, 50}, {50, 100}, {0, 100}, {0, 150}}
	boxSize := 50

	relX := x - faces[faceId].x
	relY := y - faces[faceId].y
	switch faceId {
	case 0:
		switch dir {
		case Right:
			return x + 1, y, dir
		case down:
			return x, y + 1, dir
		case Left:
			newFace := 4
			newDir := Right
			newRelX := relX
			newRelY := boxSize - relY - 1
			if newRelX != 0 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case Up:
			newFace := 5
			newDir := Right
			newRelX := relY
			newRelY := relX
			if newRelX != 0 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		}
	case 1:
		switch dir {
		case Right:
			newFace := 3
			newDir := Left
			newRelX := relX
			newRelY := boxSize - relY - 1
			if newRelX != boxSize-1 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case down:
			newFace := 2
			newDir := Left
			newRelX := relY
			newRelY := relX
			if newRelX != boxSize-1 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case Left:
			return x - 1, y, dir
		case Up:
			newFace := 5
			newDir := Up
			newRelX := relX
			newRelY := boxSize - relY - 1
			if newRelY != boxSize-1 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		}
	case 2:
		switch dir {
		case Right:
			newFace := 1
			newDir := Up
			newRelX := relY
			newRelY := relX
			if newRelY != boxSize-1 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case down:
			return x, y + 1, dir
		case Left:
			newFace := 4
			newDir := down
			newRelX := relY
			newRelY := relX
			if newRelY != 0 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case Up:
			return x, y - 1, dir
		}
	case 3:
		switch dir {
		case Right:
			newFace := 1
			newDir := Left
			newRelX := relX
			newRelY := boxSize - relY - 1
			if newRelX != boxSize-1 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case down:
			newFace := 5
			newDir := Left
			newRelX := relY
			newRelY := relX
			if newRelX != boxSize-1 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case Left:
			return x - 1, y, dir
		case Up:
			return x, y - 1, dir
		}
	case 4:
		switch dir {
		case Right:
			return x + 1, y, dir
		case down:
			return x, y + 1, dir
		case Left:
			newFace := 0
			newDir := Right
			newRelX := relX
			newRelY := boxSize - relY - 1
			if newRelX != 0 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case Up:
			newFace := 2
			newDir := Right
			newRelX := relY
			newRelY := relX
			if newRelX != 0 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		}
	case 5:
		switch dir {
		case Right:
			newFace := 3
			newDir := Up
			newRelX := relY
			newRelY := relX
			if newRelY != boxSize-1 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case down:
			newFace := 1
			newDir := down
			newRelX := relX
			newRelY := boxSize - relY - 1
			if newRelY != 0 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case Left:
			newFace := 0
			newDir := down
			newRelX := relY
			newRelY := relX
			if newRelY != 0 {
				panic("unexpected")
			}
			newX := faces[newFace].x + newRelX
			newY := faces[newFace].y + newRelY
			return newX, newY, newDir
		case Up:
			return x, y - 1, dir
		}
	}
	panic("unreachable")
}

func (s *game) faceId(x, y int) int {
	var faces []point
	var boxSize int
	faces = []point{{50, 0}, {100, 0}, {50, 50}, {50, 100}, {0, 100}, {0, 150}}
	boxSize = 50
	for id, face := range faces {
		if (face.x <= x) && ((face.x + boxSize) > x) && (face.y <= y) && ((face.y + boxSize) > y) {
			return id
		}
	}
	return -1
}
