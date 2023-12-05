package day1

import (
	"fmt"
	"testing"

	"github.com/paupin2/aoc2023/common"
	"github.com/paupin2/aoc2023/daydata"
)

func TestSum(t *testing.T) {
	common.Equals(t, sum(daydata.One.Read(1), false), 142)
	for _, l := range daydata.One.Read(2) {
		fmt.Printf("%q first:%d (%d) last:%d (%d)\n", l,
			findNumber(l, forward, false), findNumber(l, forward, true),
			findNumber(l, backwards, false), findNumber(l, backwards, true),
		)
	}
	common.Equals(t, sum(daydata.One.Read(2), true), 281)
}

func Test_firstNumber(t *testing.T) {
	common.Equals(t, findNumber("hello1", forward, false), 1)
	common.Equals(t, findNumber("hello1", forward, true), 1)
	common.Equals(t, findNumber("hellotwo1", forward, false), 1)
	common.Equals(t, findNumber("hellotwo1", forward, true), 2)
	common.Equals(t, findNumber("hellotwo1", backwards, true), 1)
}

func Test_Run(t *testing.T) {
	common.Equals(t, Run(1), 54597)
	common.Equals(t, Run(2), 54504)
}
