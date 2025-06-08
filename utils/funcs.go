package utils

import (
	"fmt"
	"task-time-logger-go/internal/models/enums/constants"
	"time"
)

func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	// Handle future times (just in case)
	if diff < 0 {
		return "in the future"
	}

	// Today's date for comparison
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	// Yesterday's date
	yesterday := today.AddDate(0, 0, -1)

	switch {
	// Less than a minute
	case diff < time.Minute:
		return "just now"

	// Less than an hour
	case diff < time.Hour:
		minutes := int(diff.Minutes())
		return fmt.Sprintf("%dm ago", minutes)

	// Less than a day (today)
	case t.After(today):
		hours := int(diff.Hours())
		return fmt.Sprintf("%dh ago", hours)

	// Yesterday
	case t.After(yesterday) && t.Before(today):
		return "yesterday"

	// Less than a week
	case diff < 7*24*time.Hour:
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%dd ago", days)

	// Less than a month (~4 weeks)
	case diff < 4*7*24*time.Hour:
		weeks := int(diff.Hours() / (24 * 7))
		return fmt.Sprintf("%dw ago", weeks)

	// Less than a year
	case diff < 365*24*time.Hour:
		months := int(diff.Hours() / (24 * 30))
		return fmt.Sprintf("%dmo ago", months)

	// Years
	default:
		years := int(diff.Hours() / (24 * 365))
		return fmt.Sprintf("%dy ago", years)
	}
}

func CalculateWorkDuration(startTime, endTime time.Time) time.Duration {
	if endTime.Before(startTime) {
		return 0
	}
	var totalDuration time.Duration
	workingDayStart := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), constants.WorkingDayStart, 0, 0, 0, time.Local)
	fmt.Println(totalDuration, workingDayStart)
	return 0
}
