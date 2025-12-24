package didrootrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupAccountAbstractionRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	api.Group("/did-root")

}
