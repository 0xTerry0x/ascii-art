package main

import "strings"

func GenerateArt(text string, lines []string) string {
	charHeight := 9
	var sb strings.Builder
	inputLines := strings.SplitSeq(text, "\\n")

	for line := range inputLines {
		if line == "" {
			sb.WriteByte('\n')
			continue
		}
		for row := 0; row < charHeight-1; row++ {
			for _, ch := range line {
				value := int(ch)
				if value < 32 || value > 126 {
					continue
				}

				blockStart := (value - 32) * charHeight
				if blockStart+row < len(lines) {
					sb.WriteString(lines[blockStart+row])
				}
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
