package eighteen_test

import (
	"fmt"
	"testing"

	"github.com/paunik/adventofcode/eighteen"
)

func TestDoIt(t *testing.T) {
	testInput := []string{"+1", "-1"}

	for i, v := range testInput {
		fmt.Println(i, v)
		eighteen.DoIt(v)
	}

	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	 }


}
