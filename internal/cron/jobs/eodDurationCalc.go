package jobs

import (
	"fmt"
	"task-time-logger-go/internal/logger"
	"task-time-logger-go/internal/storage"
	"task-time-logger-go/utils"
	"time"
)

func EodDurationJob() {
	tickets := storage.GetAllTasks()
	fmt.Printf("Gathered (%d) ticket(s), \nStarting duration calculations...\n", len(tickets))

	for _, ticket := range tickets {
		if time.Time(ticket.EndedOn).IsZero() {
			if ticket.PrevDuration == "" {
				// Calculate duration between StartedOn and now
				duration := utils.CalculateWorkDuration(time.Time(ticket.StartedOn), time.Now())
				fmt.Printf("Calculated duration for ticket %s: %s\n", ticket.ID, duration)
			} else {
				// Parse previous duration and add 9h30m
				prevTime, err := utils.ParseTimeAgo(ticket.PrevDuration)
				if err != nil {
					logger.AppLogger.Printf("Error parsing previous duration for ticket %s: %v", ticket.ID, err)
					continue
				}

				// Add 9h30m to the previous time
				newTime := prevTime.Add(9*time.Hour + 30*time.Minute)

				// Calculate duration from StartedOn to the new time
				duration := utils.CalculateWorkDuration(time.Time(ticket.StartedOn), newTime)
				fmt.Printf("Calculated duration for ticket %s with previous duration: %s\n", ticket.ID, duration)
			}
		}
	}
}
