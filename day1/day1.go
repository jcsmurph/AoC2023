package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1Solution() {

	input, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	result := 0
	scanner := bufio.NewScanner(input)

	var numbers []int
	for scanner.Scan() {
		for _, char := range scanner.Text() {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			numbers = append(numbers, number)
			break
		}

		for i := len(scanner.Text()) - 1; i >= 0; i-- {
			number, err := strconv.Atoi(string(scanner.Text()[i]))
			if err != nil {
				continue
			}
			numbers = append(numbers, number)
			break
		}

        result += sliceToInt(numbers)
        numbers = nil

	}

    fmt.Println(result)
}

func sliceToInt(slice []int) int {
	result := 0
	for i := 0; i < len(slice); i++ {
		result = result*10 + slice[i]
	}
	return result
}
