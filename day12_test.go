package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay12(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

	p1, p2 := day12(input)
	assert.Equal(t, 31, p1)
	assert.Equal(t, 29, p2)

	p1, p2 = day12(d12)
	assert.Equal(t, 330, p1)
	assert.Equal(t, 321, p2)
}
