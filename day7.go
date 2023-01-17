package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input/day7
var d7 string

type Node struct {
	name   string
	parent *Node
	size   int
	kids   []*Node
}

func day7(input string) (int, int) {
	lines := strings.Split(input, "\n")

	var root Node
	var curr *Node
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "$ ls") || strings.Contains(line, "dir ") {
			continue
		}

		if strings.Contains(line, "$ cd") {
			arg := strings.Split(line, "cd ")
			dirName := arg[len(arg)-1]

			if dirName == "/" {
				root = Node{name: dirName}
				curr = &root
			} else if dirName == ".." {
				curr = curr.parent
			} else {
				curr.kids = append(
					curr.kids,
					&Node{name: dirName, parent: curr},
				)
				curr = curr.kids[len(curr.kids)-1]
			}
			continue
		}

		fileInfo := strings.Split(line, " ")
		size, err := strconv.Atoi(fileInfo[0])
		if err != nil {
			panic(err)
		}
		curr.size += size
		backportSize(curr, size)
	}

	currentUnusedSpace := 70000000 - root.size
	smallestEnough := root.size
	return sumOfDirWithMaxSize(&root, 100000), findSizeDirToDelete(&root, currentUnusedSpace, &smallestEnough)
}

func backportSize(n *Node, size int) {
	if n.parent == nil {
		return
	}
	n.parent.size += size
	backportSize(n.parent, size)
}

func sumOfDirWithMaxSize(n *Node, target int) (tot int) {
	if n == nil {
		return 0
	}

	for _, n := range n.kids {
		if n.size <= target {
			tot += n.size
		}

		tot += sumOfDirWithMaxSize(n, target)
	}

	return
}

func findSizeDirToDelete(n *Node, currentUnusedSpace int, smallest *int) int {
	if n == nil {
		return *smallest
	}

	for _, n := range n.kids {
		if n.size > (30000000-currentUnusedSpace) && n.size < *smallest {
			*smallest = n.size
		}
		findSizeDirToDelete(n, currentUnusedSpace, smallest)
	}

	return *smallest
}
