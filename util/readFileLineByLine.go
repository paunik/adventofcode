package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// ReadFileLineByLine read file line by line
func ReadFileLineByLine(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
