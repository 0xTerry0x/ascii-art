package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) > 2 || len(args) < 1 {
		fmt.Println("Incorrect amount of arguments")
		os.Exit(1)
	}

	Text := args[0]
	if len(Text) == 0 {
		fmt.Println("Text argument is empty")
		os.Exit(1)
	}

	inputFile := "text/standard.txt"
	if len(args) == 2 {
		switch args[1] {
		case "--standard":
			inputFile = "text/standard.txt"
		case "--shadow":
			inputFile = "text/shadow.txt"
		case "--thinkertoy":
			inputFile = "text/thinkertoy.txt"
		default:
			fmt.Println("Invalid style argument")
			os.Exit(1)
		}
	}

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	charHeight := 9

	inputLines := strings.SplitSeq(Text, "\\n")

	for line := range inputLines {
		if line == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < charHeight-1; row++ {
			for _, ch := range line {
				value := int(ch)
				if value < 32 || value > 126 {
					continue
				}

				blockStart := (value - 32) * charHeight
				fmt.Print(lines[blockStart+row])
			}
			fmt.Println()
		}
	}
}
