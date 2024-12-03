package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	option1("input2.txt")
	option2("input2.txt")

}

func option1(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	current := 0
	sums := make([]int, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			sums = append(sums, current)
			current = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		current += num
	}

	sums = append(sums, current)

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	fmt.Println(sums[0] + sums[1] + sums[2])
}

// insert inserts a number into the top 3 array
func insert(array []int, num int) {
	for i := 0; i < 3; i++ {
		if num > array[i] {
			for j := 2; j > i; j-- {
				array[j] = array[j-1]
			}
			array[i] = num
			break
		}
	}
}

func option2(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	current := 0
	top3 := make([]int, 3)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			insert(top3, current)
			current = 0
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		current += num
	}

	insert(top3, current)

	sum := 0
	for i := 0; i < 3; i++ {
		sum += top3[i]
	}
	fmt.Println(sum)

}
