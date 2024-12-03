package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	result := 0
	for s.Scan() {
		line := s.Text()
		// iterate over parts
		firstNumber := 0
		lastNumber := 0

		tokens := tokenizeFromStart(line)
		fromEnd := tokenizeFromEnd(line)

		firstNumber = parse(tokens[0])
		lastNumber = parse(fromEnd[0])

		value := firstNumber*10 + lastNumber
		// println(line, "["+strings.Join(tokens, " ")+"]", value)
		// println(line, "["+strings.Join(fromEnd, " ")+"]", value)

		result = result + value

	}
	fmt.Println(result)
}

func tokenizeFromStart(str string) []string {
	chars := strings.Split(str, "")
	var tokens []string

	ptr1 := 0
	ptr2 := ptr1 + 1
	for ptr1 < len(chars) && ptr2 <= len(chars) {
		candidate := strings.Join(chars[ptr1:ptr2], "")
		if isToken(candidate) {
			tokens = append(tokens, candidate)
			ptr1 = ptr2
			ptr2 = ptr1 + 1
		} else {
			ptr2++
		}
		if ptr2 > len(chars) {
			ptr1++
			ptr2 = ptr1 + 1
		}
	}

	return tokens
}

func reverse(s []string) []string {
	var result []string
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func tokenizeFromEnd(str string) []string {
	chars := reverse(strings.Split(str, ""))
	var tokens []string

	ptr1 := 0
	ptr2 := ptr1 + 1
	for ptr1 < len(chars) && ptr2 <= len(chars) {
		candidate := strings.Join(reverse(chars[ptr1:ptr2]), "")
		if isToken(candidate) {
			tokens = append(tokens, candidate)
			ptr1 = ptr2
			ptr2 = ptr1 + 1
		} else {
			ptr2++
		}
		if ptr2 > len(chars) {
			ptr1++
			ptr2 = ptr1 + 1
		}
	}

	return tokens

}

var TOKENS = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func parse(str string) int {
	// if it's a number => return true
	val, err := strconv.Atoi(str)
	if err == nil {
		return val
	}

	for i, token := range TOKENS {
		if token == str {
			return i + 1
		}
	}
	return 0
}

func isToken(str string) bool {
	// if it's a number => return true
	_, err := strconv.Atoi(str)
	if err == nil {
		return true
	}

	for _, token := range TOKENS {
		if token == str {
			return true
		}
	}
	return false
}
