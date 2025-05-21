package handlers

import (
	"fmt"
	"strings"
	"task-time-logger-go/utils"
	"task-time-logger-go/utils/enums/params"
	"task-time-logger-go/utils/vars"

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
		{ID: "123", Title: "Sample Title", StartedOn: "2025-05-10T12:00"},
	}

	utils.SaveTicketsToFile(tickets)

	return ctx.SendString("Ticket successfully posted!")
}
