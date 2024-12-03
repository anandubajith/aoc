package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)|(don't\\(\\))|(do\\(\\))")

	sum := 0
	enable := true
	for fileScanner.Scan() {
		line := fileScanner.Text()
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch {
			case strings.Contains(match[0], "mul"):
				if enable {
					n1, _ := strconv.Atoi(match[1])
					n2, _ := strconv.Atoi(match[2])
					fmt.Println(n1, n2)
					product := n1 * n2
					sum += product
				}
			case match[0] == "do()":
				enable = true
			case match[0] == "don't()":
				enable = false
			}

			fmt.Println(match)
			fmt.Println(len(match))
		}
	}
	fmt.Println(sum)
}

func solve1() {
	readFile, err := os.Open("input2.txt")

	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			// fmt.Println("key=%s value=%s", match[1], match[2])
			n1, _ := strconv.Atoi(match[1])
			n2, _ := strconv.Atoi(match[2])
			fmt.Println(n1, n2)
			product := n1 * n2
			sum += product
		}
	}
	fmt.Println(sum)
}
