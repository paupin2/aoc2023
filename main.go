package main

import (
	"fmt"

	"github.com/paupin2/aoc2023/day1"
	"github.com/paupin2/aoc2023/day2"
	"github.com/paupin2/aoc2023/day3"
	"github.com/paupin2/aoc2023/day4"
)

func main() {
	type runner func(int) int
	runners := []runner{day1.Run, day2.Run, day3.Run, day4.Run}
	for idx, r := range runners {
		day := idx + 1
		fmt.Printf("day %d part 1: %d\n", day, r(1))
		fmt.Printf("day %d part 2: %d\n", day, r(2))
	}
}
