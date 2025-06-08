package api

import (
	"task-time-logger-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllProjectsKeys(c *fiber.Ctx) error {
	projects, err := services.GetJiraProjects()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(projects)
}
