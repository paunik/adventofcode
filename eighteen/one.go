package eighteen

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	tmpFile = "/tmp/input.txt"
)

// LoadFile is loading
func LoadFile() []int {

	res := make([]int, 0)

	file, err := os.Open(tmpFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		no := scanner.Text()

		integerNo, err := strconv.Atoi(no)

		if err != nil {
			log.Fatal(err)
		}
		res = append(res, integerNo)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}

// ChronalCalibration does calculation
func ChronalCalibration(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}

	return sum
}

// ChronalCalibrationDouble does calculation
func ChronalCalibrationDouble(input []int) int {
	mFreq := make(map[int]int)

	res := 0
	mFreq[0] = 1

	shouldExit := false
	cnt := 0
	for {
		for _, i := range input {
			res += i
			if _, found := mFreq[res]; found {
				shouldExit = true
				break
			}

			mFreq[res] = 1
		}

		if shouldExit == true {
			break
		}

		cnt++
	}

	return res
}

// DoIt actually does the work for first exercise
func DoIt() int {
	res := 0
	input := LoadFile()
	for _, inp := range input {
		res += inp
	}
	return res
}

// DoItSecond actually does the work for first exercise
func DoItSecond() int {
	res := 0
	input := LoadFile()
	res = ChronalCalibrationDouble(input)

	return res
}
