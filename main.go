package main

import (
	"os"

	"asciiArt/internal"
)

func main() {
	text, inputFile := internal.ParseArgs(os.Args)

	cleanText, indexes := internal.FilterPrintable(text)

	lines := internal.LoadFontFile(inputFile)

	internal.PrintAscii(cleanText, lines, indexes)
}
