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

	app.Get("/", handlers.GetHomePage)

	apiGroup := app.Group("/api")
	tasks := apiGroup.Group("/tasks")
	projects := apiGroup.Group("/projects")

	tasks.Get("/", handlers.GetTasks)
	tasks.Get("/:"+params.TICKET_ID, handlers.GetTaskByID)
	projects.Get("/", handlers.GetAllProjectsKeys)

	app.Listen(":8080")
}
