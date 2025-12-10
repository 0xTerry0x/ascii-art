package internal

import (
	"os"
	"bufio"
	"fmt"
)

func ReadFile(inputFile string) []string {
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

	if scanner.Err() != nil {
		fmt.Println("Error loading font:", scanner.Err())
		os.Exit(1)
	}

	return lines
}