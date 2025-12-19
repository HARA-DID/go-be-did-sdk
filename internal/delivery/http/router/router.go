package router

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/delivery/http/handler"
	"github.com/meQlause/go-be-did/internal/infrastructure/sdk"
	"github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg"
)

func Setup(app *fiber.App, cfg *config.Config, bc *pkg.Blockchain) {
	// Middleware

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Initialize repositories (Infrastructure Layer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30))
	defer cancel()

	aaRepo, _ := sdk.NewAccountAbstractionSDK(ctx, cfg.HNS.AccountAbstraction, bc)

	// Initialize use cases (Business Logic Layer) - ONE instance per module
	aaUC := accountabstraction.New(aaRepo)

	// Initialize handlers (Delivery Layer)
	aaHandler := handler.NewAccountAbstractionHandler(aaUC)

	// API routes
	api := app.Group("/api/v1")

	// Account Abstraction routes
	aa := api.Group("/account-abstraction")
	aa.Post("/create", aaHandler.CreateAccount)
	aa.Post("/execute", aaHandler.ExecuteOperation)
	// aa.Get("/:address", aaHandler.GetAccountInfo)

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
