package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16(t *testing.T) {
	input := `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`
	{
		p1, p2 := day16(input)
		assert.Equal(t, 1651, p1)
		assert.Equal(t, 1707, p2)
	}
	{
		p1, p2 := day16(d16)
		assert.Equal(t, 1873, p1)
		assert.Equal(t, 2425, p2)
	}
}
