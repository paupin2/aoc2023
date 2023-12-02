package day1

import (
	"fmt"
	"testing"

	"github.com/paupin2/aoc2023/common"
)

func equals[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		t.Fatalf("expected %v but got %v instead", expected, actual)
	}
}

func TestSum(t *testing.T) {
	equals(t, Sum("part1test.txt", false), 142)
	for _, l := range common.FileLines("part2test.txt") {
		fmt.Printf("%q first:%d (%d) last:%d (%d)\n", l,
			findNumber(l, forward, false), findNumber(l, forward, true),
			findNumber(l, backwards, false), findNumber(l, backwards, true),
		)
	}
	equals(t, Sum("part2test.txt", true), 281)
}

func Test_firstNumber(t *testing.T) {
	equals(t, findNumber("hello1", forward, false), 1)
	equals(t, findNumber("hello1", forward, true), 1)
	equals(t, findNumber("hellotwo1", forward, false), 1)
	equals(t, findNumber("hellotwo1", forward, true), 2)
	equals(t, findNumber("hellotwo1", backwards, true), 1)
}
