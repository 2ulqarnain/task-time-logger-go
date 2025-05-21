package types

type Ticket struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	StatusChangeDate string `json:"statusChangeDate"`
	Status           string `json:"status"`
}

type TicketDB struct {
	Tickets map[string]Ticket
}
