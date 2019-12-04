package three

// Coordinate is that
type Coordinate struct {
	x, y int
}

// Line represents vector
type Line struct {
	start, end  Coordinate
	orientation string
}

// ConvertWiresToLines convert input string to collection of Lines
func ConvertWiresToLines(wires []string) []Line {
	return make([]Line, 0)
}
