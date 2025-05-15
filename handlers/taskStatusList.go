package handlers

import "github.com/gofiber/fiber/v2"

func GetTaskStatusList(ctx *fiber.Ctx) error {
	Statuses := map[int]string{
		2: "To Do",
		4: "In Progress",
	}

	return ctx.JSON(Statuses)
}
