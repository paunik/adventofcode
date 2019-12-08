package three_test

import (
	"reflect"
	"testing"

	"github.com/paunik/adventofcode/nineteen/three"
)

var wireToLine = []struct {
	name   string
	wOneIn string
	wTwoIn string
	out    int
}{
	{"first", "R8,U5,L5,D3", "U7,R6,D4,L4", 6},
	{"second", "R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
}

func TestWireToLines(t *testing.T) {
	for _, tt := range wireToLine {
		t.Run(tt.name, func(t *testing.T) {
			result := three.ConvertWiresToLines(tt.wOneIn, tt.wTwoIn)
			if reflect.DeepEqual(result, tt.out) == false {
				t.Errorf("Error for input %s, %s. Expected output: %d got %d", tt.wOneIn, tt.wTwoIn, tt.out, result)
			}
		})
	}
}
