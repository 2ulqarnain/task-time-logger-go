package utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"task-time-logger-go/utils/out"
	"task-time-logger-go/utils/vars"
	"time"
)

type Ticket struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	StartedOn time.Time `json:"statedOn"`
}

type TicketDB struct {
	Tickets map[string]Ticket
}

func LoadTickets(filename string) (*TicketDB, error) {
	db := &TicketDB{
		Tickets: make(map[string]Ticket),
	}

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

func saveTickets(filename string, db *TicketDB) error {
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
	return true
}

func SaveTicketsToFile(tickets []Ticket) {
	filename := vars.DB_FILENAME

	if filename == "" {
		out.Errorln("No file name provided for saving data...proceeding with backup filename: data_backup.gob")
		filename = "data_backup.gob"
	}

	filePath := filepath.Join("db", filename)

	db, err := LoadTickets(filePath)
	if err != nil {
		fmt.Printf("Error loading tickets: %v\n", err)
		return
	}

	for _, ticket := range tickets {
		if db.addTicketIfNotExists(ticket) {
			fmt.Printf("Added new ticket: %s\n", ticket.ID)
		} else {
			out.Warningf("Ticket %s already exists, skipping\n", ticket.ID)
		}
	}

	err = saveTickets(filePath, db)
	if err != nil {
		fmt.Printf("Error saving tickets: %v\n", err)
		return
	}

	out.Successln("Ticket database updated successfully!")
}
