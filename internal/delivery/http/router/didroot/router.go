package didrootrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
)

func SetupDIDRootRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	dr := api.Group("/did-root")

	SetupRootStorageRoutes(dr, didrootsdk.GetDIDRootSDK(), cfg, bc)
}
