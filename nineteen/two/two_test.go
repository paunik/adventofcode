package two_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/paunik/adventofcode/nineteen/two"
)

var firstAlarmTestCases = []struct {
	name string
	in   []int
	out  []int
}{
	{"first", []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	{"second", []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
	{"third", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	{"fourth", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
	{"fifth", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

var secondAlarmTestCases = []struct {
	name string
	in   []int
	out  []int
}{
	{"first", []int{}, []int{}},
	{"second", []int{}, []int{}},
	{"third", []int{}, []int{}},
}

func TestAlarm(t *testing.T) {
	for _, tt := range firstAlarmTestCases {
		t.Run(tt.name, func(t *testing.T) {
			fuel := two.AlarmOpcode(tt.in)
			if reflect.DeepEqual(fuel, tt.out) == false {
				t.Errorf("got %d, want %d for %d", fuel, tt.out, tt.in)
			}
		})
	}
}

// func TestFuelWithFuel(t *testing.T) {
// 	for _, tt := range secondAlarmTestCases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			opCodeResult := nineteen.DoItFuel(tt.in, 0)
// 			if fuel != tt.out {
// 				t.Errorf("got %d, want %d for %d", fuel, tt.out, tt.in)
// 			}
// 		})
// 	}
// }

func TestAlarmOne(t *testing.T) {

	var inputTwo = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 19, 5, 23, 1, 6, 23, 27, 1, 27, 5, 31, 2, 31, 10, 35, 2, 35, 6, 39, 1, 39, 5, 43, 2, 43, 9, 47, 1, 47, 6, 51, 1, 13, 51, 55, 2, 9, 55, 59, 1, 59, 13, 63, 1, 6, 63, 67, 2, 67, 10, 71, 1, 9, 71, 75, 2, 75, 6, 79, 1, 79, 5, 83, 1, 83, 5, 87, 2, 9, 87, 91, 2, 9, 91, 95, 1, 95, 10, 99, 1, 9, 99, 103, 2, 103, 6, 107, 2, 9, 107, 111, 1, 111, 5, 115, 2, 6, 115, 119, 1, 5, 119, 123, 1, 123, 2, 127, 1, 127, 9, 0, 99, 2, 0, 14, 0}
	var expected = []int{3267740, 12, 2, 2, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 48, 1, 19, 5, 49, 1, 6, 23, 51, 1, 27, 5, 52, 2, 31, 10, 208, 2, 35, 6, 416, 1, 39, 5, 417, 2, 43, 9, 1251, 1, 47, 6, 1253, 1, 13, 51, 1258, 2, 9, 55, 3774, 1, 59, 13, 3779, 1, 6, 63, 3781, 2, 67, 10, 15124, 1, 9, 71, 15127, 2, 75, 6, 30254, 1, 79, 5, 30255, 1, 83, 5, 30256, 2, 9, 87, 90768, 2, 9, 91, 272304, 1, 95, 10, 272308, 1, 9, 99, 272311, 2, 103, 6, 544622, 2, 9, 107, 1633866, 1, 111, 5, 1633867, 2, 6, 115, 3267734, 1, 5, 119, 3267735, 1, 123, 2, 3267737, 1, 127, 9, 0, 99, 2, 0, 14, 0}

	// modify gravity setting 1202 alarm
	// inputTwo[1] = 12
	// inputTwo[2] = 2

	opCodeResult := two.AlarmOpcode(inputTwo)
	if reflect.DeepEqual(opCodeResult, expected) == false {
		t.Errorf("got %d, for %d", opCodeResult, expected)
	}
}

func TestFindNounAndVerb(t *testing.T) {
	var inputTwo = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 19, 5, 23, 1, 6, 23, 27, 1, 27, 5, 31, 2, 31, 10, 35, 2, 35, 6, 39, 1, 39, 5, 43, 2, 43, 9, 47, 1, 47, 6, 51, 1, 13, 51, 55, 2, 9, 55, 59, 1, 59, 13, 63, 1, 6, 63, 67, 2, 67, 10, 71, 1, 9, 71, 75, 2, 75, 6, 79, 1, 79, 5, 83, 1, 83, 5, 87, 2, 9, 87, 91, 2, 9, 91, 95, 1, 95, 10, 99, 1, 9, 99, 103, 2, 103, 6, 107, 2, 9, 107, 111, 1, 111, 5, 115, 2, 6, 115, 119, 1, 5, 119, 123, 1, 123, 2, 127, 1, 127, 9, 0, 99, 2, 0, 14, 0}
	var haltOpcode = 99
	exit := false

	for nIndex := 0; nIndex < haltOpcode; nIndex++ {
		for vIndex := 0; vIndex < haltOpcode; vIndex++ {

			inputCopy := make([]int, 137)

			copy(inputCopy, inputTwo)

			fmt.Println(inputCopy[0:3])
			fmt.Println(inputTwo[0:3])

			inputCopy[0] = 1
			inputCopy[1] = nIndex
			inputCopy[2] = vIndex

			fmt.Println(inputCopy[0:3])

			result := two.AlarmOpcode(inputCopy)

			if result[0] == 19690720 {
				exit = true
				if nIndex != -1 && vIndex != -1 {
					t.Errorf("got noun=%d, verb=%d", nIndex, vIndex)
				}
			}

			inputCopy = nil
			result = nil
		}

		if exit {
			break
		}
	}
}
