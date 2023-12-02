package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var lookUpTable = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func match(s string) int {
	if val, ok := lookUpTable[s]; ok {
		return val
	}
	return -1
}

func Day1Solution() {

	input, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	result := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		leftMost := -1
		rightMost := -1
		for i := 0; i < len(scanner.Text()); i++ {
			if _, err := strconv.Atoi(string(scanner.Text()[i])); err == nil {
				if leftMost == -1 {
					leftMost, err = strconv.Atoi(string(scanner.Text()[i]))
					rightMost, err = strconv.Atoi(string(scanner.Text()[i]))
				} else {
					rightMost, err = strconv.Atoi(string(scanner.Text()[i]))
				}
			}
		}
		result += (leftMost * 10) + rightMost

	}

	fmt.Printf("Day one solution: %v", result)
    fmt.Println()
}
