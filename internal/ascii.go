package internal

import (
	"fmt"
	"strings"
)

func PrintAscii(text string, lines []string, indexes []int) {
	charHeight := 9

	inputLines := strings.SplitSeq(text, "\\n")

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

	if len(indexes) > 0 {
		fmt.Println("Non-Printable characters:")
		for i, v := range indexes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}
