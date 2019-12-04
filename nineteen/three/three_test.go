package three_test

import (
	"testing"

	"github.com/paunik/adventofcode/nineteen/three"
)

var wireToLine = []struct {
	name   string
	wOneIn []string
	out    []three.Coordinate
}{
	{"first", []string{}, []Line{}, 0},
	{"second", []string{"R8", "U5", "L5", "D3"}, []string{makethree.Line{8, 0, "h"}, three.Line{5, 8, "v"}, three.Line{5, 3, "h"}, three.Line{2, 3, "v"}}},
}

func TestWireToLines(t *testing.T) {
	for _, tt := range wireToLine {
		t.Run(tt.name, func(t *testing.T) {
			// result := three.ConvertWiresToLines(tt.wOneIn)
			// if reflect.DeepEqual(fuel, tt.out) == false {
			// 	t.Errorf("got %d, want %d for %d", fuel, tt.out, tt.in)
			// }
		})
	}
}
