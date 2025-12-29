package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/pkg/response"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
)

func (ah *AccountAbstractionHandler) ValidateUserOps(c *fiber.Ctx) error {
	var input aado.ValidateUserOpsInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	aasdk.ChangeWalletImplementationAddress(input.Wallet, config.Blockchain())
	err := aasdk.GetAccountAbstractionSDK().Wallet.ValidateUserOps(c.Context(), input.Input)

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "UserOp Valid")
}
