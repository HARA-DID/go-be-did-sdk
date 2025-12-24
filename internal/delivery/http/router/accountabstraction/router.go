package accountabstractionrouter

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	aahandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/accountabstraction"
	"github.com/meQlause/go-be-did/internal/infrastructure/sdk"
	aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupAccountAbstractionRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	// Initialize repository
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	aaRepo, err := sdk.NewAccountAbstractionSDK(ctx, cfg.HNS.AccountAbstraction, bc)
	if err != nil {
		// Handle error appropriately - log or panic depending on your needs
		panic(err)
	}

	// Initialize use case
	aaUC := aauc.New(aaRepo)

	// Initialize handler
	aaHandler := aahandler.NewAccountAbstractionHandler(aaUC)

	// Account Abstraction routes
	aa := api.Group("/account-abstraction")
	aa.Post("/create", aaHandler.CreateAccount)
	aa.Post("/execute", aaHandler.ExecuteOperation)
	// aa.Get("/:address", aaHandler.GetAccountInfo)
}
