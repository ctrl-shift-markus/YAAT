package activity

import "fmt"

// Item represents a logged activity item
type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}

// FormatDuration formats a duration in minutes into a formatted, human-readable string
func FormatDuration(duration int) string {
	switch {
	case duration <= 0:
		return "0m"
	case duration < 60:
		return fmt.Sprintf("%dm", duration)
	case duration%60 == 0:
		return fmt.Sprintf("%dh", duration/60)
	default:
		hours := duration / 60
		minutes := duration % 60
		return fmt.Sprintf("%dh %dm", hours, minutes)

	}
}
