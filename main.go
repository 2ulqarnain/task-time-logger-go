package main

import (
	"task-time-logger-go/internal/api"
	"task-time-logger-go/internal/config"
	"task-time-logger-go/internal/models/params"
	"task-time-logger-go/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := config.Load(); err != nil {
		panic("Couldn't Load configuration!")
	}

	if err := storage.Initialize(); err != nil {
		panic("Couldn't initialize storage!")
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", api.GetHomePage)

	apiGroup := app.Group("/api")
	tasks := apiGroup.Group("/tasks")
	projects := apiGroup.Group("/projects")

	tasks.Get("/", api.GetTasks)
	tasks.Get("/added/", api.GetAddedTasks)
	tasks.Get("/:"+params.TICKET_ID, api.GetTaskByID)
	tasks.Post("/:"+params.TICKET_ID, api.InitTaskTimeById)
	tasks.Delete("/all", api.DeleteAllTasks)
	tasks.Delete("/:"+params.TICKET_ID, api.DeleteTaskById)
	projects.Get("/", api.GetAllProjectsKeys)

	app.Listen(":8080")
}
