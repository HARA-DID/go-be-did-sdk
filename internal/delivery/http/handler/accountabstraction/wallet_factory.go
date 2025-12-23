package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"
	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/pkg/response"
)

func (ah *AccountAbstractionHandler) CreateAccount(c *fiber.Ctx) error {
	var input aado.CreateAccountInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	result, err := ah.uc.CreateAccount(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	txSuccess, txErrors := config.Blockchain().
		CheckTxs(c.Context(), result.TxHash)

	resp := TxCheckResponse{
		Success: make(map[string]bool),
		Errors:  make(map[string]string),
	}

	for h, ok := range txSuccess {
		resp.Success[h.Hex()] = ok
		resp.Errors[h.Hex()] = "No Error Message"
	}

	for h, err := range txErrors {
		if err != nil {
			resp.Errors[h.Hex()] = err.Error()
		}
	}

	return response.Success(c, resp)
}
