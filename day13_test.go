package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13(t *testing.T) {
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`
	{
		p1, p2 := day13(input)
		assert.Equal(t, 13, p1)
		assert.Equal(t, 140, p2)
	}
	{
		p1, p2 := day13(d13)
		assert.Equal(t, 5198, p1)
		assert.Equal(t, 22344, p2)
	}
}

func Test_areOrdered(t *testing.T) {
	type args struct {
		lhs string
		rhs string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{lhs: "[1,1,3,1,1]", rhs: "[1,1,5,1,1]"},
			want: right,
		},
		{
			name: "2",
			args: args{lhs: "[[1],[2,3,4]]", rhs: "[[1],4]"},
			want: right,
		},
		{
			name: "3",
			args: args{lhs: "[9]", rhs: "[[8,7,6]]"},
			want: left,
		},
		{
			name: "4",
			args: args{lhs: "[[4,4],4,4]", rhs: "[[4,4],4,4,4]"},
			want: right,
		},
		{
			name: "5",
			args: args{lhs: "[7,7,7,7]", rhs: "[7,7,7]"},
			want: left,
		},
		{
			name: "6",
			args: args{lhs: "[]", rhs: "[3]"},
			want: right,
		},
		{
			name: "7",
			args: args{lhs: "[[[]]]", rhs: "[[]]"},
			want: left,
		},
		{
			name: "8",
			args: args{lhs: "[1,[2,[3,[4,[5,6,7]]]],8,9]", rhs: "[1,[2,[3,[4,[5,6,0]]]],8,9]"},
			want: left,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lhs := parsePacket(tt.args.lhs)
			rhs := parsePacket(tt.args.rhs)
			assert.Equalf(t, tt.want, whichGreater(lhs, rhs), "whichGreater(%v, %v)", tt.args.lhs, tt.args.rhs)
		})
	}
}
