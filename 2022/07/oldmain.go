package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Command   string
	Arguments []string
	Output    []string
}

type Tree struct {
	Name     string
	parent   *Tree
	Children []Tree
}

func main() {
	readFile, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var currentCommand Command

	var root Tree

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// fmt.Println(line)
		if line[0] == '$' {
			if currentCommand.Command != "" {
				processCommand(currentCommand, root)
				currentCommand = Command{}
			}
			parts := strings.Split(line, " ")
			currentCommand.Command = parts[1]
			currentCommand.Arguments = parts[2:]
		} else {
			currentCommand.Output = append(currentCommand.Output, line)
		}
	}
	processCommand(currentCommand, root)

}

func processCommand(command Command, root Tree) {
	fmt.Println("processCommand = ", command)
	if command.Command == "cd" && command.Arguments[0] == "/" {
		// Setup root
		root.Name = command.Arguments[0]
		root.Children = make([]Tree, 0)
	} else if command.Command == "ls" {

	}
}
