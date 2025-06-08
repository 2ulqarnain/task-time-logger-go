package api

import (
	"task-time-logger-go/internal/models/params"
	"task-time-logger-go/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
	tasks := storage.GetAllTasks()
	return c.JSON(tasks)
}

func GetAddedTasks(c *fiber.Ctx) error {
	tasks := storage.GetAddedTasks()
	return c.JSON(tasks)
}

func GetTaskByID(c *fiber.Ctx) error {
	ticketID := c.Params(params.TICKET_ID)
	task := storage.GetTaskByID(ticketID)
	return c.JSON(task)
}

func InitTaskTimeById(c *fiber.Ctx) error {
	ticketID := c.Params(params.TICKET_ID)
	task := storage.InitTaskTimeById(ticketID)
	return c.JSON(task)
}

func DeleteAllTasks(c *fiber.Ctx) error {
	storage.DeleteAllTasks()
	return c.SendStatus(fiber.StatusOK)
}

func DeleteTaskById(c *fiber.Ctx) error {
	ticketID := c.Params(params.TICKET_ID)
	storage.DeleteTaskById(ticketID)
	return c.SendStatus(fiber.StatusOK)
}
