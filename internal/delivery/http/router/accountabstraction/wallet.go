package accountabstractionrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	aahandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/accountabstraction"
	"github.com/meQlause/go-be-did/internal/repository"
	aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupWalletRoutes(aa fiber.Router, aaRepo repository.AccountAbstractionRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	aaUC := aauc.New(aaRepo)
	aaHandler := aahandler.NewAccountAbstractionHandler(aaUC)

	aa.Post("/execute", aaHandler.ExecuteOperation)
	// aa.Get("/:address", aaHandler.GetAccountInfo)
}
