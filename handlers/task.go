package handlers

import (
	"fmt"
	"task-time-logger-go/utils"
	"task-time-logger-go/utils/enums/params"
	"task-time-logger-go/utils/vars"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GetTasks(ctx *fiber.Ctx) error {
	ticketID := ctx.Params(params.TICKET_ID)
	url := vars.JIRA_BASE_URL + "/rest/api/3/issue/" + ticketID
	body, err := utils.CallJiraApi(url)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response := gjson.Get(string(body), "{key,fields.summary,statusUpdatedAt:fields.statuscategorychangedate,status:fields.statusCategory.name,statusId:fields.statusCategory.id}")
	return ctx.JSON(response.Value())
}

func GetTaskByID(ctx *fiber.Ctx) error {
	ticketID := ctx.Params(params.TICKET_ID)
	url := vars.JIRA_BASE_URL + "/rest/api/3/issue/" + ticketID
	body, err := utils.CallJiraApi(url)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response := gjson.Get(string(body), "{key,fields.summary,statusUpdatedAt:fields.statuscategorychangedate,status:fields.statusCategory.name,statusId:fields.statusCategory.id}")
	return ctx.JSON(response.Value())
}
