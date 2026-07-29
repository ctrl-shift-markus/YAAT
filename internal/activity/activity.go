package activity

import "fmt"

type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}

func FormatDuration(duration int) string {
	switch {
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
