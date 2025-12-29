package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/pkg/response"

	aaevent "github.com/meQlause/go-be-did/internal/delivery/event/accountabstraction"
	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
)

// CreateWallet godoc
// @Summary      Create Account Abstraction Wallet
// @Tags         account-abstraction
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateWalletRequest true "Create wallet payload"
// @Success      200 {object} map[string]Response
// @Router       /account-abstraction/create [post]
func (ah *AccountAbstractionHandler) CreateWallet(c *fiber.Ctx) error {
	var input aado.CreateWalletInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := ah.uc.CreateWallet(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	txSuccess, txErrors := config.Blockchain().CheckTxs(c.Context(), result.TxHash)

	resp := make(map[string]Response)

	for txHash, ok := range txSuccess {
		eventData, err := aaevent.DecodeWalletDeployedEvent(c.Context(), txHash)
		if err != nil {
			continue
		}
		if eventData != nil {
			resp[txHash.Hex()] = Response{
				Success:  ok,
				Errors:   "No Error Message",
				Returned: eventData,
			}
		}
	}

	for h, err := range txErrors {
		resp[h.Hex()] = Response{
			Success:  false,
			Errors:   err.Error(),
			Returned: nil,
		}
	}

	return response.Success(c, resp)
}
