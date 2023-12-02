package main

import (
	"fmt"

	"github.com/jcsmurph/AoC2023/day1"
	"github.com/jcsmurph/AoC2023/day2"
)

func main() {
    day1.Day1Solution()
    invalidGames := day2.CheckValidGame()
    fmt.Println("Day 2 solution: ", invalidGames)
}
