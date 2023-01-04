package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {
	input :=
		`A Y
B X
C Z`

	p1, p2 := day2(input)
	assert.Equal(t, 15, p1)
	assert.Equal(t, 12, p2)

	p1, p2 = day2(d2)
	assert.Equal(t, 10595, p1)
	assert.Equal(t, 9541, p2)
}
