package main

import (
	"bufio"
	"fmt"
	"os"

	"asciiArt/internal"
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

	inputFile := "../text/standard.txt"
	if len(args) == 2 {
		switch args[1] {
		case "--standard":
			inputFile = "../text/standard.txt"
		case "--shadow":
			inputFile = "../text/shadow.txt"
		case "--thinkertoy":
			inputFile = "../text/thinkertoy.txt"
		default:
			fmt.Println("Invalid style argument")
			os.Exit(1)
		}
	}

	newText := ""
	var indexes []int

	runeIndex := 0
	for _, ch := range Text {
		if ch < 32 || ch > 126 {
			indexes = append(indexes, runeIndex)
		} else {
			newText += string(ch)
		}
		runeIndex++
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

	output := internal.GenerateArt(newText, lines)
	fmt.Print(output)

	if len(indexes) > 0 {

		fmt.Print("Non-Printable characters: \n")
		for i, v := range indexes {
			fmt.Print(v)

			if i != len(indexes)-1 {
				fmt.Print(", ")
			} else {
				fmt.Println()
			}

		}
	}
}
