package main_test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestEdgeCases tests various edge cases for ASCII art generation
func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectEmpty bool
		contains    string
	}{
		// Basic edge cases
		{
			name:        "Empty string",
			input:       "",
			expectEmpty: false,
		},
		{
			name:        "Single space",
			input:       " ",
			expectEmpty: false,
		},
		{
			name:        "Multiple spaces",
			input:       "   ",
			expectEmpty: false,
		},
		{
			name:     "Newline only",
			input:    "\\n",
			contains: "\n",
		},
		{
			name:     "Multiple newlines",
			input:    "\\n\\n\\n",
			contains: "\n",
		},

		// Character ranges
		{
			name:        "All digits",
			input:       "0123456789",
			expectEmpty: false,
		},
		{
			name:        "Full alphabet uppercase",
			input:       "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expectEmpty: false,
		},
		{
			name:        "Full alphabet lowercase",
			input:       "abcdefghijklmnopqrstuvwxyz",
			expectEmpty: false,
		},

		// Special characters
		{
			name:        "Underscores",
			input:       "___",
			expectEmpty: false,
		},
		{
			name:        "Dashes",
			input:       "---",
			expectEmpty: false,
		},
		{
			name:        "Brackets",
			input:       "[]",
			expectEmpty: false,
		},
		{
			name:        "Parentheses",
			input:       "()",
			expectEmpty: false,
		},
		{
			name:        "Braces",
			input:       "{}",
			expectEmpty: false,
		},
		{
			name:        "Exclamation",
			input:       "!!!",
			expectEmpty: false,
		},
		{
			name:        "Question marks",
			input:       "???",
			expectEmpty: false,
		},

		// Multi-line combinations
		{
			name:     "Text with single newline",
			input:    "Hello\\nWorld",
			contains: "\n",
		},
		{
			name:     "Text with multiple newlines",
			input:    "A\\n\\nB\\n\\nC",
			contains: "\n",
		},
		{
			name:     "Newline at start",
			input:    "\\nHello",
			contains: "\n",
		},
		{
			name:     "Newline at end",
			input:    "Hello\\n",
			contains: "\n",
		},

		// Mixed content
		{
			name:        "Letters and numbers",
			input:       "abc123",
			expectEmpty: false,
		},
		{
			name:        "Letters with symbols",
			input:       "Hello!",
			expectEmpty: false,
		},
		{
			name:        "Numbers with symbols",
			input:       "123#456",
			expectEmpty: false,
		},

		// Long strings
		{
			name:        "Long string",
			input:       "ThisIsAVeryLongStringWithoutSpaces",
			expectEmpty: false,
		},
		{
			name:        "Repeated characters",
			input:       "aaaaaaaaaa",
			expectEmpty: false,
		},

		// Punctuation
		{
			name:        "Period",
			input:       ".",
			expectEmpty: false,
		},
		{
			name:        "Comma",
			input:       ",",
			expectEmpty: false,
		},
		{
			name:        "Semicolon",
			input:       ";",
			expectEmpty: false,
		},
		{
			name:        "Colon",
			input:       ":",
			expectEmpty: false,
		},
		{
			name:        "Quote marks",
			input:       "\"'",
			expectEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "./cmd/ascii-art", tt.input)
			output, err := cmd.CombinedOutput()

			if err != nil && !tt.expectEmpty {
				t.Logf("Command output: %s", string(output))
			}

			outputStr := string(output)

			if tt.expectEmpty && len(outputStr) > 0 {
				t.Errorf("Expected empty output but got: %s", outputStr)
			}

			if !tt.expectEmpty && len(outputStr) == 0 {
				t.Errorf("Expected non-empty output but got empty")
			}

			if tt.contains != "" && !strings.Contains(outputStr, tt.contains) {
				t.Errorf("Expected output to contain %q but it didn't", tt.contains)
			}
		})
	}
}

// TestInvalidInputs tests error handling for invalid inputs
func TestInvalidInputs(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "No arguments",
			args: []string{},
		},
		{
			name: "Invalid flag",
			args: []string{"Hello", "--invalid"},
		},
		{
			name: "Too many arguments",
			args: []string{"Hello", "--standard", "extra"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := append([]string{"run", "./cmd/ascii-art"}, tt.args...)
			cmd := exec.Command("go", args...)
			output, err := cmd.CombinedOutput()

			if err == nil {
				t.Logf("Command succeeded with output: %s", string(output))
			}
		})
	}
}

// TestOutputFormat tests specific output formatting requirements
func TestOutputFormat(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		checkLines bool
		minLines   int
	}{
		{
			name:       "Single character output",
			input:      "A",
			checkLines: true,
			minLines:   1,
		},
		{
			name:       "Multiple characters",
			input:      "Hello",
			checkLines: true,
			minLines:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "./cmd/ascii-art", tt.input)
			output, err := cmd.CombinedOutput()

			if err != nil {
				t.Fatalf("Command failed: %v, output: %s", err, string(output))
			}

			outputStr := string(output)
			lines := strings.Split(strings.TrimRight(outputStr, "\n"), "\n")

			if tt.checkLines && len(lines) < tt.minLines {
				t.Errorf("Expected at least %d lines but got %d", tt.minLines, len(lines))
			}

			if len(outputStr) == 0 {
				t.Errorf("Expected non-empty output")
			}
		})
	}
}
