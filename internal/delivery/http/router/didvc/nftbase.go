package didvcrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didvchandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/didvc"
	didvccase "github.com/meQlause/go-be-did/internal/usecase/didvc"
)

func SetupNFTBaseRoutes(dr fiber.Router, didvcRepo repository.DIDVCRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	didVCUC := didvccase.New(didvcRepo)
	didVCHandler := didvchandler.NewDIDVCHandler(didVCUC)

	dr.Get("/get-metadata", didVCHandler.GetMetadata)
	dr.Get("/is-credential-valid", didVCHandler.IsCredentialValid)
	dr.Get("/get-credentials-with-metadata", didVCHandler.GetCredentialsWithMetadata)
	dr.Get("/get-unclaimed-token-id", didVCHandler.GetUnclaimedTokenId)
	dr.Get("/total-tokens-to-be-claimed-by-did", didVCHandler.TotalTokensToBeClaimedByDid)
	dr.Get("/get-to-be-claimed-tokens-by-did", didVCHandler.GetToBeClaimedTokensByDid)
	dr.Get("/is-approved-for-all", didVCHandler.IsApprovedForAll)
}
