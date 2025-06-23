package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"task-time-logger-go/internal/logger"
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

func CalculateWorkDuration(startTime, endTime time.Time) string {
	if endTime.Before(startTime) {
		return "0"
	}

	totalDuration := endTime.Sub(startTime)
	workingDayStart := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), constants.WorkingDayStart, 0, 0, 0, time.Local)
	logger.AppLogger.Printf("Total duration: %v, Working day start: %v", totalDuration, workingDayStart)

	return fmt.Sprintf("%v", totalDuration)
}

func ParseTimeAgo(timeAgoStr string) (time.Time, error) {
	now := time.Now()
	timeAgoStr = strings.TrimSpace(strings.ToLower(timeAgoStr))

	switch timeAgoStr {
	case "just now":
		return now, nil
	case "yesterday":
		return now.AddDate(0, 0, -1), nil
	case "in the future":
		return now.Add(time.Hour), nil // Return 1 hour in future as default
	}

	// Use regex to parse patterns like "5m ago", "2h ago", "3d ago", etc.
	re := regexp.MustCompile(`^(\d+)(m|h|d|w|mo|y)\s+ago$`)
	matches := re.FindStringSubmatch(timeAgoStr)

	if len(matches) != 3 {
		return time.Time{}, fmt.Errorf("invalid time ago format: %s", timeAgoStr)
	}

	value, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid number in time ago string: %s", matches[1])
	}

	unit := matches[2]
	var duration time.Duration

	switch unit {
	case "m":
		duration = time.Duration(value) * time.Minute
	case "h":
		duration = time.Duration(value) * time.Hour
	case "d":
		duration = time.Duration(value) * 24 * time.Hour
	case "w":
		duration = time.Duration(value) * 7 * 24 * time.Hour
	case "mo":
		duration = time.Duration(value) * 30 * 24 * time.Hour // Approximate
	case "y":
		duration = time.Duration(value) * 365 * 24 * time.Hour // Approximate
	default:
		return time.Time{}, fmt.Errorf("unknown time unit: %s", unit)
	}

	return now.Add(-duration), nil
}
