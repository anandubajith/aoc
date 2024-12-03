package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input2.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	itemCount := 0

	var state [][]byte

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)

		if itemCount == 0 {
			length := len(line)
			itemCount = (length + 1) / 4
			fmt.Println("itemCount = ", itemCount)
			state = make([][]byte, itemCount)
			for i := range state {
				state[i] = make([]byte, 0)
			}
			processInputLine(line, itemCount, state)
		} else if len(line) == (4*itemCount-1) && strings.Contains(line, "[") {
			processInputLine(line, itemCount, state)
		} else if len(line) == (4*itemCount-1) && !strings.Contains(line, "[") {
			fmt.Println("initial state")
			visualize(state)
			fmt.Println()
		} else if strings.Contains(line, "move") {
			// move [count] from [source] to [target]
			parts := strings.Split(line, " ")

			count, _ := strconv.Atoi(parts[1])
			source, _ := strconv.Atoi(parts[3])
			target, _ := strconv.Atoi(parts[5])
			// processMoveV1(count, source-1, target-1, data)
			processMoveV2(count, source-1, target-1, state)

			visualize(state)
			fmt.Println("___")
		}
	}

	// fmt.Println("data = ", data)

	fmt.Println()
	for i := 0; i < itemCount; i++ {
		fmt.Printf("%c", state[i][0])
	}
	fmt.Println()

}

func processInputLine(line string, itemCount int, data [][]byte) {
	for i := 0; i < itemCount; i++ {
		startIndex := i * 4
		endIndex := startIndex + 3
		character := line[startIndex:endIndex][1]
		if character == ' ' {
			continue
		}
		data[i] = append(data[i], character)
	}
}

func processMoveV1(count int, source int, target int, data [][]byte) {
	fmt.Printf("processMove count = %d, source = %d, target = %d\n", count, source, target)
	for i := 0; i < count; i++ {
		item := data[source][0]
		data[source] = data[source][1:]
		data[target] = slices.Insert(data[target], 0, item)
	}
}

func processMoveV2(count int, source int, target int, data [][]byte) {
	fmt.Printf("processMove count = %d, source = %d, target = %d\n", count, source, target)

	items := data[source][0:count]
	data[source] = data[source][count:]
	data[target] = slices.Insert(data[target], 0, items...)
}

func visualize(data [][]byte) {
	length := len(data)

	itemCount := 0
	for i := 0; i < length; i++ {
		itemCount = int(math.Max(float64(itemCount), float64(len(data[i]))))
	}

	for i := 0; i < length; i++ {
		padding := itemCount - len(data[i])
		for j := 0; j < padding; j++ {
			fmt.Printf("   ")
		}
		for j := 0; j < len(data[i]); j++ {
			fmt.Printf("[%c]", data[i][j])
		}
		fmt.Println()
	}
}
