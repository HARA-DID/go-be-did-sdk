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
	didaliasrouter "github.com/meQlause/go-be-did/internal/delivery/http/router/didalias"
	didrootrouter "github.com/meQlause/go-be-did/internal/delivery/http/router/didroot"
	didvcrouter "github.com/meQlause/go-be-did/internal/delivery/http/router/didvc"
	helperrouter "github.com/meQlause/go-be-did/internal/delivery/http/router/helper"
)

func Setup(app *fiber.App, cfg *config.Config, bc *blockchain.Blockchain) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api/" + config.GetApp().Version)

	aarouter.SetupAccountAbstractionRoutes(api, cfg, bc)
	didrootrouter.SetupDIDRootRoutes(api, cfg, bc)
	didaliasrouter.SetupDIDAliasRoutes(api, cfg, bc)
	didvcrouter.SetupDIDVCRoutes(api, cfg, bc)
	helperrouter.SetupHelperRoutes(api)

	api.Get("/swagger/*", swagger.HandlerDefault)

	// Health godoc
	// @Summary      Health Check
	// @Description  Returns the health status of the API service. This endpoint provides information about the service status and available modules.
	// @Description
	// @Description  ## Response Structure
	// @Description  Success responses (HTTP 200) contain:
	// @Description  - `status` (string): Service status (always "ok" when service is running)
	// @Description  - `service` (string): Service name identifier
	// @Description  - `modules` (array): List of available service modules
	// @Description
	// @Description  ## Available Modules
	// @Description  - `didroot`: DID Root management module
	// @Description  - `account-abstraction`: Account Abstraction wallet module
	// @Description  - `alias`: Alias management module
	// @Description  - `credentials`: Credentials management module
	// @Tags         system
	// @Accept       json
	// @Produce      json
	// @Success      200 {object} map[string]interface{} "Service is healthy and operational"
	// @Router       /health [get]
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
