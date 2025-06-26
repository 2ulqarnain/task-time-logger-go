package jobs

import (
	"fmt"
	"task-time-logger-go/internal/logger"
	"task-time-logger-go/internal/storage"
	"time"
)

func EodDurationJob() {
	tickets := storage.GetAllTasks()
	fmt.Printf("Gathered (%d) ticket(s), \nStarting duration calculations...\n", len(tickets))

	for _, ticket := range tickets {
		if time.Time(ticket.EndedOn).IsZero() {
			if ticket.PrevDuration == 0 {
				// Calculate duration between StartedOn and now
				duration := time.Since(time.Time(ticket.StartedOn))
				minutes := int32(duration.Minutes())
				fmt.Printf("Calculated duration for ticket %s: %d minutes\n", ticket.ID, minutes)

				// Save the calculated duration
				if err := storage.UpdateTaskDuration(ticket.ID, minutes); err != nil {
					logger.AppLogger.Printf("Error updating duration for ticket %s: %v", ticket.ID, err)
				}
			} else {
				// Add 9h30m (570 minutes) to previous duration
				totalMinutes := ticket.PrevDuration + 570
				fmt.Printf("Calculated duration for ticket %s with previous duration: %d minutes (total: %d minutes)\n", ticket.ID, ticket.PrevDuration, totalMinutes)

				// Save the updated duration
				if err := storage.UpdateTaskDuration(ticket.ID, totalMinutes); err != nil {
					logger.AppLogger.Printf("Error updating duration for ticket %s: %v", ticket.ID, err)
				}
			}
		}
	}
}
