package activity

import "fmt"

// Item represents a logged activity item
type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}

// FormatDuration formats a duration in minutes into a formatted, human-readable string
func FormatDuration(durationInMin int) string {
	switch {
	case durationInMin <= 0:
		return "0m"
	case durationInMin < 60:
		return fmt.Sprintf("%dm", durationInMin)
	case durationInMin%60 == 0:
		return fmt.Sprintf("%dh", durationInMin/60)
	default:
		hours := durationInMin / 60
		minutes := durationInMin % 60
		return fmt.Sprintf("%dh %dm", hours, minutes)

	}
}
