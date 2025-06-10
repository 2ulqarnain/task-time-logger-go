package main

import (
	"task-time-logger-go/internal/api"
	"task-time-logger-go/internal/config"
	"task-time-logger-go/internal/logger"
	"task-time-logger-go/internal/middlewares"
	"task-time-logger-go/internal/models/enums/params"
	"task-time-logger-go/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	if err := config.Load(); err != nil {
		logger.AppLogger.Fatalf("Couldn't Load configuration: %v", err)
	}

	if err := storage.Initialize(); err != nil {
		logger.AppLogger.Fatalf("Couldn't initialize storage, %sERROR: %s%v", logger.ColorLightRed, logger.ColorReset, err)
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(middlewares.LoggerMiddleware)

	app.Get("/", api.GetHomePage)

	apiGroup := app.Group("/api")
	tasks := apiGroup.Group("/tasks")
	projects := apiGroup.Group("/projects")

	tasks.Get("/", api.GetAllTasks)
	tasks.Post("/", api.InitTaskTimeById)
	tasks.Delete("/", api.DeleteAllTasks)
	tasks.Get("/:"+params.TICKET_ID, api.GetTaskByID)
	tasks.Delete("/:"+params.TICKET_ID, api.DeleteTaskById)
	projects.Get("/", api.GetAllProjectsKeys)

	logger.AppLogger.Fatal(app.Listen(":8080"))
}
