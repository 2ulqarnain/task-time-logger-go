package handlers

import (
	"fmt"
	"path/filepath"
	"strings"
	"task-time-logger-go/utils"
	"task-time-logger-go/utils/enums/params"
	"task-time-logger-go/utils/out"
	"task-time-logger-go/utils/vars"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GetTasks(ctx *fiber.Ctx) error {
	FieldsRequired := [5]string{"summary", "created", "statuscategorychangedate", "statusCategory"}
	url := vars.JIRA_BASE_URL + "/rest/api/3/search?fields=" + strings.Join(FieldsRequired[:], ",")
	body, err := utils.CallJiraApi(url)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	data := gjson.GetBytes(body, "issues.#.{fields}")

	return ctx.JSON(data.Value())
}

func GetTaskByID(ctx *fiber.Ctx) error {
	ticketID := ctx.Params(params.TICKET_ID)
	url := vars.JIRA_BASE_URL + "/rest/api/3/issue/" + ticketID
	body, err := utils.CallJiraApi(url)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response := gjson.Get(string(body), "{errorMessages,key,fields.summary,statusUpdatedAt:fields.statuscategorychangedate,status:fields.statusCategory.name,statusId:fields.statusCategory.id}")
	return ctx.JSON(response.Value())
}

func InitTaskTimeLog(ctx *fiber.Ctx) error {
	ticketID := ctx.Params(params.TICKET_ID)
	if ticketID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Ticket ID not provided!")
	}

	var tickets []utils.Ticket = []utils.Ticket{
		{ID: ticketID, Title: "Sample Title", StartedOn: time.Now()},
	}

	utils.SaveTicketsToFile(tickets)

	return ctx.SendString("Ticket successfully posted!")
}

func GetAddedTasks(ctx *fiber.Ctx) error {
	filename := vars.DB_FILENAME

	if filename == "" {
		out.Errorln("No file name provided for retrieving data...proceeding with backup filename: data_backup.gob")
		filename = "data_backup.gob"
	}

	filePath := filepath.Join("db", filename)

	db, err := utils.LoadTickets(filePath)

	if err != nil {
		out.Errorln(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal server error while retrieving files from db!")
	}

	type TicketResponse struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		StartedOn time.Time `json:"startedOn"`
		Age       string    `json:"age"`
	}

	var response []TicketResponse
	for _, ticket := range db.Tickets {
		response = append(response, TicketResponse{
			ID:        ticket.ID,
			Title:     ticket.Title,
			StartedOn: ticket.StartedOn,
			Age:       utils.TimeAgo(ticket.StartedOn),
		})
	}

	return ctx.JSON(response)
}
