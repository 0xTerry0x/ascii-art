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

			expected := ""
			
			Input string in the test case mimics what `ascii.go` receives.

			parts := strings.Split(tt.text, "\\n")
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
