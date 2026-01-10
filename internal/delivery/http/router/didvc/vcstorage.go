package didvcrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didvchandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/didvc"
	didvccase "github.com/meQlause/go-be-did/internal/usecase/didvc"
)

func SetupVCStorageRoutes(dr fiber.Router, didvcRepo repository.DIDVCRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	didVCUC := didvccase.New(didvcRepo)
	didVCHandler := didvchandler.NewDIDVCHandler(didVCUC)

	// VC Storage endpoints
	dr.Get("/get-identity-token-count", didVCHandler.GetIdentityTokenCount)
	dr.Get("/get-certificate-token-count", didVCHandler.GetCertificateTokenCount)
	dr.Get("/get-identity-token-ids", didVCHandler.GetIdentityTokenIds)
	dr.Get("/get-certificate-token-ids", didVCHandler.GetCertificateTokenIds)
	dr.Get("/get-all-identity-token-ids", didVCHandler.GetAllIdentityTokenIds)
	dr.Get("/get-all-certificate-token-ids", didVCHandler.GetAllCertificateTokenIds)
	dr.Get("/get-did-root-storage", didVCHandler.GetDIDRootStorage)
}
