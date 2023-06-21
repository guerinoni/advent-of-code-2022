package main

import (
	_ "embed"
	"math"
	"strings"
)

//go:embed input/day25
var d25 string

func day25(input string) string {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	sum := 0

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			var num int
			switch line[len(line)-1-i] {
			case '2':
				num = 2
			case '1':
				num = 1
			case '0':
				num = 0
			case '-':
				num = -1
			case '=':
				num = -2
			}
			sum += int(math.Pow(5, float64(i))) * num
		}
	}

	ret := ""
	for sum > 0 {
		switch sum % 5 {
		case 2:
			ret = "2" + ret
		case 1:
			ret = "1" + ret
		case 0:
			ret = "0" + ret
		case 4:
			ret = "-" + ret
			sum += 2
		case 3:
			ret = "=" + ret
			sum += 3
		}
		sum /= 5
	}

	return ret
}
