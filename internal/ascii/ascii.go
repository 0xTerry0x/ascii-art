package ascii

import "strings"

func GenerateArt(text string, lines []string) string {
	// Get the height of the characters in the font
	charHeight := 9
	var sb strings.Builder
	// Split the text into lines
	inputLines := strings.SplitSeq(text, "\\n")

	for line := range inputLines {
		// If the line is empty, add a newline
		if line == "" {
			sb.WriteByte('\n')
			continue
		}
		// For each row of the character, add the corresponding line from the font
		for row := 0; row < charHeight-1; row++ {
			for _, ch := range line {
				value := int(ch)
				// If the character is not printable, skip it
				if value < 32 || value > 126 {
					continue
				}

				// Calculate the starting index of the character in the font
				blockStart := (value - 32) * charHeight
				if blockStart+row < len(lines) {
					// Add the corresponding line from the font
					sb.WriteString(lines[blockStart+row])
				}
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
