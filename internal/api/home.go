package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString("<b>Hello World</b>")
}
