package didrootrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	didroothandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/didroot"
	didrootusecase "github.com/meQlause/go-be-did/internal/usecase/didroot"
)

func SetupRootStorageRoutes(dr fiber.Router, didRootRepo repository.DIDRootRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	didRootUC := didrootusecase.New(didRootRepo)
	didRootHandler := didroothandler.NewDIDRootHandler(didRootUC)

	dr.Get("/resolve-did", didRootHandler.ResolveDID)
	dr.Get("/verify-did-ownership", didRootHandler.VerifyDIDOwnership)
	dr.Get("/get-key", didRootHandler.GetKey)
	dr.Get("/get-keys-by-did", didRootHandler.GetKeysByDID)
	dr.Get("/get-claim", didRootHandler.GetClaim)
	dr.Get("/get-claims-by-did", didRootHandler.GetClaimsByDID)
	dr.Get("/verify-claim", didRootHandler.VerifyClaim)
	dr.Get("/get-data", didRootHandler.GetData)
	dr.Get("/get-did-key-data-count", didRootHandler.GetDIDKeyDataCount)
	dr.Get("/get-did-key-data-by-index", didRootHandler.GetDIDKeyDataByIndex)
	dr.Get("/get-original-key", didRootHandler.GetOriginalKey)
	dr.Get("/did-index-map", didRootHandler.DIDIndexMap)
	dr.Get("/did-index-map-reverse", didRootHandler.DIDIndexMapReverse)
}
