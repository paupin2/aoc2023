package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/paupin2/aoc2023/daydata"
)

type hand struct {
	red, green, blue int
}

func (h hand) under(max hand) bool {
	return h.red <= max.red && h.green <= max.green && h.blue <= max.blue
}

func (h hand) power() int {
	return h.red * h.green * h.blue
}

type game struct {
	id    int
	hands []hand
}

func (g game) findMin() hand {
	var min hand
	for _, h := range g.hands {
		maxInt(&min.red, h.red)
		maxInt(&min.green, h.green)
		maxInt(&min.blue, h.blue)
	}
	return min
}

func parse(lines []string) (games []game) {
	for _, l := range lines {
		gamepart, handspart, _ := strings.Cut(l, ":")
		gamefields := strings.Fields(gamepart)
		var g game
		g.id, _ = strconv.Atoi(gamefields[1])

		for _, handstr := range strings.Split(handspart, ";") {
			var h hand
			for _, grabstr := range strings.Split(handstr, ",") {
				var count int
				var color string
				fmt.Sscanf(grabstr, "%d %s", &count, &color)
				switch color {
				case "red":
					h.red += count
				case "green":
					h.green += count
				case "blue":
					h.blue += count
				}
			}
			g.hands = append(g.hands, h)
		}
		games = append(games, g)
	}
	return
}

func maxInt(dest *int, src int) {
	if src > *dest {
		*dest = src
	}
}

func underMax(games []game, max hand) (idSum int) {
	for _, g := range games {
		under := true
		for _, h := range g.hands {
			if !h.under(max) {
				under = false
			}
		}
		if under {
			idSum += g.id
		}
	}
	return
}

func Run(part int) int {
	switch part {
	case 1:
		games := parse(daydata.Two.Read(2))
		max := hand{red: 12, green: 13, blue: 14}
		return underMax(games, max)

	case 2:
		games := parse(daydata.Two.Read(2))
		austin := 0
		for _, g := range games {
			austin += g.findMin().power()
		}
		return austin

	default:
		panic("bad part")
	}
}
