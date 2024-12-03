package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput()
	for _, row := range input {
		fmt.Println(row)
	}

	rowPos := 0
	for rowPos < len(input) {
		colPos := 0
		for colPos < len(input[rowPos]) {
			// fmt.Println(input[rowPos][colPos])
			if isNumber(input[rowPos][colPos]) {
				isTouchingSymbol := false
				if rowPos > 0 && colPos > 0 && isSymbol(input[rowPos-1][colPos-1]) {
					isTouchingSymbol = true
				}
				if rowPos > 0 && isSymbol(input[rowPos-1][colPos]) {
					isTouchingSymbol = true
				}
				if rowPos > 0 && colPos < len(input[rowPos])-1 && isSymbol(input[rowPos-1][colPos+1]) {
					isTouchingSymbol = true
				}

				if colPos > 0 && isSymbol(input[rowPos][colPos-1]) {
					isTouchingSymbol = true
				}
				if colPos < len(input[rowPos])-1 && isSymbol(input[rowPos][colPos+1]) {
					isTouchingSymbol = true
				}
				if colPos > 0 && rowPos < len(input)-1 && isSymbol(input[rowPos+1][colPos-1]) {
					isTouchingSymbol = true
				}
				if rowPos < len(input)-1 && isSymbol(input[rowPos+1][colPos]) {
					isTouchingSymbol = true
				}
				if rowPos < len(input)-1 && colPos+1 < len(input[rowPos+1])-1 && isSymbol(input[rowPos+1][colPos+1]) {
					isTouchingSymbol = true
				}

				// if whatever to my left isNumber => i inherit their configuration
				if (colPos > 0 && strings.HasPrefix(input[rowPos][colPos-1], "AN")) || isTouchingSymbol {
					input[rowPos][colPos] = "AN" + input[rowPos][colPos]
				}
			}
			colPos++
		}
		rowPos++
	}

	rowPos = 0

	for rowPos < len(input) {
		colPos := len(input[rowPos]) - 1
		for colPos >= 0 {
			if isNumber(input[rowPos][colPos]) && !strings.HasPrefix(input[rowPos][colPos], "AN") {
				if colPos < len(input[rowPos])-1 && strings.HasPrefix(input[rowPos][colPos+1], "AN") {
					input[rowPos][colPos] = "AN" + input[rowPos][colPos]
				}
			}

			colPos--
		}
		rowPos++
	}

	fmt.Println()
	for _, row := range input {
		fmt.Println(row)
	}

	// cleanup
	rowPos = 0
	for rowPos < len(input) {
		colPos := 0
		for colPos < len(input[rowPos]) {
			if strings.HasPrefix(input[rowPos][colPos], "AN") {
				input[rowPos][colPos] = strings.TrimPrefix(input[rowPos][colPos], "AN")
			} else {
				input[rowPos][colPos] = "."
			}
			colPos++
		}
		rowPos++
	}

	fmt.Println()
	for _, row := range input {
		fmt.Println(row)
	}

	// extract the sum
	sum := 0
	rowPos = 0
	for rowPos < len(input) {
		colPos := 0
		currentNumber := 0
		for colPos < len(input[rowPos]) {
			if isNumber(input[rowPos][colPos]) {
				val, _ := strconv.Atoi(input[rowPos][colPos])
				currentNumber = currentNumber*10 + val
			} else {
				sum += currentNumber
				currentNumber = 0
			}
			colPos++
		}
		if currentNumber > 0 {
			sum += currentNumber
		}
		rowPos++
	}

	fmt.Println("sum", sum)

}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isSymbol(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil && s != "."
}

func parseInput() [][]string {
	s := bufio.NewScanner(os.Stdin)
	input := [][]string{}
	for s.Scan() {
		line := s.Text()
		input = append(input, strings.Split(line, ""))
	}

	return input
}
