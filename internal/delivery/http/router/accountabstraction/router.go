package accountabstractionrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
)

func SetupAccountAbstractionRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	aa := api.Group("/account-abstraction")

	SetupWalletFactoryRoutes(aa, aasdk.GetAccountAbstractionSDK(), cfg, bc)
	SetupEntryPointRoutes(aa, aasdk.GetAccountAbstractionSDK(), cfg, bc)
	SetupWalletRoutes(aa, aasdk.GetAccountAbstractionSDK(), cfg, bc)
}
