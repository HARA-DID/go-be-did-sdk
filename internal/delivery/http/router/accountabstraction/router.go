package accountabstractionrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupAccountAbstractionRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	aa := api.Group("/account-abstraction")
	SetupWalletFactoryRoutes(aa, cfg, bc)
	SetupWalletRoutes(aa, cfg, bc)
}
