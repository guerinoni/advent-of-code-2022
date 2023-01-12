package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay6(t *testing.T) {
	input := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

	p1, p2 := day6(input)
	assert.Equal(t, 7, p1)
	assert.Equal(t, 19, p2)

	p1, p2 = day6(d6)
	assert.Equal(t, 1760, p1)
	assert.Equal(t, 2974, p2)
}
