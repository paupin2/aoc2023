package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/paupin2/aoc2023/daydata"
)

type card struct {
	num  int
	win  []int
	have []int
}

func (c card) isWinning(nr int) bool {
	for _, w := range c.win {
		if w == nr {
			return true
		}
	}
	return false
}

func (c card) wins() (wins int) {
	for _, h := range c.have {
		if c.isWinning(h) {
			wins++
		}
	}
	return
}

func (c card) points() int {
	if w := c.wins(); w > 0 {
		return int(math.Pow(2, float64(c.wins()-1)))
	}
	return 0
}

func parseCard(lines []string) (cards []card) {
	fields := func(s string) (fs []int) {
		for _, n := range strings.Fields(s) {
			i, _ := strconv.Atoi(n)
			fs = append(fs, i)
		}
		return
	}

	for _, line := range lines {
		var c card
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		cardname, rest, _ := strings.Cut(line, ":")
		fmt.Sscanf(cardname, "Card %d", &c.num)
		winning, have, _ := strings.Cut(rest, "|")
		c.win = fields(winning)
		c.have = fields(have)
		cards = append(cards, c)
	}
	return
}

func sumPile(cards []card) (sum int) {
	for _, c := range cards {
		sum += c.points()
	}
	return
}

func countCards(cards []card) int {
	copies := make([]int, len(cards))
	for i, card := range cards {
		copies[i]++
		wins := card.wins()
		for ; wins > 0; wins-- {
			copies[i+wins] += copies[i]
		}
	}

	count := 0
	for _, n := range copies {
		count += n
	}
	return count
}

func Run(part int) int {
	switch part {
	case 1:
		return sumPile(parseCard(daydata.Four.Read(2)))

	case 2:
		return countCards(parseCard(daydata.Four.Read(2)))

	default:
		panic("bad part")
	}
}
