package helperrouter

import (
	"github.com/gofiber/fiber/v2"

	helperhandler "github.com/meQlause/go-be-did/internal/delivery/http/handler/helper"
	internalhelper "github.com/meQlause/go-be-did/internal/infrastructure/helper"
	helperuc "github.com/meQlause/go-be-did/internal/usecase/helper"
)

func SetupHelperRoutes(api fiber.Router) {
	helperRepo, err := internalhelper.NewInternalHelper()
	if err != nil {
		panic(err)
	}

	helperUC := helperuc.New(helperRepo)
	helperHandler := helperhandler.NewHelperHandler(helperUC)
	helper := api.Group("/helper")

	helper.Post("/string-2-hex32", helperHandler.StringToHex32)

	// DID encoding endpoints
	helper.Post("/encode-create-did-param", helperHandler.EncodeCreateDIDParam)
	helper.Post("/encode-update-did-param", helperHandler.EncodeUpdateDIDParam)
	helper.Post("/encode-deactivate-did-param", helperHandler.EncodeDeactiveDIDParam)
	helper.Post("/encode-reactivate-did-param", helperHandler.EncodeReactiveDIDParam)
	helper.Post("/encode-transfer-did-owner-param", helperHandler.EncodeTransferDIDOwnerParam)

	// Data encoding endpoints
	helper.Post("/encode-store-data-param", helperHandler.EncodeStoreDataParam)
	helper.Post("/encode-delete-data-param", helperHandler.EncodeDeleteDataParam)

	// Key encoding endpoints
	helper.Post("/encode-add-key-param", helperHandler.EncodeAddKeyParam)
	helper.Post("/encode-remove-key-param", helperHandler.EncodeRemoveKeyParam)

	// Claim encoding endpoints
	helper.Post("/encode-add-claim-param", helperHandler.EncodeAddClaimParam)
	helper.Post("/encode-remove-claim-param", helperHandler.EncodeRemoveClaimParam)
}
