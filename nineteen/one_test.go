package nineteen_test

import (
	"testing"

	"github.com/paunik/adventofcode/nineteen"
)

var firstFueltTestCases = []struct {
	name string
	in   int
	out  int
}{
	{"first", 12, 2},
	{"second", 14, 2},
	{"third", 1969, 654},
	{"fourth", 100756, 33583},
}

var secondFueltTestCases = []struct {
	name string
	in   int
	out  int
}{
	{"first", 14, 2},
	{"second", 1969, 966},
	{"third", 100756, 50346},
}

func TestFuel(t *testing.T) {
	for _, tt := range firstFueltTestCases {
		t.Run(tt.name, func(t *testing.T) {
			fuel := nineteen.DoIt(tt.in)
			if fuel != tt.out {
				t.Errorf("got %d, want %d for %d", fuel, tt.out, tt.in)
			}
		})
	}
}

func TestFuelWithFuel(t *testing.T) {
	for _, tt := range secondFueltTestCases {
		t.Run(tt.name, func(t *testing.T) {
			fuel := nineteen.DoItFuel(tt.in, 0)
			if fuel != tt.out {
				t.Errorf("got %d, want %d for %d", fuel, tt.out, tt.in)
			}
		})
	}
}

func TestSumItForAll(t *testing.T) {
	total := nineteen.SumItForAll()
	if total != 0 {
		t.Errorf("Sum is: %d.", total)
	}
}

func TestSumItForAllWithFuel(t *testing.T) {
	total := nineteen.SumItForAllWithFuel()
	if total != 0 {
		t.Errorf("Sum it for all with fuel is: %d.", total)
	}
}
