package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14(t *testing.T) {
	input := `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`
	{
		p1, p2 := day14(input)
		assert.Equal(t, 24, p1)
		assert.Equal(t, 93, p2)
	}
	{
		p1, p2 := day14(d14)
		assert.Equal(t, 655, p1)
		assert.Equal(t, 26484, p2)
	}
}
