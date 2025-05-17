package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GetHomePage(ctx *fiber.Ctx) error {

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	htmlFilePath := filepath.Join("res", "index.html")
	content, err := os.ReadFile(htmlFilePath)

	if err != nil {
		fmt.Println(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).SendString("Couldn't read the HTML file!\nWelcome, anyways! :)")
	}

	return ctx.SendString(string(content))

}
