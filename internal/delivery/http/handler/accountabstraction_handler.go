package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/response"
)

type AccountAbstractionHandler struct {
	uc *aauc.UseCase
}

func NewAccountAbstractionHandler(uc *aauc.UseCase) *AccountAbstractionHandler {
	return &AccountAbstractionHandler{uc: uc}
}

func (h *AccountAbstractionHandler) CreateAccount(c *fiber.Ctx) error {
	var input accountabstraction.CreateAccountInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	result, err := h.uc.CreateAccount(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, result)
}

func (h *AccountAbstractionHandler) ExecuteOperation(c *fiber.Ctx) error {
	var input accountabstraction.ExecuteOperationInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := h.uc.ExecuteOperation(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, result)
}
