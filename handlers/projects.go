package handlers

import (
	"task-time-logger-go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func GetAllProjectsKeys(ctx *fiber.Ctx) error {
	url := "https://zulqarnainhaider.atlassian.net/rest/api/3/project/search"

	body, err := utils.CallJiraApi(url)

	if err != nil {
		return ctx.SendString(err.Error())
	}

	gjsonValue := gjson.Get(string(body), "values.#.key")

	var keys []string

	for _, item := range gjsonValue.Array() {
		keys = append(keys, item.String())
	}

	return ctx.JSON(keys)
}
