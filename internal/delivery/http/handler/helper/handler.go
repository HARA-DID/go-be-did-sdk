package helperhandler

import (
	"github.com/gofiber/fiber/v2"
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
	helperuc "github.com/meQlause/go-be-did/internal/usecase/helper"
	"github.com/meQlause/go-be-did/pkg/response"
)

type HelperHandler struct {
	uc *helperuc.HelperUseCase
}

func NewHelperHandler(uc *helperuc.HelperUseCase) *HelperHandler {
	return &HelperHandler{uc: uc}
}

func (hh *HelperHandler) StringToByte32(c *fiber.Ctx) error {
	var input helperdo.StringToByte32Input
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	resp := hh.uc.StringToByte32(input.Input)

	return response.Success(c, resp)
}
