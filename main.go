package main

import (
	"fmt"

	"github.com/paupin2/aoc2023/day1"
	"github.com/paupin2/aoc2023/day2"
	"github.com/paupin2/aoc2023/day3"
)

func main() {
	fmt.Printf("day 1 part 1: %d\n", day1.Run(false))
	fmt.Printf("day 1 part 2: %d\n", day1.Run(true))
	fmt.Printf("day 2 part 1: %d\n", day2.Run(1))
	fmt.Printf("day 2 part 2: %d\n", day2.Run(2))
	fmt.Printf("day 3 part 1: %d\n", day3.Run(1))
	fmt.Printf("day 3 part 2: %d\n", day3.Run(2))
}
