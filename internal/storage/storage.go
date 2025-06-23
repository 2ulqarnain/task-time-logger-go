package storage

import (
	"encoding/gob"
	"fmt"
	"os"
	"slices"
	"time"

	"task-time-logger-go/internal/config"
	"task-time-logger-go/internal/logger"
	"task-time-logger-go/internal/models/structs"
)

type Database struct {
	Tickets []structs.Ticket
}

var db Database

// Initialize loads the database from disk
func Initialize() {
	file, err := os.Open(config.AppConfig.DBFilename)
	if err != nil {
		// If file doesn't exist, create an empty database
		if os.IsNotExist(err) {
			db = Database{Tickets: []structs.Ticket{}}
			if err := SaveTickets(); err != nil {
				logger.AppLogger.Fatalf("Could not initialize database storage, Error: %v", err)
			}
			logger.AppLogger.Println("Created new database file !")
			return
		} else {
			logger.AppLogger.Fatal(err)
		}
	}
	fi, err := file.Stat()
	if err != nil {
		logger.AppLogger.Printf("Failed to get file stats: %v", err)
	}
	fmt.Printf("%sDatabase file size: %s%d bytes \n", logger.ColorGray, logger.ColorReset, fi.Size())
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&db); err != nil {
		logger.AppLogger.Fatalln(err)
	}
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

func GetAllTasks() []structs.Ticket {
	return db.Tickets
}

func GetTaskByID(ticketID string) *structs.Ticket {
	for _, ticket := range db.Tickets {
		if ticket.ID == ticketID {
			return &ticket
		}
	}
	return nil
}

func InitTaskTimeById(ticketID string, ticketTitle string, ticketStatus string) *structs.Ticket {
	ticket := &structs.Ticket{
		ID:        ticketID,
		Title:     ticketTitle,
		Status:    ticketStatus,
		StartedOn: time.Now(),
		EndedOn:   structs.NullTime(time.Time{}),
	}
	db.Tickets = append(db.Tickets, *ticket)
	SaveTickets()
	return ticket
}

func DeleteAllTasks() (int, error) {
	ticketsCount := len(db.Tickets)
	db.Tickets = []structs.Ticket{}
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

// UpdateTicket updates an existing ticket by ID
func UpdateTicket(ticketID string, updatedTicket structs.Ticket) error {
	for i, ticket := range db.Tickets {
		if ticket.ID == ticketID {
			db.Tickets[i] = updatedTicket
			return SaveTickets()
		}
	}
	return fmt.Errorf("ticket with ID %s not found", ticketID)
}
