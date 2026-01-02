package helperhandler

import (
	"context"
	"math/big"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/domain/dto"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	accountabstractionsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	helperuc "github.com/meQlause/go-be-did/internal/usecase/helper"
)

type HelperHandler struct {
	uc *helperuc.HelperUseCase
}

func NewHelperHandler(uc *helperuc.HelperUseCase) *HelperHandler {
	return &HelperHandler{uc: uc}
}

// StringToByte32 godoc
// @Summary      Convert String to Byte32
// @Description  Converts a string input to a 32-byte array (byte32). This is a utility endpoint commonly used for encoding data that needs to fit into a single bytes32 storage slot in smart contracts.
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true for HTTP 200
// @Description  - `data` (string): The hex-encoded 32-byte array result
// @Description  - `meta` (object): Contains timestamp and API version
// @Description
// @Description  ## Common Use Cases
// @Description  - Encoding strings for smart contract storage
// @Description  - Preparing data for keccak256 hashing
// @Description  - Converting identifiers to fixed-size byte arrays
// @Description
// @Description  ## Important Notes
// @Description  - Input strings longer than 32 bytes will be truncated
// @Description  - Shorter strings will be padded with zeros
// @Description  - The result is returned as a hex string with 0x prefix
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.StringToHex32DTO true "String input to convert to byte32"
// @Success      200 {object} response.Response{data=string} "Successfully converted string to byte32"
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON or missing required fields"
// @Failure      500 {object} response.Response "Internal server error - conversion failed"
// @Router       /helper/string-2-byte32 [post]
func (hh *HelperHandler) StringToHex32(c *fiber.Ctx) error {
	var input dto.StringToHex32DTO
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	stringToHex32Input := input.Into()
	resp := hh.uc.StringToHex32(stringToHex32Input)
	return response.Success(c, resp)
}

// EncodeCreateDIDParam godoc
// @Summary      Encode Create DID Parameters
// @Description  Encodes parameters for creating a DID (Decentralized Identifier) and prepares the transaction data. This endpoint encodes the DID creation parameters, retrieves the current nonce for the wallet, and returns the encoded call data along with the target contract address.
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true for HTTP 200
// @Description  - `data` (object): EncodeCreateDIDParamResponse containing:
// @Description    - `Data` (string): The hex-encoded call data for the DID creation transaction
// @Description    - `Target` (string): The target contract address (DID Root Factory address)
// @Description    - `Nonce` (uint64): The current nonce for the wallet address
// @Description  - `meta` (object): Contains timestamp and API version
// @Description
// @Description  ## CreateDIDParam Structure
// @Description  The DIDParam object contains:
// @Description  - `DID` (string): The Decentralized Identifier string to create
// @Description
// @Description  ## Common Error Scenarios
// @Description  - "Can Not Get Nonce" - Failed to retrieve nonce from the blockchain
// @Description  - "failed to encode DID" - Failed to encode the DID parameter
// @Description  - "failed to pack" - Failed to pack the function call data
// @Description
// @Description  ## Important Notes
// @Description  - The nonce is retrieved from the EntryPoint contract for the specified wallet address
// @Description  - The encoded data is ready to be used in a user operation
// @Description  - The KeyIdentifier is used to identify which key will sign the operation
// @Description  - The address must be a valid Ethereum address
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body helperdomain.EncodeCreateDIDParamInput true "Encode DID creation parameters with wallet address, DID parameter, and key identifier"
// @Success      200 {object} response.Response{data=helperhandler.EncodeCreateDIDParamResponse} "Successfully encoded DID creation parameters"
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON, missing required fields, or invalid address format"
// @Failure      500 {object} response.Response "Internal server error - encoding failed, nonce retrieval failed, or RPC node errors"
// @Router       /helper/encode-create-did-param [post]
func (hh *HelperHandler) EncodeCreateDIDParam(c *fiber.Ctx) error {
	var input dto.EncodeCreateDIDDTO
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	encodeCreateDIDInput := input.Into()

	decodedData, err := hh.uc.EncodeCreateDIDParam(encodeCreateDIDInput)
	if err != nil {
		response.Error(c, fiber.StatusInternalServerError, "Can Not Get Nonce")
	}

	walletAddress := utils.HexToAddress(input.Address)

	key := new(big.Int).SetInt64(0)
	nonce, err := accountabstractionsdk.GetAccountAbstractionSDK().EntryPoint.GetNonce(
		context.Background(),
		walletAddress,
		key,
	)
	if err != nil {
		response.Error(c, fiber.StatusInternalServerError, "Can Not Get Nonce")
	}

	nonceValue := nonce.Uint64() & ((1 << 64) - 1)

	resp := EncodeCreateDIDParamResponse{
		Data:   decodedData,
		Target: didrootsdk.GetDIDRootSDK().RootFactory.Address.Hex(),
		Nonce:  nonceValue,
	}

	return response.Success(c, resp)
}
