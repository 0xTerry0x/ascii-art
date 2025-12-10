package internal

import (
	"os"
	"fmt"
)

func GetArgs(args []string) (string, string) {
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

	return Text, inputFile
}