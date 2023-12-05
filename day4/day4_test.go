package day4

import (
	"testing"

	"github.com/paupin2/aoc2023/common"
	"github.com/paupin2/aoc2023/daydata"
)

func Test_parseCard(t *testing.T) {
	common.DeepEqual(t, parseCard(daydata.Four.Read(1)), []card{
		{num: 1, win: []int{41, 48, 83, 86, 17}, have: []int{83, 86, 6, 31, 17, 9, 48, 53}},
		{num: 2, win: []int{13, 32, 20, 16, 61}, have: []int{61, 30, 68, 82, 17, 32, 24, 19}},
		{num: 3, win: []int{1, 21, 53, 59, 44}, have: []int{69, 82, 63, 72, 16, 21, 14, 1}},
		{num: 4, win: []int{41, 92, 73, 84, 69}, have: []int{59, 84, 76, 51, 58, 5, 54, 83}},
		{num: 5, win: []int{87, 83, 26, 28, 32}, have: []int{88, 30, 70, 12, 93, 22, 82, 36}},
		{num: 6, win: []int{31, 18, 13, 56, 72}, have: []int{74, 77, 10, 23, 35, 67, 36, 11}},
	})
}

func Test_card_points(t *testing.T) {
	common.Equals(t, sumPile(parseCard(daydata.Four.Read(1))), 13)
}

func TestRun(t *testing.T) {
	common.Equals(t, Run(1), 21919)
	common.Equals(t, Run(2), 9881048)
}

func Test_countCards(t *testing.T) {
	common.Equals(t, countCards(parseCard(daydata.Four.Read(1))), 30)
}
