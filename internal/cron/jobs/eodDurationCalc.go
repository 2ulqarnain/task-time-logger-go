package jobs

import (
	"fmt"
	"task-time-logger-go/internal/storage"
	"task-time-logger-go/utils"
	"time"
)

func EodDurationJob() {
	tickets := storage.GetAllTasks()
	fmt.Printf("Gathered (%d) ticket(s), \nStarting duration calculations...", len(tickets))

	for _, ticket := range tickets {
		if time.Time(ticket.EndedOn).IsZero() {
			if ticket.PrevDuration == "" {
				ticket.PrevDuration = utils.TimeAgo(ticket.StartedOn)
				if err := storage.UpdateTicket(ticket.ID, ticket); err != nil {
					fmt.Printf("Error updating ticket %s: %v\n", ticket.ID, err)
				} else {
					fmt.Printf("\n%s updated with duration %s", ticket.ID, ticket.PrevDuration)
				}
			} else {
				fmt.Print("Parsed Time Ago:")
				fmt.Println(utils.ParseTimeAgo(ticket.PrevDuration))
			}
		}
	}
}
