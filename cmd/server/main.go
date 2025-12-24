package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/delivery/http/router"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func main() {
	config.InitConfig()
	config.InitBlockchain()

	logger.Init(config.GetApp().LogLevel)

	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	router.Setup(app, config.GetConfig(), config.Blockchain())

	log.Printf("Server starting on port %s", config.GetApp().Port)
	log.Fatal(app.Listen(":" + config.GetApp().Port))
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
