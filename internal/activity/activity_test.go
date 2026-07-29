package activity

import (
	"testing"
)

func TestFormatDurationMinutes(t *testing.T) {
	formattedDuration := FormatDuration(24)

	if formattedDuration != "24m" {
		t.Errorf("Expected %q, got %q", "24m", formattedDuration)
	}
}

func TestFormatDurationHours(t *testing.T) {
	formattedDuration := FormatDuration(120)

	if formattedDuration != "2h" {
		t.Errorf("Expected %q, got %q", "2h", formattedDuration)
	}
}

func TestFormatDurationZero(t *testing.T) {
	formattedDuration := FormatDuration(0)

	if formattedDuration != "0m" {
		t.Errorf("Expected %q, got %q", "0m", formattedDuration)
	}
}

func TestFormatDurationHoursAndMinutes(t *testing.T) {
	formattedDuration := FormatDuration(124)

	if formattedDuration != "2h 4m" {
		t.Errorf("Expected %q, got %q", "2h 4m", formattedDuration)
	}
}
