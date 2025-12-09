package internal

// FilterPrintable removes non-ASCII printable characters
// and returns the cleaned string + rune indexes of removed chars.
func FilterPrintable(s string) (string, []int) {
	cleaned := ""
	var indexes []int

	runeIndex := 0
	for _, ch := range s {
		if ch < 32 || ch > 126 {
			indexes = append(indexes, runeIndex)
		} else {
			cleaned += string(ch)
		}
		runeIndex++
	}
	return cleaned, indexes
}
