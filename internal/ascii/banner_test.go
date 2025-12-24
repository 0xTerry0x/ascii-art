package ascii_test

import (
	"testing"

	"asciiArt/internal/ascii"
)

// TestGetArgs tests the argument parsing function
func TestGetArgs(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedText  string
		expectedFile  string
		shouldPanic   bool
	}{
		{
			name:         "Default standard banner",
			args:         []string{"Hello"},
			expectedText: "Hello",
			expectedFile: "internal/assets/standard.txt",
		},
		{
			name:         "Shadow banner",
			args:         []string{"Hello", "--shadow"},
			expectedText: "Hello",
			expectedFile: "internal/assets/shadow.txt",
		},
		{
			name:         "Thinkertoy banner",
			args:         []string{"Hello", "--thinkertoy"},
			expectedText: "Hello",
			expectedFile: "internal/assets/thinkertoy.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text, file := ascii.GetArgs(tt.args)
			if text != tt.expectedText {
				t.Errorf("GetArgs() text = %v, want %v", text, tt.expectedText)
			}
			if file != tt.expectedFile {
				t.Errorf("GetArgs() file = %v, want %v", file, tt.expectedFile)
			}
		})
	}
}

// TestReadFile tests the banner file loading function
func TestReadFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		wantErr  bool
	}{
		{
			name:     "Load standard banner",
			filePath: "../../internal/assets/standard.txt",
			wantErr:  false,
		},
		{
			name:     "Load shadow banner",
			filePath: "../../internal/assets/shadow.txt",
			wantErr:  false,
		},
		{
			name:     "Load thinkertoy banner",
			filePath: "../../internal/assets/thinkertoy.txt",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Skip if panic is expected
			defer func() {
				if r := recover(); r != nil && !tt.wantErr {
					t.Errorf("ReadFile() panicked: %v", r)
				}
			}()

			lines := ascii.ReadFile(tt.filePath)
			if len(lines) == 0 && !tt.wantErr {
				t.Errorf("ReadFile() returned empty lines")
			}
		})
	}
}
