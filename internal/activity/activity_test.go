package activity

import (
	"testing"
)

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"negative", -24, "0m"},
		{"zero", 0, "0m"},
		{"minutes", 24, "24m"},
		{"hours", 120, "2h"},
		{"mixed", 124, "2h 4m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatDuration(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, got)
			}
		})
	}
}
