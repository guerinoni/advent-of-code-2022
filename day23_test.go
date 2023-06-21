package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay23(t *testing.T) {
	input := `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`
	{
		p1, p2 := day23(input)
		assert.Equal(t, 110, p1)
		assert.Equal(t, 20, p2)
	}
	{
		p1, p2 := day23(d23)
		assert.Equal(t, 4056, p1)
		assert.Equal(t, 999, p2)
	}
}
