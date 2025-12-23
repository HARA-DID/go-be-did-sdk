package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/response"
)

type TxCheckResponse struct {
	Success map[string]bool   `json:"success"`
	Errors  map[string]string `json:"errors,omitempty"`
}

type AccountAbstractionHandler struct {
	uc *aauc.UseCase
}

func NewAccountAbstractionHandler(uc *aauc.UseCase) *AccountAbstractionHandler {
	return &AccountAbstractionHandler{uc: uc}
}

func (h *AccountAbstractionHandler) CreateAccount(c *fiber.Ctx) error {
	var input aado.CreateAccountInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	result, err := h.uc.CreateAccount(c.Context(), input)
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

func (h *AccountAbstractionHandler) ExecuteOperation(c *fiber.Ctx) error {
	var input aado.ExecuteOperationInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.uc.ExecuteOperation(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, result)
}
