package helperhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/pkg/response"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	helperdto "github.com/meQlause/go-be-did/internal/domain/dtos/helper"
	backendutils "github.com/meQlause/go-be-did/utils"
)

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
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeCreateDIDDTO true "Encode DID creation parameters with wallet address, DID parameter, and key identifier"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse} "Successfully encoded DID creation parameters"
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON, missing required fields, or invalid address format"
// @Failure      500 {object} response.Response "Internal server error - encoding failed, nonce retrieval failed, or RPC node errors"
// @Router       /helper/encode-create-did-param [post]
func (hh *HelperHandler) EncodeCreateDIDParam(c *fiber.Ctx) error {
	var input helperdto.EncodeCreateDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)

	}

	encodedData, err := hh.uc.EncodeCreateDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeCreateDID,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeUpdateDIDParam godoc
// @Summary      Encode Update DID Parameters
// @Description  Encodes parameters for updating a DID's URI
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeUpdateDIDDTO true "Update DID parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-update-did-param [post]
func (hh *HelperHandler) EncodeUpdateDIDParam(c *fiber.Ctx) error {
	var input helperdto.EncodeUpdateDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeUpdateDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeUpdateDID,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeDeactiveDIDParam godoc
// @Summary      Encode Deactivate DID Parameters
// @Description  Encodes parameters for deactivating a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeDeactiveDIDDTO true "Deactivate DID parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-deactivate-did-param [post]
func (hh *HelperHandler) EncodeDeactiveDIDParam(c *fiber.Ctx) error {
	var input helperdto.EncodeDeactiveDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeDeactiveDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeDeactivateDID,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeReactiveDIDParam godoc
// @Summary      Encode Reactivate DID Parameters
// @Description  Encodes parameters for reactivating a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeReactiveDIDDTO true "Reactivate DID parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-reactivate-did-param [post]
func (hh *HelperHandler) EncodeReactiveDIDParam(c *fiber.Ctx) error {
	var input helperdto.EncodeReactiveDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeReactiveDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeReactivateDID,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeTransferDIDOwnerParam godoc
// @Summary      Encode Transfer DID Owner Parameters
// @Description  Encodes parameters for transferring DID ownership
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeTransferDIDOwnerDTO true "Transfer DID ownership parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-transfer-did-owner-param [post]
func (hh *HelperHandler) EncodeTransferDIDOwnerParam(c *fiber.Ctx) error {
	var input helperdto.EncodeTransferDIDOwnerDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeTransferDIDOwnerParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeTransferDID,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeStoreDataParam godoc
// @Summary      Encode Store Data Parameters
// @Description  Encodes parameters for storing data in a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeStoreDataDTO true "Store data parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-store-data-param [post]
func (hh *HelperHandler) EncodeStoreDataParam(c *fiber.Ctx) error {
	var input helperdto.EncodeStoreDataDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeStoreDataParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeStoreData,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeDeleteDataParam godoc
// @Summary      Encode Delete Data Parameters
// @Description  Encodes parameters for deleting data from a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeDeleteDataDTO true "Delete data parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-delete-data-param [post]
func (hh *HelperHandler) EncodeDeleteDataParam(c *fiber.Ctx) error {
	var input helperdto.EncodeDeleteDataDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeDeleteDataParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeDeleteData,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeAddKeyParam godoc
// @Summary      Encode Add Key Parameters
// @Description  Encodes parameters for adding a key to a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeAddKeyDTO true "Add key parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-add-key-param [post]
func (hh *HelperHandler) EncodeAddKeyParam(c *fiber.Ctx) error {
	var input helperdto.EncodeAddKeyDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeAddKeyParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeAddKey,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeRemoveKeyParam godoc
// @Summary      Encode Remove Key Parameters
// @Description  Encodes parameters for removing a key from a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeRemoveKeyDTO true "Remove key parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-remove-key-param [post]
func (hh *HelperHandler) EncodeRemoveKeyParam(c *fiber.Ctx) error {
	var input helperdto.EncodeRemoveKeyDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRemoveKeyParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeRemoveKey,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeAddClaimParam godoc
// @Summary      Encode Add Claim Parameters
// @Description  Encodes parameters for adding a claim to a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeAddClaimDTO true "Add claim parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-add-claim-param [post]
func (hh *HelperHandler) EncodeAddClaimParam(c *fiber.Ctx) error {
	var input helperdto.EncodeAddClaimDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeAddClaimParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeAddClaim,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeRemoveClaimParam godoc
// @Summary      Encode Remove Claim Parameters
// @Description  Encodes parameters for removing a claim from a DID
// @Tags Helper - DID Root
// @Accept       json
// @Produce      json
// @Param        request body helperdto.EncodeRemoveClaimDTO true "Remove claim parameters"
// @Success      200 {object} response.Response{data=helperhandler.HelperResponse}
// @Failure      400 {object} response.Response
// @Failure      500 {object} response.Response
// @Router       /helper/encode-remove-claim-param [post]
func (hh *HelperHandler) EncodeRemoveClaimParam(c *fiber.Ctx) error {
	var input helperdto.EncodeRemoveClaimDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRemoveClaimParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDRoot,
		TxType:  backendutils.TypeRemoveClaim,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}
