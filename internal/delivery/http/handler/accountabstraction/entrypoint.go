package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	didrootevent "github.com/meQlause/go-be-did/internal/delivery/event/didroot"
	"github.com/meQlause/go-be-did/pkg/response"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
)

func (ah *AccountAbstractionHandler) HandleOps(c *fiber.Ctx) error {
	var input aado.HandleOpsInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	aasdk.ChangeWalletImplementationAddress(input.Wallet, config.Blockchain())

	result, err := ah.uc.HandleOps(c.Context(), input, config.Network())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	txSuccess, txErrors := config.Blockchain().CheckTxs(c.Context(), result.TxHash)

	resp := make(map[string]Response)

	for txHash, ok := range txSuccess {
		eventData, err := didrootevent.DecodeCreateDIDEvents(c.Context(), txHash)
		if err != nil {
			continue
		}

		resp[txHash.Hex()] = Response{
			Success:  ok,
			Errors:   "No Error Message",
			Returned: eventData,
		}
	}

	for txHash, txErr := range txErrors {
		resp[txHash.Hex()] = Response{
			Success:  false,
			Errors:   txErr.Error(),
			Returned: nil,
		}
	}

	return response.Success(c, resp)
}
