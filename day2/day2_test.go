package day2

import (
	"testing"

	"github.com/paupin2/aoc2023/common"
)

func Test_parse(t *testing.T) {
	common.DeepEqual(t, parse("calibration.txt"), []game{
		{id: 1, hands: []hand{{blue: 3, red: 4}, {red: 1, green: 2, blue: 6}, {green: 2}}},
		{id: 2, hands: []hand{{blue: 1, green: 2}, {green: 3, blue: 4, red: 1}, {green: 1, blue: 1}}},
		{id: 3, hands: []hand{{green: 8, blue: 6, red: 20}, {blue: 5, red: 4, green: 13}, {green: 5, red: 1}}},
		{id: 4, hands: []hand{{green: 1, red: 3, blue: 6}, {green: 3, red: 6}, {green: 3, blue: 15, red: 14}}},
		{id: 5, hands: []hand{{red: 6, blue: 1, green: 3}, {blue: 2, red: 1, green: 2}}},
	})
}

func Test_underMax(t *testing.T) {
	games := parse("calibration.txt")
	max := hand{red: 12, green: 13, blue: 14}
	common.Equals(t, underMax(games, max), 8)
}

func Test_game_findMin(t *testing.T) {
	games := parse("calibration.txt")
	common.Equals(t, games[0].findMin(), hand{red: 4, green: 2, blue: 6})
	common.Equals(t, games[0].findMin().power(), 48)
	common.Equals(t, games[1].findMin().power(), 12)
	common.Equals(t, games[2].findMin().power(), 1560)
	common.Equals(t, games[3].findMin().power(), 630)
	common.Equals(t, games[4].findMin().power(), 36)
}
