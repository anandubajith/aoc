package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	currentPath := make([]string, 0)

	fileMap := make(map[string]int)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line[0] == '$' {
			parts := strings.Split(line, " ")
			if parts[1] == "cd" {
				if parts[2] == "/" {
					currentPath = []string{""} // Use an empty string for the root
				} else if parts[2] == ".." {
					if len(currentPath) > 0 {
						currentPath = currentPath[:len(currentPath)-1]
					}
				} else {
					currentPath = append(currentPath, parts[2])
				}
			}
		} else {
			if strings.Contains(line, "dir") {
				continue
			}
			parts := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(parts[0])
			normalizedPath := strings.Join(currentPath, "/") + "/" + parts[1]
			fileMap[normalizedPath] = fileSize
		}
	}

	total := 0

	directoryMap := make(map[string]int)
	for k, _ := range fileMap {
		parts := strings.Split(k, "/")
		directories := make([]string, 0)
		for i := 0; i < len(parts)-1; i++ {
			directories = append(directories, parts[i])
		}
		normalizedPath := strings.Join(directories, "/")
		directoryMap[normalizedPath] = directoryMap[normalizedPath] + fileMap[k]
	}

	// fmt.Println(directoryMap)
	for _, v := range directoryMap {
		if v < 10000 {
			total = v
		}
	}

	fmt.Println(total)
}
