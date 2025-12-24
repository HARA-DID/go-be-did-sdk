package helperrouter

import (
	"github.com/gofiber/fiber/v2"
	helperhandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/helper"
	internalhelper "github.com/meQlause/go-be-did/internal/infrastructure/helper"
	helperuc "github.com/meQlause/go-be-did/internal/usecase/helper"
)

func SetupHelperRoutes(api fiber.Router) {
	// Initialize repository
	helperRepo, err := internalhelper.NewInternalHelper()
	if err != nil {
		// Handle error appropriately - log or panic depending on your needs
		panic(err)
	}

	// Initialize use case
	helperUC := helperuc.New(helperRepo)

	// Initialize handler
	helperHandler := helperhandler.NewHelperHandler(helperUC)

	// Helper routes
	helper := api.Group("/helper")
	helper.Post("/string-2-byte32", helperHandler.StringToByte32)
}
