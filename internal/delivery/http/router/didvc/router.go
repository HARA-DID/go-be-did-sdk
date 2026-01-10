package didvcrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didvcsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didvc"
)

func SetupDIDVCRoutes(api fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	dr := api.Group("/did-vc")

	SetupNFTBaseRoutes(dr, didvcsdk.GetDIDVCSDK(), cfg, bc)
	SetupVCStorageRoutes(dr, didvcsdk.GetDIDVCSDK(), cfg, bc)
}
