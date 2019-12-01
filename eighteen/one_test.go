package eighteen_test

import (
	"testing"

	"github.com/paunik/adventofcode/eighteen"
)

var chronalCalibrationTestCaseOne = []struct {
	name string
	in   []int
	out  int
}{
	{"first", []int{+1, -2, +3, +1}, 3},
	{"second", []int{+1, +1, +1}, 3},
	{"third", []int{+1, +1, -2}, 0},
	{"fourth", []int{-1, -2, -3}, -6},
}

var chronalCalibrationTestCaseTwo = []struct {
	name string
	in   []int
	out  int
}{
	{"first", []int{+1, -2, +3, +1}, 2},
	{"second", []int{+1, -1}, 0},
	{"third", []int{+3, +3, +4, -2, -4}, 10},
	{"fourth", []int{-6, +3, +8, +5, -6}, 5},
	{"fift", []int{+7, +7, -2, -7, -4}, 14},
}

func TestChronalCalibration(t *testing.T) {
	for _, tt := range chronalCalibrationTestCaseOne {
		t.Run(tt.name, func(t *testing.T) {
			res := eighteen.ChronalCalibration(tt.in)
			if res != tt.out {
				t.Errorf("got %d, want %d", res, tt.out)
			}
		})
	}
}

func TestChronalCalibrationDouble(t *testing.T) {
	for _, tt := range chronalCalibrationTestCaseTwo {
		t.Run(tt.name, func(t *testing.T) {
			res := eighteen.ChronalCalibrationDouble(tt.in)
			if res != tt.out {
				t.Errorf("got %d, want %d", res, tt.out)
			}
		})
	}
}

func TestDoIt(t *testing.T) {
	res := eighteen.DoIt()
	if res != 435 {
		t.Errorf("Sum is: %d.", res)
	}
}

func TestDoItSecond(t *testing.T) {
	res := eighteen.DoItSecond()
	if res != 245 {
		t.Errorf("Sum is: %d.", res)
	}
}
