package accountabstractionhandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/response"
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
		// Decode the UserOperationEvent (similar to WalletDeployedEvent)
		// eventData, err := aaevent.DecodeUserOperationEvent(c.Context(), txHash)

		// if err != nil {
		// fmt.Printf("Could not decode UserOperation event for tx %s: %v\n", txHash, err)
		// Still add to response even if event decoding fails
		resp[txHash.Hex()] = Response{
			Success:  ok,
			Errors:   fmt.Sprintf("Event decode error: %v", err),
			Returned: nil,
		}
		// continue
		// }

		resp[txHash.Hex()] = Response{
			Success:  ok,
			Errors:   "No Error Message",
			Returned: nil,
		}
	}

	// Process failed transactions
	for txHash, txErr := range txErrors {
		resp[txHash.Hex()] = Response{
			Success:  false,
			Errors:   txErr.Error(),
			Returned: nil,
		}
	}

	return response.Success(c, resp)
}
