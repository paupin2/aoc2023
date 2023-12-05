package day1

import (
	"strings"

	"github.com/paupin2/aoc2023/daydata"
)

var numberNames = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

type direction int

const (
	forward direction = iota + 1
	backwards
)

func findNumber(line string, dir direction, useNames bool) int {
	start, delta, end := 0, 1, len(line)
	if dir == backwards {
		start, delta, end = len(line)-1, -1, -1
	}

	for li := start; li != end; li += delta {
		if r := line[li]; r >= '0' && r <= '9' {
			return int(r - '0')
		}
		if !useNames {
			continue
		}
		for ni := range numberNames {
			if strings.HasPrefix(line[li:], numberNames[ni]) {
				return ni + 1
			}
		}
	}
	return 0
}

func firstLast(line string, useNames bool) int {
	return findNumber(line, forward, useNames)*10 +
		findNumber(line, backwards, useNames)
}

func sum(lines []string, useNames bool) int {
	var sum int
	for _, line := range lines {
		sum += firstLast(line, useNames)
	}
	return sum
}

func Run(part int) int {
	switch part {
	case 1:
		return sum(daydata.One.Read(3), false)
	case 2:
		return sum(daydata.One.Read(3), true)
	default:
		panic("bad part")
	}
}
