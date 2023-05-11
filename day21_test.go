package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay21(t *testing.T) {
	input := `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
`
	{
		p1, p2 := day21(input)
		assert.Equal(t, 152, p1)
		assert.Equal(t, 301, p2)
	}
	{
		p1, p2 := day21(d21)
		assert.Equal(t, 118565889858886, p1)
		assert.Equal(t, 3032671800353, p2)
	}
}
