package didvcrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didvchandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/didvc"
	didvccase "github.com/meQlause/go-be-did/internal/usecase/didvc"
)

func SetupNFTBaseRoutes(dv fiber.Router, didvcRepo repository.DIDVCRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	didVCUC := didvccase.New(didvcRepo)
	didVCHandler := didvchandler.NewDIDVCHandler(didVCUC)

	dv.Get("/get-metadata", didVCHandler.GetMetadata)
	dv.Get("/is-credential-valid", didVCHandler.IsCredentialValid)
	dv.Get("/get-credentials-with-metadata", didVCHandler.GetCredentialsWithMetadata)
	dv.Get("/get-unclaimed-token-id", didVCHandler.GetUnclaimedTokenId)
	dv.Get("/total-tokens-to-be-claimed-by-did", didVCHandler.TotalTokensToBeClaimedByDid)
	dv.Get("/get-to-be-claimed-tokens-by-did", didVCHandler.GetToBeClaimedTokensByDid)
	dv.Get("/is-approved-for-all", didVCHandler.IsApprovedForAll)
}
