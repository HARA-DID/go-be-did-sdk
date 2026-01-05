package accountabstractionhandler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/meQlause/go-be-did/internal/config"
	accountabstractiondto "github.com/meQlause/go-be-did/internal/domain/dto/accountabstraction"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"

	aaevent "github.com/meQlause/go-be-did/internal/delivery/event/accountabstraction"
)

// CreateWallet godoc
// @Summary      Create Account Abstraction Wallet
// @Description  Creates a new Account Abstraction wallet by deploying a smart contract wallet on the blockchain.
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
// @Description  - `returned` (object|null): WalletDeployedEvent object containing wallet details, or null if transaction failed or event decoding failed
// @Description
// @Description  ## WalletDeployedEvent Object
// @Description  - `wallet_address` (string): The deployed smart contract wallet address
// @Description  - `deployer_address` (string): The Ethereum address that deployed the wallet
// @Description  - `salt` (string): The salt value used for deterministic wallet address generation
// @Description
// @Description  ## Common Error Scenarios
// @Description  ### Transaction Errors (success: false in transaction result):
// @Description  - "execution reverted: insufficient funds for gas * price + value" - Deployer lacks ETH for gas fees
// @Description  - "execution reverted: wallet already exists" - Wallet with same salt already deployed
// @Description  - "execution reverted: invalid deployer" - Deployer address not authorized
// @Description  - "nonce too low" - Transaction nonce conflict
// @Description  - "transaction underpriced" - Gas price below network minimum
// @Description  - "invalid signature" - Transaction signing failed
// @Description
// @Description  ### Event Decode Failures (transaction success: true, returned: null):
// @Description  - Transaction succeeded but event parsing failed
// @Description  - Event signature mismatch or unexpected contract behavior
// @Description  - Wrong contract address in logs or insufficient log topics
// @Description
// @Description  ## Important Notes
// @Description  - HTTP 200 status does NOT guarantee wallet creation success
// @Description  - Always check the `success` field within each transaction result in the data map
// @Description  - Multiple transaction hashes may be returned for batch operations
// @Description  - Empty data object {} indicates no events were emitted
// @Description  - Salt must be unique per deployer address to avoid conflicts
// @Tags         account-abstraction
// @Accept       json
// @Produce      json
// @Param        request body accountabstractiondto.CreateWalletInputDTO true "Wallet creation payload with deployer address and optional salt value"
// @Success      200 {object} response.Response{data=map[string]response.BlockchainResponse} "Transaction(s) processed successfully - check individual transaction results" example(SuccessfulWalletCreation)
// @Success      200 {object} response.Response{data=map[string]response.BlockchainResponse} "Transaction failed on blockchain - returned in 200 response with error details" example(TransactionFailed)
// @Success      200 {object} response.Response{data=map[string]response.BlockchainResponse} "Multiple transactions with mixed results" example(MultipleTxMixedResults)
// @Success      200 {object} response.Response{data=map[string]response.BlockchainResponse} "Transaction succeeded but event decoding failed" example(EventDecodeFailed)
// @Success      200 {object} response.Response{data=map[string]response.BlockchainResponse} "No events emitted" example(EmptyResponse)
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON, missing required fields, or invalid Ethereum address format" example(BadRequest)
// @Failure      500 {object} response.Response "Internal server error - wallet creation use case failed, network connectivity issues, RPC node errors, or smart contract deployment failure" example(InternalServerError)
// @Router       /account-abstraction/create [post]
func (ah *AccountAbstractionHandler) CreateWallet(c *fiber.Ctx) error {
	var input accountabstractiondto.CreateWalletInputDTO
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	createWalletInput := input.Into()
	result, err := ah.uc.CreateWallet(c.Context(), createWalletInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	txSuccess, txErrors := config.Blockchain().CheckTxs(c.Context(), result.TxHash)
	resp := make(map[string]response.BlockchainResponse)

	for txHash, ok := range txSuccess {
		eventData, err := aaevent.DecodeWalletDeployedEvent(c.Context(), txHash)
		if err != nil {
			continue
		}

		if eventData != nil {
			resp[txHash.Hex()] = response.BlockchainResponse{
				Success:  ok,
				Errors:   "No Error Message",
				Returned: eventData,
			}
		}
	}

	for h, err := range txErrors {
		resp[h.Hex()] = response.BlockchainResponse{
			Success:  false,
			Errors:   err.Error(),
			Returned: nil,
		}
	}

	return response.Success(c, resp)
}
