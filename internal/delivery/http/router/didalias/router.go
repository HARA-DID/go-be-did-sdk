package didaliasrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	didaliassdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didalias"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupDIDAliasRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	dr := api.Group("/did-root")

	SetupAliasFactoryRoutes(dr, didaliassdk.GetDIDAliasSDK(), cfg, bc)
}
