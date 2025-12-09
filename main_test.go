package main

import (
	"strings"
	"testing"

	"asciiArt/internal"
)

func TestGenerateArt(t *testing.T) {
	charHeight := 9

	// Helper to create mock lines from a map of characters
	createMockLines := func(patterns map[rune]string) []string {
		lines := make([]string, 1000)
		// Initialize empty lines
		for i := range lines {
			lines[i] = ""
		}

		for char, pattern := range patterns {
			start := (int(char) - 32) * charHeight
			for i := 0; i < charHeight; i++ {
				if start+i < len(lines) {
					lines[start+i] = pattern
				}
			}
		}
		return lines
	}

	tests := []struct {
		name     string
		text     string
		patterns map[rune]string
	}{
		{
			name: "hello (lowercase)",
			text: "hello",
			patterns: map[rune]string{
				'h': "h_lower",
				'e': "e_lower",
				'l': "l_lower",
				'o': "o_lower",
			},
		},
		{
			name: "Case Sensitivity (HeLlo)",
			text: "HeLlo",
			patterns: map[rune]string{
				'H': "H_UPPER",
				'h': "h_lower",
				'e': "e_lower",
				'L': "L_UPPER",
				'l': "l_lower",
				'o': "o_lower",
			},
		},
		{
			name: "1Hello 2There",
			text: "1Hello 2There",
			patterns: map[rune]string{
				'1': "1_NUMBER",
				'2': "2_NUMBER",
				'H': "H_UPPER",
				'h': "h_lower",
				'e': "e_lower",
				'l': "l_lower",
				'o': "o_lower",
				'T': "T_UPPER",
				'r': "r_lower",
				' ': " ",
			},
		},
		{
			name: "{Hello There}",
			text: "{Hello There}",
			patterns: map[rune]string{
				'H': "H_UPPER",
				'h': "h_lower",
				'e': "e_lower",
				'l': "l_lower",
				'o': "o_lower",
				'T': "T_UPPER",
				'r': "r_lower",
				'{': "{",
				'}': "}",
				' ': " ",
			},
		},
		{
			name: "Hello\\n\\nThere",
			// We must use literal \\n because the code splits by the STRING "\n", not the character.
			text: "Hello\\n\\nThere",
			patterns: map[rune]string{
				'H': "H_UPPER",
				'h': "h_lower",
				'e': "e_lower",
				'l': "l_lower",
				'o': "o_lower",
				'T': "T_UPPER",
				'r': "r_lower",
			},
		},
		{
			name: "Hello\\nThere",
			text: "Hello\\nThere",
			patterns: map[rune]string{
				'H': "H_UPPER",
				'h': "h_lower",
				'e': "e_lower",
				'l': "l_lower",
				'o': "o_lower",
				'T': "T_UPPER",
				'r': "r_lower",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := createMockLines(tt.patterns)

			// Construct expected output programmatically
			// We must mimic the logic of splitting by \n
			expected := ""

			// Handle literal newline chars if that's what input is
			// The original code uses strings.SplitSeq(text, "\\n")
			// If the input string contains ACTUAL newlines (rune 10), SplitSeq("\\n") won't split them unless they are escaped.
			// Wait, let's look at ascii.go again.
			// inputLines := strings.SplitSeq(text, "\\n")
			// This expects literal backslash+n.

			// If the input is "Hello\n\nThere" (with real newlines), SplitSeq wont split it.
			// BUT, usually in these projects, the input args from command line might come in as "Hello\nThere"
			// where \n is a single character? Or as "Hello\\nThere"?

			// The user ran: go run . "Hello\n\nThere" in PowerShell.
			// PowerShell passes \n as literal newlines usually? Or does it escape them?
			// The output shows it working, which means the program IS splitting them.

			// If the program splits by "\\n", then the input string MUST contain literal backslash followed by n.
			// However, if the user input "Hello\n\nThere" in Go string literal syntax, that is ACTUAL newlines.

			// Let's adjust the mock logic to split by newlines for the expected generation
			// assuming the Input string in the test case mimics what `ascii.go` receives.

			// We will replicate the split logic from ascii.go partially to build expectation.
			parts := strings.Split(tt.text, "\\n")
			// If the test case uses actual newlines, we might need to replace them with \\n first or just split by \n
			if strings.Contains(tt.text, "\n") {
				parts = strings.Split(tt.text, "\n")
			}

			for _, part := range parts {
				if part == "" {
					expected += "\n"
					continue
				}

				for row := 0; row < charHeight-1; row++ {
					rowStr := ""
					for _, char := range part {
						if p, ok := tt.patterns[char]; ok {
							rowStr += p
						} else {
							// If the test case expects a character but didn't define a pattern for it,
							// we should probably fail or mark it clearly so the test doesn't silently pass on empty.
							t.Errorf("Test case %q uses character %q but defines no pattern for it", tt.name, char)
						}
					}
					expected += rowStr + "\n"
				}
			}

			got := internal.GenerateArt(tt.text, lines)
			if got != expected {
				t.Errorf("GenerateArt() = \n%q\nwant \n%q", got, expected)
			}
		})
	}
}
