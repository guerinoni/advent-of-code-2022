package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay25(t *testing.T) {
	input := `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`
	{
		p1 := day25(input)
		assert.Equal(t, "2=-1=0", p1)
	}
	{
		p1 := day25(d25)
		assert.Equal(t, "2---0-1-2=0=22=2-011", p1)
	}
}
