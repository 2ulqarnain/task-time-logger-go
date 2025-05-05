package handlers

import (
	"task-time-logger-go/db"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(ctx *fiber.Ctx) error {
	rows, err := db.DB.Query("SELECT * FROM tickets;")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer rows.Close()
	tickets := []db.Ticket{}
	for rows.Next() {
		var ticket db.Ticket
		if err := rows.Scan(&ticket.ID, &ticket.Title, &ticket.TicketNo); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		tickets = append(tickets, ticket)
	}
	return ctx.JSON(tickets)
}

func PostTask(ctx *fiber.Ctx) error {
	return ctx.SendString("<html><p>Task Posted <b>Successfully!</b></p></html>")
}
