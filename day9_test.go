package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay9(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

	p1, p2 := day9(input)

	assert.Equal(t, 13, p1)
	assert.Equal(t, 1, p2)

	input = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

	_, p2 = day9(input)
	assert.Equal(t, 36, p2)

	p1, p2 = day9(d9)
	assert.Equal(t, 6745, p1)
	assert.Equal(t, 2793, p2)
}
