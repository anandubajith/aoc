package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Draw struct {
	redCount   int
	blueCount  int
	greenCount int
}

type Game struct {
	ID    int
	draws []Draw
}

func parseInput() []Game {
	s := bufio.NewScanner(os.Stdin)
	var gameList []Game
	for s.Scan() {
		line := s.Text()
		gameParts := strings.Split(line, ":")
		gameId := strings.Split(gameParts[0], " ")[1]

		gameItems := strings.Split(strings.Trim(gameParts[1], " "), ";")
		var itemsStructList []Draw
		for _, item := range gameItems {
			itemsWithoutSpaces := strings.Trim(item, " ")
			itemParts := strings.Split(itemsWithoutSpaces, ", ")
			currentItem := Draw{}
			for _, itemPart := range itemParts {
				parts := strings.Split(itemPart, " ")
				valAsNumber, _ := strconv.Atoi(parts[0])
				if parts[1] == "red" {
					currentItem.redCount += valAsNumber
				} else if parts[1] == "blue" {
					currentItem.blueCount += valAsNumber
				} else if parts[1] == "green" {
					currentItem.greenCount += valAsNumber
				}
			}
			itemsStructList = append(itemsStructList, currentItem)
		}
		gameIdNumber, _ := strconv.Atoi(gameId)
		gameList = append(
			gameList,
			Game{
				ID:    gameIdNumber,
				draws: itemsStructList,
			})
	}
	return gameList
}

func main() {
	gameList := parseInput()
	p1result := solveP1(gameList)
	p2result := solveP2(gameList)

	fmt.Printf("P1: %d\n", p1result)
	fmt.Printf("P2: %d\n", p2result)
}

func solveP2(gameList []Game) int {
	result := 0

	for _, game := range gameList {
		maxBlue := 0
		maxRed := 0
		maxGreen := 0
		for _, item := range game.draws {
			maxBlue = max(maxBlue, item.blueCount)
			maxRed = max(maxRed, item.redCount)
			maxGreen = max(maxGreen, item.greenCount)
		}
		fmt.Printf("Game %d: Power %d\n", game.ID, maxBlue*maxRed*maxGreen)
		result += maxBlue * maxRed * maxGreen
	}

	return result
}

func solveP1(gameList []Game) int {
	result := 0

	for _, game := range gameList {
		maxBlue := 0
		maxRed := 0
		maxGreen := 0
		for _, item := range game.draws {
			maxBlue = max(maxBlue, item.blueCount)
			maxRed = max(maxRed, item.redCount)
			maxGreen = max(maxGreen, item.greenCount)
		}
		if maxBlue <= 14 && maxRed <= 12 && maxGreen <= 13 {
			fmt.Printf("Game %d is VALID\n", game.ID)
			result += game.ID
		}
	}

	return result
}
