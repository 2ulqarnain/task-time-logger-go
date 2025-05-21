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
	vars.DB_FILENAME = os.Getenv("DB_FILENAME")
	vars.JIRA_BASE_URL = os.Getenv("JIRA_BASE_URL")
	vars.JIRA_USERNAME = os.Getenv("JIRA_USERNAME")
	vars.JIRA_API_TOKEN = os.Getenv("JIRA_API_TOKEN")
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", handlers.GetHomePage)

	apiGroup := app.Group("/api")
	tasks := apiGroup.Group("/tasks")
	projects := apiGroup.Group("/projects")

	tasks.Get("/", handlers.GetTasks)
	tasks.Get("/:"+params.TICKET_ID, handlers.GetTaskByID)
	tasks.Post("/:"+params.TICKET_ID, handlers.InitTaskTimeLog)
	projects.Get("/", handlers.GetAllProjectsKeys)

	app.Listen(":8080")
}
