package internal

func Filter(Text string) (string, []int) {
	cleanText := ""
	var indexes []int

	runeIndex := 1
	for _, ch := range Text {
		if ch < 32 || ch > 126 {
			indexes = append(indexes, runeIndex)
		} else {
			cleanText += string(ch)
		}
		runeIndex++
	}

	return cleanText, indexes
}
