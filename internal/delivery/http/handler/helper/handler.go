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
// @Param        request body dto.EncodeCreateDIDDTO true "Encode DID creation parameters with wallet address, DID parameter, and key identifier"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse} "Successfully encoded DID creation parameters"
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON, missing required fields, or invalid address format"
// @Failure      500 {object} response.Response "Internal server error - encoding failed, nonce retrieval failed, or RPC node errors"
// @Router       /helper/encode-create-did-param [post]
func (hh *HelperHandler) EncodeCreateDIDParam(c *fiber.Ctx) error {
	var input dto.EncodeCreateDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)

	}

	encodedData, err := hh.uc.EncodeCreateDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeUpdateDIDParam godoc
// @Summary      Encode Update DID Parameters
// @Description  Encodes parameters for updating a DID's URI
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeUpdateDIDDTO true "Update DID parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-update-did-param [post]
func (hh *HelperHandler) EncodeUpdateDIDParam(c *fiber.Ctx) error {
	var input dto.EncodeUpdateDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeUpdateDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeDeactiveDIDParam godoc
// @Summary      Encode Deactivate DID Parameters
// @Description  Encodes parameters for deactivating a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeDeactiveDIDDTO true "Deactivate DID parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-deactivate-did-param [post]
func (hh *HelperHandler) EncodeDeactiveDIDParam(c *fiber.Ctx) error {
	var input dto.EncodeDeactiveDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeDeactiveDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeReactiveDIDParam godoc
// @Summary      Encode Reactivate DID Parameters
// @Description  Encodes parameters for reactivating a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeReactiveDIDDTO true "Reactivate DID parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-reactivate-did-param [post]
func (hh *HelperHandler) EncodeReactiveDIDParam(c *fiber.Ctx) error {
	var input dto.EncodeReactiveDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeReactiveDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeTransferDIDOwnerParam godoc
// @Summary      Encode Transfer DID Owner Parameters
// @Description  Encodes parameters for transferring DID ownership
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeTransferDIDOwnerDTO true "Transfer DID ownership parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-transfer-did-owner-param [post]
func (hh *HelperHandler) EncodeTransferDIDOwnerParam(c *fiber.Ctx) error {
	var input dto.EncodeTransferDIDOwnerDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeTransferDIDOwnerParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeStoreDataParam godoc
// @Summary      Encode Store Data Parameters
// @Description  Encodes parameters for storing data in a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeStoreDataDTO true "Store data parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-store-data-param [post]
func (hh *HelperHandler) EncodeStoreDataParam(c *fiber.Ctx) error {
	var input dto.EncodeStoreDataDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeStoreDataParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeDeleteDataParam godoc
// @Summary      Encode Delete Data Parameters
// @Description  Encodes parameters for deleting data from a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeDeleteDataDTO true "Delete data parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-delete-data-param [post]
func (hh *HelperHandler) EncodeDeleteDataParam(c *fiber.Ctx) error {
	var input dto.EncodeDeleteDataDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeDeleteDataParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeAddKeyParam godoc
// @Summary      Encode Add Key Parameters
// @Description  Encodes parameters for adding a key to a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeAddKeyDTO true "Add key parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-add-key-param [post]
func (hh *HelperHandler) EncodeAddKeyParam(c *fiber.Ctx) error {
	var input dto.EncodeAddKeyDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeAddKeyParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeRemoveKeyParam godoc
// @Summary      Encode Remove Key Parameters
// @Description  Encodes parameters for removing a key from a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeRemoveKeyDTO true "Remove key parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-remove-key-param [post]
func (hh *HelperHandler) EncodeRemoveKeyParam(c *fiber.Ctx) error {
	var input dto.EncodeRemoveKeyDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRemoveKeyParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeAddClaimParam godoc
// @Summary      Encode Add Claim Parameters
// @Description  Encodes parameters for adding a claim to a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeAddClaimDTO true "Add claim parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-add-claim-param [post]
func (hh *HelperHandler) EncodeAddClaimParam(c *fiber.Ctx) error {
	var input dto.EncodeAddClaimDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeAddClaimParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeRemoveClaimParam godoc
// @Summary      Encode Remove Claim Parameters
// @Description  Encodes parameters for removing a claim from a DID
// @Tags         helper
// @Accept       json
// @Produce      json
// @Param        request body dto.EncodeRemoveClaimDTO true "Remove claim parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-remove-claim-param [post]
func (hh *HelperHandler) EncodeRemoveClaimParam(c *fiber.Ctx) error {
	var input dto.EncodeRemoveClaimDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRemoveClaimParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	resp, err := hh.buildHelperResponse(input.Address, encodedData)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

func (hh *HelperHandler) parseAndValidate(c *fiber.Ctx, input any) any {
	if err := c.BodyParser(&input); err != nil {
		return "Invalid request body"
	}

	if err := validator.Validate.Struct(input); err != nil {
		validationErrors := validator.FormatError(err)
		return validationErrors
	}

	return nil
}

func (hh *HelperHandler) buildHelperResponse(address string, encodedData string) (HelperResponse, error) {
	walletAddress := utils.HexToAddress(address)

	key := new(big.Int).SetInt64(0)
	nonce, err := accountabstractionsdk.GetAccountAbstractionSDK().EntryPoint.GetNonce(
		context.Background(),
		walletAddress,
		key,
	)
	if err != nil {
		return HelperResponse{}, err
	}

	nonceValue := nonce.Uint64() & ((1 << 64) - 1)

	return HelperResponse{
		Data:   encodedData,
		Target: didrootsdk.GetDIDRootSDK().RootFactory.Address.Hex(),
		Nonce:  nonceValue,
	}, nil
}
