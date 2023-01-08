package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay5(t *testing.T) {
	input := `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

	p1, p2 := day5(input)
	assert.Equal(t, "CMZ", p1)
	assert.Equal(t, "MCD", p2)

	p1, p2 = day5(d5)
	assert.Equal(t, "ZWHVFWQWW", p1)
	assert.Equal(t, "HZFZCCWWV", p2)
}
