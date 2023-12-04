package common

import (
	"bufio"
	"embed"
	"reflect"
	"strings"
	"testing"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func FileLines(fs embed.FS, filename string) (lines []string) {
	f := Must(fs.Open(filename))
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if line := strings.TrimSpace(s.Text()); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func Equals[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		t.Fatalf("expected %v but got %v instead", expected, actual)
	}
}

func SliceEquals[T comparable](t *testing.T, actual, expected []T) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("expected len %d but got %d instead", len(expected), len(actual))
	} else {
		for i := range actual {
			Equals(t, actual[i], expected[i])
		}
	}
}

func DeepEqual[T any](t *testing.T, actual, expected T) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected:\n%+v\nbut got:\n%+v\ninstead", expected, actual)
	}
}
