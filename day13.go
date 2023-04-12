package main

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed input/day13
var d13 string

type packet struct {
	parent *packet
	value  int
	items  []*packet
}

func day13(input string) (int, int) {
	pairs := strings.Split(input, "\n\n")

	var packets []packet
	packets = append(packets, parsePacket("[[2]]"))
	packets = append(packets, parsePacket("[[6]]"))

	sumOfIndex := 0
	for idx, pair := range pairs {
		lines := strings.Split(pair, "\n")
		p1 := parsePacket(lines[0])
		p2 := parsePacket(lines[1])
		packets = append(packets, p1, p2)
		if whichGreater(p1, p2) == right {
			sumOfIndex += idx + 1
		}
	}

	sort.Slice(packets, func(i, j int) bool {
		return whichGreater(packets[i], packets[j]) == right
	})

	idxSentinel2 := -1
	idxSentinel6 := -1
	for i := 0; i < len(packets); i++ {
		if idxSentinel6 != -1 && idxSentinel2 != -1 {
			break
		}

		if len(packets[i].items) == 1 {
			if len(packets[i].items[0].items) == 1 {
				if packets[i].items[0].items[0].value == 2 {
					idxSentinel2 = i + 1
				}

				if packets[i].items[0].items[0].value == 6 {
					idxSentinel6 = i + 1
				}
			}
		}
	}

	return sumOfIndex, idxSentinel2 * idxSentinel6
}

const (
	left = -1 + iota
	equal
	right
)

func whichGreater(lhs packet, rhs packet) int {
	if len(lhs.items) == 0 && len(rhs.items) == 0 {
		if lhs.value == rhs.value {
			return equal
		}

		if lhs.value > rhs.value {
			return left
		} else {
			return right
		}
	}

	if lhs.value >= 0 {
		return whichGreater(packet{nil, -1, []*packet{&lhs}}, rhs)
	}

	if rhs.value >= 0 {
		return whichGreater(lhs, packet{nil, -1, []*packet{&rhs}})
	}

	var i int
	for i = 0; i < len(lhs.items) && i < len(rhs.items); i++ {
		if ordered := whichGreater(*lhs.items[i], *rhs.items[i]); ordered != equal {
			return ordered
		}
	}

	if len(lhs.items) > len(rhs.items) {
		return left
	} else if len(lhs.items) < len(rhs.items) {
		return right
	}

	return equal
}

func parsePacket(line string) packet {
	root := packet{nil, -1, []*packet{}}
	current := &root

	var currentNumber string
	for _, c := range line {
		switch c {
		case '[':
			newpacket := packet{current, -1, []*packet{}}
			current.items = append(current.items, &newpacket)
			current = &newpacket
		case ']':
			if len(currentNumber) > 0 {
				current.value = mustAtoi(currentNumber)
				currentNumber = ""
			}
			current = current.parent
		case ',':
			if len(currentNumber) > 0 {
				current.value = mustAtoi(currentNumber)
				currentNumber = ""
			}
			current = current.parent
			current = newPacket(current)
		default:
			currentNumber += string(c)
		}
	}
	return root
}

func newPacket(current *packet) *packet {
	newPacket := packet{current, -1, []*packet{}}
	current.items = append(current.items, &newPacket)
	current = &newPacket
	return &newPacket
}
