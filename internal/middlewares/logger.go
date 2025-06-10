package middlewares

import (
	"fmt"
	"task-time-logger-go/internal/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {

	start := time.Now()

	err := c.Next()

	duration := time.Since(start)
	fmt.Printf("%s[%s] %s %s - %v%s\n", logger.ColorGray, c.Method(), c.Path(), c.IP(), duration, logger.ColorReset)

	return err
}
