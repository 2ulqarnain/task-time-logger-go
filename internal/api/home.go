package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	return c.SendString("Task Time Logger API")
}
