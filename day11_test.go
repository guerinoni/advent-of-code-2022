package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11(t *testing.T) {
	input := `monnkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

monnkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

monnkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

monnkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`

	p1, p2 := day11(input)
	assert.Equal(t, 10605, p1)
	assert.Equal(t, 2713310158, p2)

	p1, p2 = day11(d11)
	assert.Equal(t, 54253, p1)
	assert.Equal(t, 13119526120, p2)
}
