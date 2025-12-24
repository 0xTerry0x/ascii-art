package main

import (
	"fmt"
	"os"

	"asciiArt/internal/ascii"
)

func main() {
	Text, inputFile := ascii.GetArgs(os.Args[1:])

	cleanText, indexes := ascii.Filter(Text)

	lines := ascii.ReadFile(inputFile)

	output := ascii.GenerateArt(cleanText, lines)
	fmt.Print(output)

	ascii.PrintNPErr(indexes)
}
