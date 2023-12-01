package common

import (
	"bufio"
	"os"
	"strings"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func FileLines(filename string) (lines []string) {
	f := Must(os.Open(filename))
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if line := strings.TrimSpace(s.Text()); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
