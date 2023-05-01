package main

import (
	_ "embed"
	"strings"
)

//go:embed input/day8
var d8 string

func day8(input string) (int, int) {
	rows := strings.Split(input, "\n")

	var forest [][]int32

	totalVisibile := (len(rows))*2 + (len(rows[0])-2)*2 // edge

	for _, n := range rows {
		var row []int32
		for _, v := range n {
			row = append(row, v-'0')
		}

		forest = append(forest, row)
	}

	vlen := len(rows) - 1
	hlen := len(rows[0]) - 1

	maxScenic := 0

	for r, t := range forest {
		if r == 0 || r == vlen {
			continue
		}

		for c := range t {
			if c == 0 || c == hlen {
				continue
			}

			if isVisible(r, c, hlen, vlen, forest) {
				totalVisibile += 1
			}

			scenic := calcScenic(r, c, hlen, forest)
			if scenic > maxScenic {
				maxScenic = scenic
			}
		}
	}

	return totalVisibile, maxScenic
}

func isVisible(r, c, hlen, vlen int, forest [][]int32) bool {
	if rightSideVisible(r, c, vlen, forest) {
		return true
	}

	if leftSideVisible(r, c, forest) {
		return true
	}

	if bottomSideVisible(r, c, hlen, forest) {
		return true
	}

	if topSideVisible(r, c, forest) {
		return true
	}

	return false
}

func rightSideVisible(r int, c int, vlen int, forest [][]int32) bool {
	for i := c + 1; i <= vlen; i++ {
		if forest[r][c] <= forest[r][i] {
			return false
		}
	}

	return true
}

func leftSideVisible(r, c int, forest [][]int32) bool {
	for i := c - 1; i >= 0; i-- {
		if forest[r][c] <= forest[r][i] {
			return false
		}
	}

	return true
}

func bottomSideVisible(r, c, hlen int, forest [][]int32) bool {
	for i := r + 1; i <= hlen; i++ {
		if forest[r][c] <= forest[i][c] {
			return false
		}
	}

	return true
}

func topSideVisible(r, c int, forest [][]int32) bool {
	for i := r - 1; i >= 0; i-- {
		if forest[r][c] <= forest[i][c] {
			return false
		}
	}

	return true
}

func calcScenic(r int, c int, hlen int, forest [][]int32) int {
	right := calcRight(forest, r, c, forest[r][c], true)
	left := calcLeft(forest, r, c, forest[r][c], true)
	top := calcTop(forest, r, c, forest[r][c], true)
	bottom := calcBottom(forest, r, c, forest[r][c], true)

	return right * left * top * bottom
}

func calcRight(forest [][]int32, r, c int, house int32, firstCall bool) int {
	if c == len(forest[0])-1 || !firstCall && forest[r][c] >= house {
		return 0
	}
	return 1 + calcRight(forest, r, c+1, house, false)
}

func calcLeft(forest [][]int32, r, c int, house int32, firstCall bool) int {
	if c == 0 || !firstCall && forest[r][c] >= house {
		return 0
	}
	return 1 + calcLeft(forest, r, c-1, house, false)
}

func calcBottom(forest [][]int32, r, c int, house int32, firstCall bool) int {
	if r == len(forest)-1 || !firstCall && forest[r][c] >= house {
		return 0
	}
	return 1 + calcBottom(forest, r+1, c, house, false)
}

func calcTop(forest [][]int32, r, c int, house int32, firstCall bool) int {
	if r == 0 || !firstCall && forest[r][c] >= house {
		return 0
	}
	return 1 + calcTop(forest, r-1, c, house, false)
}
