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
	fileScanner := bufio.NewScanner(readFile)

	fullOverlap := 0
	partialOverlap := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		parts := strings.Split(line, ",")

		pair1 := strings.Split(parts[0], "-")
		a1, _ := strconv.Atoi(pair1[0])
		a2, _ := strconv.Atoi(pair1[1])
		fmt.Printf("pair1 = (%d, %d)\n", a1, a2)

		pair2 := strings.Split(parts[1], "-")
		b1, _ := strconv.Atoi(pair2[0])
		b2, _ := strconv.Atoi(pair2[1])
		fmt.Printf("pair2 = (%d, %d)\n", b1, b2)

		if a1 <= b1 && a2 >= b2 || b1 <= a1 && b2 >= a2 {
			// fmt.Printf("overlap")
			fullOverlap += 1
		}

		if (a1 <= b1 && a2 >= b2) ||
			(b1 <= a1 && b2 >= a2) ||
			(a1 <= b2 && a2 >= b1) ||
			(b1 <= a2 && b2 >= a1) {
			partialOverlap += 1
		}

		fmt.Println()
	}
	fmt.Println("part1 = fullOverlap = ", fullOverlap)
	fmt.Println("part2 = partialOverlap = ", partialOverlap)

}
