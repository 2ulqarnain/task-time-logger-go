package api

import (
	"fmt"
	"task-time-logger-go/internal/logger"
	"task-time-logger-go/internal/models/enums/params"
	"task-time-logger-go/internal/models/structs"
	"task-time-logger-go/internal/storage"
	"task-time-logger-go/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllTasks(c *fiber.Ctx) error {
	tasks := storage.GetAllTasks()
	type responseData struct {
		Count int              `json:"count"`
		Tasks []storage.Ticket `json:"tasks"`
	}
	return c.JSON(structs.ApiResponse(false, "All tasks fetched successfully", responseData{
		Count: len(tasks),
		Tasks: tasks,
	}))
}

func GetTaskByID(c *fiber.Ctx) error {
	ticketID := c.Params(params.TICKET_ID)
	task := storage.GetTaskByID(ticketID)
	taskWithDuration := struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		StartedOn time.Time `json:"startedOn"`
		Duration  string    `json:"duration"`
	}{
		ID:        task.ID,
		Title:     task.Title,
		StartedOn: task.StartedOn,
		Duration:  utils.TimeAgo(task.StartedOn),
	}
	logger.AppLogger.Printf("Task duration: %f", time.Since(task.StartedOn).Hours())
	return c.JSON(structs.ApiResponse(false, "Task fetched successfully", taskWithDuration))
}

func InitTaskTimeById(c *fiber.Ctx) error {
	var requestBody struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	logger.AppLogger.Printf("Received request body: %+v", requestBody)

	if requestBody.ID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Ticket ID is required",
		})
	}

	ticket := storage.InitTaskTimeById(requestBody.ID, requestBody.Title)
	return c.JSON(structs.ApiResponse(false, "Task started successfully", ticket))
}

func DeleteAllTasks(c *fiber.Ctx) error {
	ticketsCount, err := storage.DeleteAllTasks()
	if ticketsCount == 0 {
		return c.JSON(structs.ApiResponse(false, "No tasks to delete", nil))
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			structs.ApiResponse(true, "Failed to delete all tasks", nil),
		)
	}
	responseMessage := fmt.Sprintf("All %d tasks deleted successfully", ticketsCount)
	return c.JSON(structs.ApiResponse(false, responseMessage, nil))
}

func DeleteTaskById(c *fiber.Ctx) error {
	ticketID := c.Params(params.TICKET_ID)
	storage.DeleteTaskById(ticketID)
	return c.SendStatus(fiber.StatusOK)
}
