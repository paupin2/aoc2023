package day3

import (
	"testing"

	"github.com/paupin2/aoc2023/common"
	"github.com/paupin2/aoc2023/daydata"
)

func Test_schematic_findNumbers(t *testing.T) {
	common.DeepEqual(t,
		load(daydata.Three.Read(1)).findNumbers(),
		[]number{
			{value: 467, isPartNumber: true, gears: []coord{{1, 3}}},
			{value: 114, isPartNumber: false},
			{value: 35, isPartNumber: true, gears: []coord{{1, 3}}},
			{value: 633, isPartNumber: true},
			{value: 617, isPartNumber: true, gears: []coord{{4, 3}}},
			{value: 58, isPartNumber: false},
			{value: 592, isPartNumber: true},
			{value: 755, isPartNumber: true, gears: []coord{{8, 5}}},
			{value: 664, isPartNumber: true},
			{value: 598, isPartNumber: true, gears: []coord{{8, 5}}},
		},
	)
}

func Test_sumParts(t *testing.T) {
	common.Equals(t, sumParts(load(daydata.Three.Read(1)).findNumbers()), 4361)
}

func Test_sumGearRatios(t *testing.T) {
	common.Equals(t, sumGearRatios(load(daydata.Three.Read(1)).findNumbers()), 467835)
}

func TestRun(t *testing.T) {
	common.Equals(t, Run(1), 539637)
	common.Equals(t, Run(2), 82818007)
}
