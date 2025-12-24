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

	helper.Post("/string-2-byte32", helperHandler.StringToByte32)
	helper.Post("/encode-create-did-param", helperHandler.EncodeCreateDIDParam)
}
