package accountabstractionrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	aahandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/accountabstraction"
	aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
)

func SetupEntryPointRoutes(aa fiber.Router, aaRepo repository.AccountAbstractionRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	aaUC := aauc.New(aaRepo)
	aaHandler := aahandler.NewAccountAbstractionHandler(aaUC)

	aa.Post("/handle-ops", aaHandler.HandleOps)
	// aa.Get("/:address", aaHandler.GetAccountInfo)
}
