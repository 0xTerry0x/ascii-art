package internal

import (
	"fmt"
	"os"
)

func ParseArgs(args []string) (string, string) {
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Incorrect number of arguments")
		os.Exit(1)
	}

	text := args[1]
	if text == "" {
		fmt.Println("Text argument is empty")
		os.Exit(1)
	}

	inputFile := "text/standard.txt"

	if len(args) == 3 {
		switch args[2] {
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

	return text, inputFile
}
