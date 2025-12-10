package main

import (
	"fmt"
	"os"

	"asciiArt/internal"
)

func main() {
	Text, inputFile := internal.GetArgs(os.Args[1:])

	cleanText, indexes := internal.Filter(Text)

	lines := internal.ReadFile(inputFile)

	output := internal.GenerateArt(cleanText, lines)
	fmt.Print(output)

	internal.PrintNPErr(indexes)
}
