package router

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/meQlause/go-be-did/internal/config"
	aahandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/accountabstraction"
	helperhandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/helper"
	internalhelper "github.com/meQlause/go-be-did/internal/infrastructure/helper"
	"github.com/meQlause/go-be-did/internal/infrastructure/sdk"
	aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
	helperuc "github.com/meQlause/go-be-did/internal/usecase/helper"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func Setup(app *fiber.App, cfg *config.Config, bc *blockchain.Blockchain) {
	// Middleware

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Initialize repositories (Infrastructure Layer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30))
	defer cancel()

	aaRepo, _ := sdk.NewAccountAbstractionSDK(ctx, cfg.HNS.AccountAbstraction, bc)
	helperRepo, _ := internalhelper.NewInternalHelper()

	// Initialize use cases (Business Logic Layer) - ONE instance per module
	aaUC := aauc.New(aaRepo)
	helperUC := helperuc.New(helperRepo)

	// Initialize handlers (Delivery Layer)
	aaHandler := aahandler.NewAccountAbstractionHandler(aaUC)
	helperHandler := helperhandler.NewHelperHandler(helperUC)

	// API routes
	api := app.Group("/api/v1")

	// Account Abstraction routes
	aa := api.Group("/account-abstraction")
	aa.Post("/create", aaHandler.CreateAccount)
	aa.Post("/execute", aaHandler.ExecuteOperation)
	// aa.Get("/:address", aaHandler.GetAccountInfo)

	helper := api.Group("/helper")
	helper.Post("/string-2-byte32", helperHandler.StringToByte32)

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
