package main

import (
	_ "embed"
	"strings"
)

//go:embed input/day2
var d2 string

func day2(input string) (int, int) {
	lines := strings.Split(input, "\n")

	points1 := 0
	points2 := 0
	for _, line := range lines {

		chars := strings.Split(line, " ")
		if len(chars) == 1 {
			continue
		}

		points1 += calculatePointOfMyChoice(chars[1])
		points1 += calculateRound(chars[0], chars[1])

		switch chars[1] {
		case "Y":
			chars[1] = chars[0]
		case "X":
			chars[1] = convertToLoose(chars[0])
		case "Z":
			chars[1] = convertToWin(chars[0])
		}

		points2 += calculatePointOfMyChoice(chars[1])
		points2 += calculateRound(chars[0], chars[1])
	}

	return points1, points2
}

func convertToWin(ch string) string {
	if ch == "A" {
		return "Y"
	}
	if ch == "B" {
		return "Z"
	}
	if ch == "C" {
		return "X"
	}
	return ""
}

func convertToLoose(ch string) string {
	if ch == "A" {
		return "Z"
	}
	if ch == "B" {
		return "X"
	}
	if ch == "C" {
		return "Y"
	}

	return ""
}

func calculateRound(ch1 string, ch2 string) int {
	if isSissor(ch1) && isRock(ch2) || isRock(ch1) && isPaper(ch2) || isPaper(ch1) && isSissor(ch2) {
		return 6
	}

	if isRock(ch1) && isRock(ch2) || isPaper(ch1) && isPaper(ch2) || isSissor(ch1) && isSissor(ch2) {
		return 3
	}
	return 0
}

func calculatePointOfMyChoice(ch string) int {
	if isRock(ch) {
		return 1
	}
	if isPaper(ch) {
		return 2
	}
	if isSissor(ch) {
		return 3
	}
	return 0
}

func isRock(ch string) bool {
	return ch == "X" || ch == "A"
}

func isPaper(ch string) bool {
	return ch == "Y" || ch == "B"
}

func isSissor(ch string) bool {
	return ch == "Z" || ch == "C"
}
