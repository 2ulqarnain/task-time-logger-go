package main

import (
	"os"
	"task-time-logger-go/handlers"
	"task-time-logger-go/utils/enums/params"
	"task-time-logger-go/utils/vars"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Couldn't Load .env file!")
	}
	vars.JIRA_USERNAME = os.Getenv("JIRA_USERNAME")
	vars.JIRA_PASSWORD = os.Getenv("JIRA_API_TOKEN")
	vars.JIRA_BASE_URL = os.Getenv("JIRA_BASE_URL")
	// db.InitDB()
	// defer db.DB.Close()
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return ctx.SendString("<html><div style='font-family:sans-serif;display:grid;place-items:center;width:100vw;height:100vh;'><p>Go Fiber - <b>Task Time Logger</b></p></html>")
	})

	apiGroup := app.Group("/api")
	tasks := apiGroup.Group("/tasks/:" + params.TICKET_ID)
	projects := apiGroup.Group("/projects")

	tasks.Get("/", handlers.GetTasks)
	projects.Get("/", handlers.GetAllProjectsKeys)

	app.Listen(":8080")
}
