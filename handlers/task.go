package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetTasks(ctx *fiber.Ctx) error {

	return ctx.SendString("work in progress...")
}
