package storage

import (
	"encoding/gob"
	"os"
	"time"

	"task-time-logger-go/internal/config"
)

type Ticket struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	StartedOn time.Time `json:"startedOn"`
}

type Database struct {
	Tickets []Ticket
}

var db Database

// Initialize loads the database from disk
func Initialize() error {
	file, err := os.Open(config.AppConfig.DBFilename)
	if err != nil {
		// If file doesn't exist, create an empty database
		if os.IsNotExist(err) {
			db = Database{Tickets: []Ticket{}}
			return SaveTickets()
		}
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	return decoder.Decode(&db)
}

func SaveTickets() error {
	file, err := os.Create(config.AppConfig.DBFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(db)
}

func GetAllTasks() []Ticket {
	return db.Tickets
}

func GetAddedTasks() []Ticket {
	return db.Tickets
}

func GetTaskByID(ticketID string) *Ticket {
	for _, ticket := range db.Tickets {
		if ticket.ID == ticketID {
			return &ticket
		}
	}
	return nil
}

func InitTaskTimeById(ticketID string) *Ticket {
	ticket := &Ticket{
		ID:        ticketID,
		Title:     "Sample Title", // TODO: Get from JIRA
		StartedOn: time.Now(),
	}
	db.Tickets = append(db.Tickets, *ticket)
	SaveTickets()
	return ticket
}

func DeleteAllTasks() error {
	db.Tickets = []Ticket{}
	return SaveTickets()
}

func DeleteTaskById(ticketID string) error {
	for i, ticket := range db.Tickets {
		if ticket.ID == ticketID {
			db.Tickets = append(db.Tickets[:i], db.Tickets[i+1:]...)
			return SaveTickets()
		}
	}
	return nil
}
