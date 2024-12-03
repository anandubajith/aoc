package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lookup = map[byte][3]int{
	'X': {4, 1, 7},
	'Y': {8, 5, 2},
	'Z': {3, 9, 6},
}

var lookup2 = map[byte][3]int{
	// I need to lose
	'X': {
		3, // They chose Rock  => I Choose Scissors => 2
		1, // They chose Paper => Choose rock => 1
		2, // They chose Scissors => Choose Paper => 3
	},
	// I need to draw
	'Y': {
		4, // They chose Rock  => Choose Rock => 1 + 3
		5, // They chose Paper => Choose Paper => 2 + 3
		6, // They chose Scissors => Choose Scissors => 3 + 3
	},
	// I need to win
	'Z': {
		8, // They chose Rock  => Choose Paper => 2 + 6
		9, // They chose Paper => Choose Scissors => 3 + 6
		7, // They chose Scissors => Choose Rock => 1 + 6
	},
}

func main() {
	solve2()
}

func solve1() {
	readFile, err := os.Open("input2.txt")
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	sum := 0
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		theirSelection := parts[0][0]
		mySelection := parts[1][0]
		myScore := lookup[mySelection][theirSelection-'A']
		sum += myScore
	}

	fmt.Println("part1 = total = ", sum)
}

func solve2() {
	readFile, err := os.Open("input2.txt")
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	sum := 0
	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		theirSelection := parts[0][0]
		outcome := parts[1][0]
		// fmt.Printf("theirSelection = %c, outcome = %c\n", theirSelection, outcome)
		score := lookup2[outcome][theirSelection-'A']
		sum += score
	}

	fmt.Println("part1 = total = ", sum)
}
