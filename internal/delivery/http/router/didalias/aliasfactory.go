package didaliasrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	dahandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/didalias"
	dauc "github.com/meQlause/go-be-did/internal/usecase/didalias"
)

func SetupAliasFactoryRoutes(da fiber.Router, daRepo repository.DIDAliasRepository, cfg *config.Config, bc *blockchain.Blockchain) {
	daUC := dauc.New(daRepo)
	daHandler := dahandler.NewDIDAliasHandler(daUC)

	// Read operations (GET methods)
	da.Get("/resolve", daHandler.Resolve)
	da.Get("/resolve-from-string", daHandler.ResolveFromString)
	da.Get("/status", daHandler.GetAliasStatus)
	da.Get("/status-from-string", daHandler.GetAliasStatusFromString)
	da.Get("/owner", daHandler.GetOwner)
	da.Get("/owner-from-string", daHandler.GetOwnerFromString)
	da.Get("/did", daHandler.GetDID)
	da.Get("/did-from-string", daHandler.GetDIDFromString)
	da.Get("/namehash", daHandler.Namehash)
	da.Get("/registration-period", daHandler.GetRegistrationPeriod)

	da.Post("/register-tld", daHandler.RegisterTLD)
}
