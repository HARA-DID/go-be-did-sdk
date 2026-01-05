package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	accountabstractiondto "github.com/meQlause/go-be-did/internal/domain/dto/accountabstraction"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"

	didrootevent "github.com/meQlause/go-be-did/internal/delivery/event/handleops"
	aasdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
)

// HandleOps godoc
// @Summary      Handle User Operations
// @Description  Executes user operations through the Account Abstraction entry point. This endpoint processes user operations by signing and submitting them to the blockchain via the entry point contract.
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true for HTTP 200
// @Description  - `data` (object): Map where each key is a transaction hash and value contains transaction result
// @Description  - `meta` (object): Contains timestamp and API version
// @Description
// @Description  ## Transaction Result Object
// @Description  Each transaction hash in the data map contains:
// @Description  - `success` (boolean): Indicates if the transaction was successful on the blockchain
// @Description  - `errors` (string): Error message if failed, or "No Error Message" if successful
// @Description  - `returned` (object|null): CreateDIDEvent object containing DID creation details, or null if transaction failed or event decoding failed
// @Description
// @Description  ## Common Error Scenarios
// @Description  ### Transaction Errors (success: false in transaction result):
// @Description  - "execution reverted" - Transaction reverted on the blockchain
// @Description  - "nonce too low" - Transaction nonce conflict
// @Description  - "transaction underpriced" - Gas price below network minimum
// @Description  - "invalid signature" - Transaction signing failed
// @Description  - "insufficient funds" - Relayer lacks funds for gas fees
// @Description
// @Description  ## Important Notes
// @Description  - HTTP 200 status does NOT guarantee operation success
// @Description  - Always check the `success` field within each transaction result in the data map
// @Description  - Multiple transaction hashes may be returned for batch operations
// @Description  - The PrivKey is used by the relayer to sign and submit the transaction
// @Tags         account-abstraction
// @Accept       json
// @Produce      json
// @Param        request body accountabstractiondto.HandleOpsDTO true "HandleOps payload with private key, wallet address, target address, data, and nonce"
// @Success      200 {object} response.Response{data=map[string]backendutils.Response} "Transaction(s) processed successfully - check individual transaction results"
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON, missing required fields, or invalid address format"
// @Failure      500 {object} response.Response "Internal server error - handle ops use case failed, network connectivity issues, RPC node errors, or transaction submission failure"
// @Router       /account-abstraction/handle-ops [post]
func (ah *AccountAbstractionHandler) HandleOps(c *fiber.Ctx) error {
	var input accountabstractiondto.HandleOpsDTO
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	handleOpsInput := input.Into()
	aasdk.ChangeWalletImplementationAddress(handleOpsInput.Wallet, config.Blockchain())

	result, err := ah.uc.HandleOps(c.Context(), handleOpsInput, config.Network())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	txSuccess, txErrors := config.Blockchain().CheckTxs(c.Context(), result.TxHash)

	resp := make(map[string]response.BlockchainResponse)
	decodeEventFunc, _ := didrootevent.Registry[input.Details.Service][input.Details.TxType]

	for txHash, ok := range txSuccess {
		eventData, err := decodeEventFunc(c.Context(), txHash)
		if err != nil {
			continue
		}

		resp[txHash.Hex()] = response.BlockchainResponse{
			Success:  ok,
			Errors:   "No Error Message",
			Returned: eventData,
		}
	}

	for txHash, txErr := range txErrors {
		resp[txHash.Hex()] = response.BlockchainResponse{
			Success:  false,
			Errors:   txErr.Error(),
			Returned: nil,
		}
	}

	return response.Success(c, resp)
}
