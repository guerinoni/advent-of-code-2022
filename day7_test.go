package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay7(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	p1, p2 := day7(input)
	assert.Equal(t, 95437, p1)
	assert.Equal(t, 24933642, p2)

	p1, p2 = day7(d7)
	assert.Equal(t, 1582412, p1)
	assert.Equal(t, 3696336, p2)
}
