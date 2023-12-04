package day1

import (
	"fmt"
	"testing"

	"github.com/paupin2/aoc2023/common"
)

func TestSum(t *testing.T) {
	common.Equals(t, sum("part1test.txt", false), 142)
	for _, l := range common.FileLines(inputfs, "part2test.txt") {
		fmt.Printf("%q first:%d (%d) last:%d (%d)\n", l,
			findNumber(l, forward, false), findNumber(l, forward, true),
			findNumber(l, backwards, false), findNumber(l, backwards, true),
		)
	}
	common.Equals(t, sum("part2test.txt", true), 281)
}

func Test_firstNumber(t *testing.T) {
	common.Equals(t, findNumber("hello1", forward, false), 1)
	common.Equals(t, findNumber("hello1", forward, true), 1)
	common.Equals(t, findNumber("hellotwo1", forward, false), 1)
	common.Equals(t, findNumber("hellotwo1", forward, true), 2)
	common.Equals(t, findNumber("hellotwo1", backwards, true), 1)
}
