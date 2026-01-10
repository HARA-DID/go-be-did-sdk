package didvcrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didvchandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/didvc"
	didvccase "github.com/meQlause/go-be-did/internal/usecase/didvc"
)

func SetupVCStorageRoutes(dv fiber.Router, didvcRepo repository.DIDVCRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	didVCUC := didvccase.New(didvcRepo)
	didVCHandler := didvchandler.NewDIDVCHandler(didVCUC)

	dv.Get("/get-identity-token-count", didVCHandler.GetIdentityTokenCount)
	dv.Get("/get-certificate-token-count", didVCHandler.GetCertificateTokenCount)
	dv.Get("/get-identity-token-ids", didVCHandler.GetIdentityTokenIds)
	dv.Get("/get-certificate-token-ids", didVCHandler.GetCertificateTokenIds)
	dv.Get("/get-all-identity-token-ids", didVCHandler.GetAllIdentityTokenIds)
	dv.Get("/get-all-certificate-token-ids", didVCHandler.GetAllCertificateTokenIds)
	dv.Get("/get-did-root-storage", didVCHandler.GetDIDRootStorage)
}
