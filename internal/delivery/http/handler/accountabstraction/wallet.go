package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"
	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/response"
)

func (ah *AccountAbstractionHandler) ExecuteOperation(c *fiber.Ctx) error {
	var input aado.ExecuteOperationInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := ah.uc.ExecuteOperation(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, result)
}
