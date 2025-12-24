package accountabstractionrouter

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupAccountAbstractionRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	aa := api.Group("/account-abstraction")

	aaRepo, err := aasdk.NewAccountAbstractionSDK(ctx, cfg.HNS.AccountAbstraction, bc)
	if err != nil {
		panic(err)
	}

	SetupWalletFactoryRoutes(aa, aaRepo, cfg, bc)
	SetupWalletRoutes(aa, aaRepo, cfg, bc)
}
