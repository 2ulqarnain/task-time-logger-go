package structs

import "time"

type Ticket struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Status       string    `json:"status"`
	StartedOn    time.Time `json:"startedOn"`
	PrevDuration string    `json:"duration"` /* Duration will be in minutes */
	EndedOn      NullTime  `json:"endedOn"`
}
