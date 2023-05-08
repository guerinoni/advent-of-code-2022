package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay19(t *testing.T) {
	input := `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`
	{
		p1, p2 := day19(input)
		assert.Equal(t, 33, p1)
		assert.Equal(t, 56, p2)
	}
	{
		p1, p2 := day19(d19)
		assert.Equal(t, 1766, p1)
		assert.Equal(t, 0, p2)
	}
}
