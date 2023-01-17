package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay3(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	p1, p2 := day3(input)

	assert.Equal(t, 157, p1)
	assert.Equal(t, 70, p2)

	p1, p2 = day3(d3)
	assert.Equal(t, 7446, p1)
	assert.Equal(t, 2646, p2)
}
