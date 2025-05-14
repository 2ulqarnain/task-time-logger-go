package handlers

import (
	"fmt"
	"task-time-logger-go/utils"
	"task-time-logger-go/utils/vars"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GetAllProjectsKeys(ctx *fiber.Ctx) error {
	url := vars.JIRA_BASE_URL + "/rest/api/3/project/search"

	body, err := utils.CallJiraApi(url)
	if err != nil {
		fmt.Println(err)
		return ctx.SendString(err.Error())
	}

	gjsonValue := gjson.Get(string(body), "values.#.key")

	var keys []string

	for _, item := range gjsonValue.Array() {
		keys = append(keys, item.String())
	}

	if len(keys) == 0 {
		return ctx.Status(fiber.StatusNotFound).SendString("No Projects!")
	}

	return ctx.JSON(keys)
}
