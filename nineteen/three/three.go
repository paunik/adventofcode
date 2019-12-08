package three

import (
	"fmt"
	"strconv"
	"strings"
)

// ConvertWiresToLines convert input string to collection of Lines
func ConvertWiresToLines(wireOne, wireTwo string) int {

	w1 := strings.Split(wireOne, ",")
	w2 := strings.Split(wireTwo, ",")

	sX := 0
	sY := 0
	var o string

	m1 := make(map[string]string)

	totalWireLength := 0

	for _, w := range w1 {
		input := strings.SplitN(w, "", 2)
		step, _ := strconv.Atoi(input[1])

		k := fmt.Sprintf("x:%d,y:%d", sX, sY)
		m1[k] = "0"

		if input[0] == "L" || input[0] == "R" {
			o = "H"
		} else {
			o = "V"
		}

		for index := 0; index < step; index++ {
			if input[0] == "R" {
				sX = sX + 1
			}
			if input[0] == "L" {
				sX = sX - 1
			}
			if input[0] == "U" {
				sY = sY + 1
			}
			if input[0] == "D" {
				sY = sY - 1
			}

			totalWireLength++

			k := fmt.Sprintf("x:%d,y:%d", sX, sY)

			m1[k] = fmt.Sprintf("%s:%d", o, totalWireLength)
		}
	}

	sX = 0
	sY = 0

	m2 := make(map[string]string)

	totalWireLength = 0

	for _, w := range w2 {
		input := strings.SplitN(w, "", 2)
		step, _ := strconv.Atoi(input[1])

		if input[0] == "L" || input[0] == "R" {
			o = "H"
		} else {
			o = "V"
		}

		k := fmt.Sprintf("x:%d,y:%d", sX, sY)
		m2[k] = "0"

		for index := 0; index < step; index++ {
			if input[0] == "R" {
				sX = sX + 1
			}
			if input[0] == "L" {
				sX = sX - 1
			}
			if input[0] == "U" {
				sY = sY + 1
				o = "V"
			}
			if input[0] == "D" {
				sY = sY - 1
				o = "V"
			}

			totalWireLength++

			k := fmt.Sprintf("x:%d,y:%d", sX, sY)

			m2[k] = fmt.Sprintf("%s:%d", o, totalWireLength)
		}
	}

	x := make([]string, 0)

	for k := range m1 {

		v1, c1 := m1[k]
		v2, c2 := m2[k]

		rV1 := strings.SplitN(v1, ":", 2)
		rV2 := strings.SplitN(v2, ":", 2)

		v1 = rV1[0]
		v2 = rV2[0]

		if c2 != false && c1 != false && ((v1 == "H" && v2 == "V") || (v1 == "V" && v2 == "H")) {

			// fmt.Println(" ")
			// fmt.Print("k: ", k)
			// fmt.Print(", v1: ", v1)
			// fmt.Print(", v2: ", v2)
			// fmt.Print(", tw1: ", rV1[1])
			// fmt.Print(", tw2: ", rV2[1])
			// x = append(x, k) first
			x = append(x, fmt.Sprintf("x:%s,y:%s", rV1[1], rV2[1]))
		}
	}
	fmt.Println(" ")

	min := make([]int, 0)

	for _, v := range x {
		www := strings.Split(v, ",")

		xxx := strings.Split(www[0], ":")
		yyy := strings.Split(www[1], ":")

		xInt, _ := strconv.Atoi(xxx[1])
		yInt, _ := strconv.Atoi(yyy[1])

		min = append(min, Abs(xInt)+Abs(yInt))
	}

	mm := min[0]
	for _, v := range min {
		if v < mm {
			mm = v
		}
	}

	return mm
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
