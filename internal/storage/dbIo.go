package storage

import (
	"encoding/gob"
	"os"
	"path/filepath"
	"task-time-logger-go/utils/out"
	"task-time-logger-go/utils/vars"
	"time"
)

type TicketDB struct {
	Tickets map[string]Ticket
}

func getDBPath() string {
	filename := vars.DB_FILENAME
	if filename == "" {
		out.Errorln("No file name provided...using backup filename: data_backup.gob")
		filename = "data_backup.gob"
	}
	return filepath.Join("db", filename)
}

func AddNewTicket(Id string, Title string) error {
	newTicket := Ticket{
		ID:        Id,
		Title:     Title,
		StartedOn: time.Now(),
	}

	db, err := LoadTickets()
	if err != nil {
		return err
	}

	db.addTicketIfNotExists(newTicket)

	return nil
}

func LoadTickets() (*TicketDB, error) {
	db := &TicketDB{
		Tickets: make(map[string]Ticket),
	}

	filename := getDBPath()
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return db, nil
		}
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func saveTickets(db *TicketDB) error {
	filename := getDBPath()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(db)
}

func (db *TicketDB) addTicketIfNotExists(ticket Ticket) bool {
	if _, exists := db.Tickets[ticket.ID]; exists {
		return false
	}
	db.Tickets[ticket.ID] = ticket
	saveTickets(db)
	return true
}

func (db *TicketDB) DeleteTicket(ticketID string) bool {
	if _, exists := db.Tickets[ticketID]; exists {
		delete(db.Tickets, ticketID)
		saveTickets(db)
		return true
	}
	return false
}

func (db *TicketDB) DeleteAllTickets() error {
	db.Tickets = make(map[string]Ticket)

	if err := saveTickets(db); err != nil {
		return err
	}
	return nil
}
