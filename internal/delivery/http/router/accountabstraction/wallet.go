package accountabstractionrouter

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	aahandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/accountabstraction"
	"github.com/meQlause/go-be-did/internal/infrastructure/sdk"
	aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

func SetupWalletRoutes(aa fiber.Router, cfg *config.Config, bc *blockchain.Blockchain) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	aaRepo, err := sdk.NewAccountAbstractionSDK(ctx, cfg.HNS.AccountAbstraction, bc)
	if err != nil {
		panic(err)
	}

	aaUC := aauc.New(aaRepo)
	aaHandler := aahandler.NewAccountAbstractionHandler(aaUC)
	aa.Post("/execute", aaHandler.ExecuteOperation)
	// aa.Get("/:address", aaHandler.GetAccountInfo)
}
