package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17(t *testing.T) {
	input := `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`
	{
		p1, p2 := day17(input)
		assert.Equal(t, 3068, p1)
		assert.Equal(t, 1514285714288, p2)
	}
	{
		p1, p2 := day17(d17)
		assert.Equal(t, 3184, p1)
		assert.Equal(t, 1577077363915, p2)
	}
}
