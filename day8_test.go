package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay8(t *testing.T) {
	input := `30373
25512
65332
33549
35390`

	p1, p2 := day8(input)
	assert.Equal(t, 21, p1)
	assert.Equal(t, 8, p2)

	p1, p2 = day8(d8)
	assert.Equal(t, 1840, p1)
	assert.Equal(t, 405769, p2)
}
