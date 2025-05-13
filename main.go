package main

import (
	"task-time-logger-go/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// db.InitDB()
	// defer db.DB.Close()
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return ctx.SendString("<html><div style='font-family:sans-serif;display:grid;place-items:center;width:100vw;height:100vh;'><p>Go Fiber - <b>Task Time Logger</b></p></html>")
	})

	apiGroup := app.Group("/api")
	tasks := apiGroup.Group("/tasks")
	projects := apiGroup.Group("/projects")

	tasks.Get("/", handlers.GetTasks)
	projects.Get("/", handlers.GetAllProjectsKeys)

	app.Listen(":8080")
}
