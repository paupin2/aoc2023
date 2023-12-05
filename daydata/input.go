package daydata

import (
	"bufio"
	"embed"
	"fmt"
	"strings"

	"github.com/paupin2/aoc2023/common"
)

type Day int

const (
	One Day = iota + 1
	Two
	Three
	Four
)

//go:embed *.txt
var inputfs embed.FS

func lines(fs embed.FS, filename string) (lines []string) {
	f := common.Must(fs.Open(filename))
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if line := strings.TrimSpace(s.Text()); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func (d Day) Read(i int) []string {
	return lines(inputfs, fmt.Sprintf("day%dinput%d.txt", d, i))
}
