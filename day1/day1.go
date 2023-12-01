package day1

import (
	"strconv"
	"unicode"

	"github.com/paupin2/aoc2023/common"
)

func firstLast(line string) int {
	var first, last rune
	for _, r := range line {
		if unicode.IsDigit(r) {
			first = r
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if r := rune(line[i]); unicode.IsDigit(r) {
			last = r
			break
		}
	}
	n, _ := strconv.Atoi(string([]rune{first, last}))
	return n
}

func Calibrate(filename string) int {
	var sum int
	for _, line := range common.FileLines(filename) {
		sum += firstLast(line)
	}
	return sum
}
