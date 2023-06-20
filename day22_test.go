package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay22(t *testing.T) {
	//input := `        ...#
	//    .#..
	//    #...
	//    ....
	//...#.......#
	//........#...
	//..#....#....
	//..........#.
	//        ...#....
	//        .....#..
	//        .#......
	//        ......#.
	//
	//10R5L5R10L4R5L5
	//`
	//{
	//	p1, p2 := day22(input)
	//	assert.Equal(t, 6032, p1)
	//	assert.Equal(t, 5031, p2)
	//}
	{
		p1, p2 := day22(d22)
		assert.Equal(t, 89224, p1)
		assert.Equal(t, 136182, p2)
	}
}
