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

	// Remove every number which is not adjacent to *
	// write an absorb funciton

}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isSymbol(s string) bool {
	_, err := strconv.Atoi(s)
	return err != nil && s == "*"
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
