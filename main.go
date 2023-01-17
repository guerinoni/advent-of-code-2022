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
}
