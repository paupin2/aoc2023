package day3

import (
	"embed"
	"strconv"

	"github.com/paupin2/aoc2023/common"
)

type coord struct{ row, col int }

type number struct {
	value        int
	isPartNumber bool
	gears        []coord
}

type schematic [][]byte

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func sumParts(parts []number) (sum int) {
	for _, p := range parts {
		if p.isPartNumber {
			sum += p.value
		}
	}
	return
}

func sumGearRatios(parts []number) (sum int) {
	gears := map[coord][]int{}
	for _, p := range parts {
		for _, g := range p.gears {
			gears[g] = append(gears[g], p.value)
		}
	}
	for _, values := range gears {
		if len(values) == 2 {
			ratio := values[0] * values[1]
			sum += ratio
		}
	}

	return
}

func (s schematic) get(row, col int) byte {
	if row >= 0 && row < len(s) && col > 0 && col < len(s[row]) {
		return s[row][col]
	}
	return 0
}

func (s schematic) isPart(row, col int) bool {
	b := s.get(row, col)
	return b != 0 && b != '.' && !isDigit(b)
}

func (s schematic) findNumbers() (found []number) {
	addPart := func(n *number, r, c int) {
		if s.isPart(r, c) {
			n.isPartNumber = true
			if s.get(r, c) == '*' {
				n.gears = append(n.gears, coord{row: r, col: c})
			}
		}
	}

	for row, line := range s {
		for col := 0; col < len(line); col++ {
			if !isDigit(line[col]) {
				continue
			}

			start := col
			for col+1 < len(line) && isDigit(line[col+1]) {
				col++
			}
			end := col

			text := line[start : end+1]
			var n number
			n.value, _ = strconv.Atoi(string(text))

			// top, bottom rows
			for c := start - 1; c <= end+1; c++ {
				addPart(&n, row-1, c)
			}
			for c := start - 1; c <= end+1; c++ {
				addPart(&n, row+1, c)
			}

			// left, right
			addPart(&n, row, start-1)
			addPart(&n, row, end+1)
			found = append(found, n)
		}
	}
	return
}

//go:embed *.txt
var inputfs embed.FS

func load(filename string) schematic {
	var s schematic
	for _, l := range common.FileLines(inputfs, filename) {
		s = append(s, []byte(l))
	}
	return s
}

func Run(part int) int {
	switch part {
	case 1:
		return sumParts(load("input.txt").findNumbers())

	case 2:
		return sumGearRatios(load("input.txt").findNumbers())

	default:
		panic("nad part")
	}
}
