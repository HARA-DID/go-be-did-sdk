package accountabstractionrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupAccountAbstractionRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	aa := api.Group("/account-abstraction")

	SetupWalletFactoryRoutes(aa, aasdk.GetAccountAbstractionSDK(), cfg, bc)
	SetupWalletRoutes(aa, aasdk.GetAccountAbstractionSDK(), cfg, bc)
}
