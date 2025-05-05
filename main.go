package main

import (
	"log"
	"task-time-logger-go/db"
	"task-time-logger-go/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file, Error:%s", err.Error())
	}
	db.InitDB()
	defer db.DB.Close()
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return ctx.SendString("<html><div style='font-family:sans-serif;display:grid;place-items:center;width:100vw;height:100vh;'><p>Go Fiber - <b>Task Time Logger</b></p></html>")
	})

	apiGroup := app.Group("/api")
	tasksGroup := apiGroup.Group("/tasks")

	tasksGroup.Get("/", handlers.GetTasks)

	app.Listen(":8080")
}
