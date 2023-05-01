package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay18(t *testing.T) {
	input := `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
`
	{
		p1, p2 := day18(input)
		assert.Equal(t, 64, p1)
		assert.Equal(t, 58, p2)
	}
	{
		p1, p2 := day18(d18)
		assert.Equal(t, 3662, p1)
		assert.Equal(t, 2060, p2)
	}
}
