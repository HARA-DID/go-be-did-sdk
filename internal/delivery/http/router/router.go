package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	aarouter "github.com/meQlause/go-be-did/internal/delivery/http/router/accountabstraction"
	helperrouter "github.com/meQlause/go-be-did/internal/delivery/http/router/helper"
)

func Setup(app *fiber.App, cfg *config.Config, bc *blockchain.Blockchain) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api/" + config.GetApp().Version)

	api.Get("/swagger/*", swagger.HandlerDefault)

	aarouter.SetupAccountAbstractionRoutes(api, cfg, bc)
	helperrouter.SetupHelperRoutes(api)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "rpc-api-gateway",
			"modules": []string{
				"didroot",
				"account-abstraction",
				"alias",
				"credentials",
			},
		})
	})
}
