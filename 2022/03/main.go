package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// solve1()
	solve2()
}
func solve2() {
	c := make(chan int)
	defer close(c)
	readFile, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	lines := make([]string, 0, 3)

	pendingCount := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
		if len(lines) == 3 {
			fmt.Println("got 3 lines", lines)
			go solveLineGroup(lines, c)
			pendingCount += 1
			lines = make([]string, 0, 3)
		}
	}

	sum := 0
	for pendingCount > 0 {
		priority := <-c
		sum += priority
		pendingCount -= 1
	}
	fmt.Println("part2 = sum = ", sum)

}

func solveLineGroup(lineGroup []string, c chan int) {
	lookupSet := make(map[byte]int)

	for _, line := range lineGroup {
		internalSet := make(map[byte]bool)
		for i := 0; i < len(line); i++ {
			internalSet[line[i]] = true
		}
		for key, value := range internalSet {
			if value == true {
				lookupSet[key] += 1
			}
		}
	}

	// iterate lookupSet
	for key, value := range lookupSet {
		if value >= len(lineGroup) {
			fmt.Printf("Found conflict %c priority = %d \n", key, value)
			fmt.Println()

			priority := 0
			if key > 'a' {
				priority = int(key - 'a' + 1)
			} else {
				priority = int(key - 'A' + 26 + 1)
			}
			c <- priority
			break
		}
	}
}

func solve1() {
	readFile, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	totalPriority := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		length := len(line)

		lookupSet := make(map[byte]bool)

		for i := 0; i < length/2; i++ {
			lookupSet[line[i]] = true
		}

		for i := length / 2; i < length; i++ {
			if lookupSet[line[i]] == true {
				conflict := 0
				if line[i] > 'a' {
					// fmt.Printf("Found conflict %c priority = %d \n", line[i], line[i]-'a'+1)
					conflict = int(line[i] - 'a' + 1)
				} else {
					// fmt.Printf("Found conflict %c priority = %d \n", line[i], line[i]-'A'+26+1)
					conflict = int(line[i] - 'A' + 26 + 1)
				}
				if conflict > 0 {
					totalPriority += conflict
					break
				}
			}
		}
	}
	fmt.Printf("part1 = totalPriority = %d\n", totalPriority)
}
