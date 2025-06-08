package types

type Ticket struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	StartedOn string `json:"startedOn"`
	EndedOn   string `json:"endedOn"`
	Duration  string `json:"duration"`
}

type TicketDB struct {
	Tickets map[string]Ticket
}
