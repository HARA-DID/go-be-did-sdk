package didvcrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didvcsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didvc"
)

func SetupDIDVCRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	dv := api.Group("/did-vc")

	SetupNFTBaseRoutes(dv, didvcsdk.GetDIDVCSDK(), cfg, bc)
	SetupVCStorageRoutes(dv, didvcsdk.GetDIDVCSDK(), cfg, bc)
}
