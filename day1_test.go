package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	p1, p2 := day1(input)
	assert.Equal(t, 24000, p1)
	assert.Equal(t, 45000, p2)

	p1, p2 = day1(d1)
	assert.Equal(t, 70720, p1)
	assert.Equal(t, 207148, p2)
}
