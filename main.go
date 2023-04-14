package main

import "fmt"

func main() {
	fmt.Println("Advent of code 2022!")

	d1p1, d1p2 := day1(d1)
	fmt.Println("day 1 part 1 -> ", d1p1)
	fmt.Println("day 1 part 2 -> ", d1p2)

	d2p1, d2p2 := day2(d2)
	fmt.Println("day 2 part 1 -> ", d2p1)
	fmt.Println("day 2 part 2 -> ", d2p2)

	d3p1, d3p2 := day3(d3)
	fmt.Println("day 3 part 1 -> ", d3p1)
	fmt.Println("day 3 part 2 -> ", d3p2)

	d4p1, d4p2 := day4(d4)
	fmt.Println("day 4 part 1 -> ", d4p1)
	fmt.Println("day 4 part 2 -> ", d4p2)

	d5p1, d5p2 := day5(d5)
	fmt.Println("day 5 part 1 -> ", d5p1)
	fmt.Println("day 5 part 2 -> ", d5p2)

	d6p1, d6p2 := day6(d6)
	fmt.Println("day 6 part 1 -> ", d6p1)
	fmt.Println("day 6 part 2 -> ", d6p2)

	d7p1, d7p2 := day7(d7)
	fmt.Println("day 7 part 1 -> ", d7p1)
	fmt.Println("day 7 part 2 -> ", d7p2)

	d8p1, d8p2 := day8(d8)
	fmt.Println("day 8 part 1 -> ", d8p1)
	fmt.Println("day 8 part 2 -> ", d8p2)

	d9p1, d9p2 := day9(d9)
	fmt.Println("day 9 part 1 -> ", d9p1)
	fmt.Println("day 9 part 2 -> ", d9p2)

	d10p1, d10p2 := day10(d10)
	fmt.Println("day 10 part 1 -> ", d10p1)
	fmt.Println("day 10 part 2 -> ", d10p2)

	d11p1, d11p2 := day11(d11)
	fmt.Println("day 11 part 1 -> ", d11p1)
	fmt.Println("day 11 part 2 -> ", d11p2)

	d12p1, d12p2 := day12(d12)
	fmt.Println("day 12 part 1 -> ", d12p1)
	fmt.Println("day 12 part 2 -> ", d12p2)

	d13p1, d13p2 := day13(d13)
	fmt.Println("day 13 part 1 -> ", d13p1)
	fmt.Println("day 13 part 2 -> ", d13p2)

	d14p1, d14p2 := day14(d14)
	fmt.Println("day 14 part 1 -> ", d14p1)
	fmt.Println("day 14 part 2 -> ", d14p2)

	d15p1, d15p2 := day15(d15, 2000000)
	fmt.Println("day 15 part 1 -> ", d15p1)
	fmt.Println("day 15 part 2 -> ", d15p2)
}
