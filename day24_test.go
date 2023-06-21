package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay24(t *testing.T) {
	input := `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
`
	{
		p1, p2 := day24(input)
		assert.Equal(t, 18, p1)
		assert.Equal(t, 54, p2)
	}
	{
		p1, p2 := day24(d24)
		assert.Equal(t, 251, p1)
		assert.Equal(t, 758, p2)
	}
}
