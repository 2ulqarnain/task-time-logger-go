package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Ticket struct {
	ID       int    `json:"id"`
	TicketNo string `json:"ticketNo"`
	Title    string `json:"title"`
}

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite3", os.Getenv("SQLITE_FILE_PATH"))
	if err != nil {
		log.Fatal(err)
	}
}
