package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay4(t *testing.T) {
	input :=
		`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	p1, p2 := day4(input)
	assert.Equal(t, 2, p1)
	assert.Equal(t, 4, p2)

	p1, p2 = day4(d4)
	assert.Equal(t, 595, p1)
	assert.Equal(t, 952, p2)
}
