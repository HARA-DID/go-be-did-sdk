package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/meQlause/go-be-did/internal/config"
	aarouter "github.com/meQlause/go-be-did/internal/delivery/http/router/accountabstraction"
	helperrouter "github.com/meQlause/go-be-did/internal/delivery/http/router/helper"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func Setup(app *fiber.App, cfg *config.Config, bc *blockchain.Blockchain) {
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// API routes
	api := app.Group("/api/" + config.GetApp().Version)

	// Initialize module routes
	aarouter.SetupAccountAbstractionRoutes(api, cfg, bc)
	helperrouter.SetupHelperRoutes(api)

	// Health check
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
