package ascii

import "strings"

func Filter(Text string) (string, []int) {
	var cleanText strings.Builder
	var indexes []int

	runeIndex := 1
	for _, ch := range Text {
		// If the character is not printable, add the index to the indexes slice
		if ch < 32 || ch > 126 {
			indexes = append(indexes, runeIndex)
		} else {
			// If the character is printable, add it to the clean text
			cleanText.WriteString(string(ch))
		}
		runeIndex++
	}

	return cleanText.String(), indexes
}
