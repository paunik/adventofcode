package nineteen_test

import (
	"testing"

	"github.com/paunik/adventofcode/nineteen"
)

func TestDoIt(t *testing.T) {
	result := nineteen.DoIt()

	if result != "test" {
		t.Errorf("Expected %s, got: %s.", "test", result)
	}

}
