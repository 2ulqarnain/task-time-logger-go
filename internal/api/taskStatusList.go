package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetTaskStatusList(c *fiber.Ctx) error {
	statuses := []string{"To Do", "In Progress", "Done"}
	return c.JSON(statuses)
}
