package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	state := make([]int, 0)
	state = append(state, 1)

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Printf("line %s\n", line)
		parts := strings.Split(line, " ")
		instruction := parts[0]
		if instruction == "noop" {
			state = append(state, 0)
		} else if instruction == "addx" {
			state = append(state, 0)
			number, _ := strconv.Atoi(parts[1])
			state = append(state, number)
		}
	}

	fmt.Println(state)

	startCycle := 20
	result := 0

	x := 0
	for i := 0; i < len(state); i++ {
		if i == startCycle {
			result += x * startCycle
			startCycle += 40
		}

		x += state[i]
		// fmt.Printf("%d ", x)
	}

	fmt.Println("result", result)

}
