package eighteen

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	input   = "https://adventofcode.com/2018/day/1/input"
	tmpFile = "/tmp/input.txt"
)

// LoadFile is loading
func LoadFile() {
	file, err := os.Open(tmpFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		DoIt(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//
// DoIt actually does the work for first exercise
func DoIt(textInteger string) int {
	// util.DownloadFile(tmpFile, input)

	mFreq := make(map[int]int)

	res := 0
	mFreq[res] = 1
	mFreq[0] = 1

	cnt := 0

	inputIntiger, err := strconv.Atoi(textInteger)
	if err != nil {
		panic(err)
	}

	res += inputIntiger

	if x, found := mFreq[res]; found {
		fmt.Println("x: ", x)
		fmt.Println("res: ", res)
		found = true
	} else {
		mFreq[res] = 1
	}

	if cnt == 1 {
		fmt.Println("res:", res)
		fmt.Println("int:", inputIntiger)
	}

	fmt.Println(res)
	return res
}
