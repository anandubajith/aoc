package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	numbers        []int
}

func main() {
	input := parseInput()
	solveP1(input)
	solveP2(input)

}

func solveP2(input []Card) {

	cardToCountMap := make(map[int]int)
	for _, card := range input {
		cardToCountMap[card.id] = 1
	}

	for _, card := range input {
		winCount := 0
		for _, n := range card.numbers {
			for _, wn := range card.winningNumbers {
				if n == wn {
					winCount++
				}
			}
		}
		for winCount > 0 {
			cardToCountMap[card.id+winCount] += cardToCountMap[card.id]
			winCount--
		}
	}

	totalCards := 0
	for _, v := range cardToCountMap {
		totalCards += v
	}
	fmt.Println("Total cards:", totalCards)
}

func solveP1(input []Card) {
	sum := 0
	for _, card := range input {
		result := 0
		for _, n := range card.numbers {
			for _, wn := range card.winningNumbers {
				if n == wn {
					if result == 0 {
						result = 1
					} else {
						result *= 2
					}
				}
			}
		}

		fmt.Printf("Card %d: %d\n", card.id, result)
		sum += result
	}

	fmt.Printf("Total: %d\n", sum)
}

func parseInput() []Card {
	s := bufio.NewScanner(os.Stdin)
	cards := make([]Card, 0)
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, ": ")
		cardNumber := strings.Trim(strings.TrimPrefix(parts[0], "Card "), " ")
		cardId, _ := strconv.Atoi(cardNumber)

		numberParts := strings.Split(parts[1], " | ")
		winningNumbers := strings.Split(strings.Trim(strings.ReplaceAll(numberParts[0], "  ", " "), " "), " ")
		numbers := strings.Split(strings.Trim(strings.ReplaceAll(numberParts[1], "  ", " "), " "), " ")

		card := Card{}
		card.id = cardId
		card.numbers = make([]int, len(numbers))
		for i, n := range numbers {
			card.numbers[i], _ = strconv.Atoi(strings.Trim(n, " "))
		}

		card.winningNumbers = make([]int, len(winningNumbers))
		for i, n := range winningNumbers {
			card.winningNumbers[i], _ = strconv.Atoi(strings.Trim(n, " "))
		}
		cards = append(cards, card)
	}

	return cards
}
