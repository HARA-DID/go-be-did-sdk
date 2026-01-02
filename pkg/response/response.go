package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
)

type Response struct {
	Success bool `json:"success"`
	Error   any  `json:"error,omitempty"`
	Data    any  `json:"data,omitempty"`
	Meta    Meta `json:"meta"`
}

type Meta struct {
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

func Success(c *fiber.Ctx, data any) error {
	return c.JSON(Response{
		Success: true,
		Data:    data,
		Meta: Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Version:   config.GetApp().Version,
		},
	})
}

func Error(c *fiber.Ctx, statusCode int, message any) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Error:   message,
		Meta: Meta{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Version:   "1.0",
		},
	})
}
