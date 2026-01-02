package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/domain/dto"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"

	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
)

// ValidateUserOps godoc
// @Summary      Validate User Operation
// @Description  Validates a user operation for an Account Abstraction wallet. This endpoint checks if a user operation is valid before execution, ensuring the operation can be processed by the wallet contract.
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true for HTTP 200
// @Description  - `data` (string): Validation result message ("UserOp Valid")
// @Description  - `meta` (object): Contains timestamp and API version
// @Description
// @Description  ## UserOp Structure
// @Description  The UserOp object contains:
// @Description  - `Target` (string): The target contract address for the operation
// @Description  - `Value` (string): The amount of ETH to send (in wei, as hex string)
// @Description  - `Data` (string): The encoded function call data (hex string)
// @Description  - `ClientBlockNumber` (string): The client block number (as hex string)
// @Description  - `UserNonce` (string): The user operation nonce (as hex string)
// @Description  - `Signature` (string): The signature for the user operation (hex string)
// @Description
// @Description  ## Common Error Scenarios
// @Description  ### Validation Errors:
// @Description  - "invalid user operation" - UserOp structure is invalid
// @Description  - "signature verification failed" - Signature does not match the operation
// @Description  - "nonce mismatch" - UserNonce does not match expected value
// @Description  - "insufficient balance" - Wallet lacks sufficient funds
// @Description  - "invalid target address" - Target contract address is invalid
// @Description
// @Description  ## Important Notes
// @Description  - This endpoint only validates the operation, it does not execute it
// @Description  - The wallet implementation address is automatically updated before validation
// @Description  - Validation ensures the operation can be safely executed
// @Tags         account-abstraction
// @Accept       json
// @Produce      json
// @Param        request body dto.ValidateUserOpsDTO true "Validation payload with wallet address and UserOp object"
// @Success      200 {object} response.Response{data=string} "User operation is valid"
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON, missing required fields, or invalid address format"
// @Failure      500 {object} response.Response "Internal server error - validation failed, network connectivity issues, or RPC node errors"
// @Router       /account-abstraction/validate-userop [post]
func (ah *AccountAbstractionHandler) ValidateUserOps(c *fiber.Ctx) error {
	var input dto.ValidateUserOpsDTO
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	validateUserOpsParams := input.Into()
	aasdk.ChangeWalletImplementationAddress(validateUserOpsParams.Wallet, config.Blockchain())
	err := aasdk.GetAccountAbstractionSDK().Wallet.ValidateUserOps(c.Context(), validateUserOpsParams.Input)

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "UserOp Valid")
}
