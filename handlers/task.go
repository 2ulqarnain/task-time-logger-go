package handlers

import (
	"encoding/json"
	"fmt"
	"task-time-logger-go/utils"
	"task-time-logger-go/utils/vars"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GetTasks(ctx *fiber.Ctx) error {
	url := vars.JIRA_BASE_URL + "/rest/api/3/issue/SMS-1"
	body, err := utils.CallJiraApi(url)
	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error !")
	}
	response := gjson.Get(string(body), "key,id")
	fmt.Println(response)
	return ctx.JSON(data)
}
