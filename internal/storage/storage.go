package storage

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"

	"task-time-logger-go/internal/config"
	"task-time-logger-go/internal/logger"
)

// NullTime represents a nullable time value
type NullTime time.Time

// MarshalJSON implements custom JSON marshaling
func (nt NullTime) MarshalJSON() ([]byte, error) {
	t := time.Time(nt)
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t)
}

// GobEncode implements the gob.GobEncoder interface
func (nt NullTime) GobEncode() ([]byte, error) {
	t := time.Time(nt)
	return t.GobEncode()
}

// GobDecode implements the gob.GobDecoder interface
func (nt *NullTime) GobDecode(data []byte) error {
	var t time.Time
	if err := t.GobDecode(data); err != nil {
		return err
	}
	*nt = NullTime(t)
	return nil
}

type Ticket struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	StartedOn time.Time `json:"startedOn"`
	EndedOn   NullTime  `json:"endedOn"`
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
	fi, err := file.Stat()
	if err != nil {
		logger.AppLogger.Printf("Failed to get file stats: %v", err)
		return err
	}
	fmt.Printf("%sDatabase file size: %s%d bytes", logger.ColorGray, logger.ColorReset, fi.Size())
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

func GetTaskByID(ticketID string) *Ticket {
	for _, ticket := range db.Tickets {
		if ticket.ID == ticketID {
			return &ticket
		}
	}
	return nil
}

func InitTaskTimeById(ticketID string, ticketTitle string) *Ticket {
	ticket := &Ticket{
		ID:        ticketID,
		Title:     ticketTitle,
		StartedOn: time.Now(),
		EndedOn:   NullTime(time.Time{}),
	}
	db.Tickets = append(db.Tickets, *ticket)
	SaveTickets()
	return ticket
}

func DeleteAllTasks() (int, error) {
	ticketsCount := len(db.Tickets)
	db.Tickets = []Ticket{}
	return ticketsCount, SaveTickets()
}

func DeleteTaskById(ticketID string) error {
	for i, ticket := range db.Tickets {
		if ticket.ID == ticketID {
			db.Tickets = slices.Delete(db.Tickets, i, i+1)
			return SaveTickets()
		}
	}
	return nil
}
