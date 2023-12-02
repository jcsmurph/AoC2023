package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Marble struct {
	color string
	count int
}

type Game struct {
	marbles map[string][][]Marble
}

var ValidGame = map[string]map[string]int{
	"ValidGame": {"red": 12, "green": 13, "blue": 14},
}

func readInput() Game {
	input, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	allGames := Game{}

	for scanner.Scan() {
		data := string(scanner.Text())
		gameNumber := strings.Split(data, ":")[0]
		games := strings.Split(data, ":")[1]

		for idx, game := range strings.Split(games, ";") {
			currentGameMarbles := make([][]Marble, 100)
			marbles := strings.Split(game, ",")
			for _, marble := range marbles {
				currentMarble := Marble{}
				color := strings.Split(marble, " ")[2]
				count := strings.Split(marble, " ")[1]
				currentMarble.color = color
				currentMarble.count, _ = strconv.Atoi(count)
				currentGameMarbles[idx] = append(currentGameMarbles[idx], currentMarble)
			}

			if allGames.marbles == nil {
				allGames.marbles = make(map[string][][]Marble)
			}

			allGames.marbles[gameNumber] = append(allGames.marbles[gameNumber], currentGameMarbles[idx])
		}

	}

	return allGames
}

func CheckValidGame() int {
	games := readInput()

	validGames := 0

	for gameID, game := range games.marbles {
		currentGameInvalid := false
		for _, marbles := range game {
			currentHandInvalid := false
			for _, marble := range marbles {
				if ValidGame["ValidGame"][marble.color] < marble.count {
					currentHandInvalid = true
				}
			}
			if currentHandInvalid {
				currentGameInvalid = true
				break
			}
		}

		if !currentGameInvalid {
			gameNumber, _ := strconv.Atoi(strings.Split(gameID, " ")[1])
			validGames += gameNumber
		}
	}
	return validGames

}
