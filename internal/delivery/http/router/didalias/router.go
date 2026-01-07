package didaliasrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	didaliassdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didalias"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupDIDAliasRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	da := api.Group("/did-alias")

	SetupAliasFactoryRoutes(da, didaliassdk.GetDIDAliasSDK(), cfg, bc)
}
