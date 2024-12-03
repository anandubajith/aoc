package main

import (
	"bufio"
	"fmt"
	"math"
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
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	safeCount := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, " ")
		numbers := make([]int, 0, 1000)

		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, num)
		}

		if solve1(numbers) {
			safeCount++
		} else {
			// remove
			for i := 0; i < len(numbers); i++ {
				new := make([]int, 0, 1000)
				for j := 0; j < len(numbers); j++ {
					if j != i {
						new = append(new, numbers[j])
					}
				}
				if solve1(new) {
					safeCount++
					break
				}
			}

		}

	}
	fmt.Println(safeCount)

}

func solve1(numbers []int) bool {

	deltaSet := make([]int, len(numbers), 1000)
	for i := 1; i < len(numbers); i++ {
		deltaSet[i] = numbers[i] - numbers[i-1]
	}

	isAllNegative := true
	isAllPositive := true
	for _, delta := range deltaSet {
		if delta > 0 {
			isAllNegative = false
		}
		if delta < 0 {
			isAllPositive = false
		}
	}

	allInRange := true

	for i := 1; i < len(numbers); i++ {
		abs := int(math.Abs(float64(deltaSet[i])))
		if abs < 1 || abs > 3 {
			allInRange = false
			break
		}
	}

	// fmt.Println(deltaSet)
	// fmt.Println("isAllNegative", isAllNegative)
	// fmt.Println("isAllPositive", isAllPositive)
	// fmt.Println("allInRange", allInRange)
	result := (isAllNegative || isAllPositive) && allInRange

	return result
}
