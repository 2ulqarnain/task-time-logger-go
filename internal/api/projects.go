package api

import (
	"task-time-logger-go/internal/models/structs"
	"task-time-logger-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllProjectsKeys(c *fiber.Ctx) error {
	projects, err := services.GetJiraProjects()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	if len(projects) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Project Found",
		})
	}
	return c.JSON(structs.ApiResponse(false, "Projects successfully fetched!", projects))
}
