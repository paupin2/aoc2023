package day1

import "testing"

func TestCalibrate(t *testing.T) {
	actual := Calibrate("testinput.txt")
	if actual != 142 {
		t.Fatalf("got %d instead", actual)
	}
}
