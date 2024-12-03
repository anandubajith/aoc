package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input2.txt")

	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	num1Array := make([]int, 0, 1000)
	num2Array := make([]int, 0, 1000)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		parts := strings.Split(line, "   ")
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		num1Array = append(num1Array, num1)
		num2Array = append(num2Array, num2)
	}
	readFile.Close()

	fmt.Printf("part1 : %d \n", solve1(num1Array, num2Array))
	fmt.Printf("part2 : %d \n", solve2(num1Array, num2Array))

}

func solve1(num1Array []int, num2Array []int) int {

	sort.Sort(sort.IntSlice(num1Array))
	sort.Sort(sort.IntSlice(num2Array))

	sum := 0
	for i := 0; i < len(num1Array); i++ {
		delta := num1Array[i] - num2Array[i]
		sum += int(math.Abs(float64(delta)))
	}
	return sum
}

func solve2(num1Array []int, num2Array []int) int {
	lookupTable := make(map[int]int)
	for i := 0; i < len(num2Array); i++ {
		lookupTable[num2Array[i]] = lookupTable[num2Array[i]] + 1
	}

	// fmt.Println(lookupTable)

	sum := 0
	for i := 0; i < len(num1Array); i++ {
		// check if key exists
		if _, ok := lookupTable[num1Array[i]]; ok {
			sum += lookupTable[num1Array[i]] * num1Array[i]
		}
	}
	// fmt.Println(sum)
	return sum
}
