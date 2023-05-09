package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay20(t *testing.T) {
	input := `1
2
-3
3
-2
0
4`
	{
		p1, p2 := day20(input)
		assert.Equal(t, int64(3), p1)
		assert.Equal(t, int64(1623178306), p2)
	}
	{
		p1, p2 := day20(d20)
		assert.Equal(t, int64(7004), p1)
		assert.Equal(t, int64(17200008919529), p2)
	}
}
