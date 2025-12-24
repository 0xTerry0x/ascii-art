package ascii

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

	inputFile := "internal/assets/standard.txt"
	if len(args) == 2 {
		switch args[1] {
		case "--standard":
			inputFile = "internal/assets/standard.txt"
		case "--shadow":
			inputFile = "internal/assets/shadow.txt"
		case "--thinkertoy":
			inputFile = "internal/assets/thinkertoy.txt"
		default:
			fmt.Println("Invalid style argument")
			os.Exit(1)
		}
	}

	return Text, inputFile
}
