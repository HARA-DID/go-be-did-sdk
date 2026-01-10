package didvchandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"

	didvcdto "github.com/meQlause/go-be-did/internal/domain/dtos/didvc"
)

// GetMetadata godoc
// @Summary      Get Credential Metadata
// @Description  Retrieves metadata for a specific credential token
// @Tags         DID VC - NFT Base
// @Accept       json
// @Produce      json
// @Param        token_id query string true "Token ID (hex-encoded big integer with 0x prefix)" example(0x1)
// @Param        options query string true "Options flag (0 for identity, 1 for certificate)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Metadata retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - metadata retrieval failed"
// @Router       /did-vc/get-metadata [get]
func (dvh *DIDVCHandler) GetMetadata(c *fiber.Ctx) error {
	var input didvcdto.GetMetadataDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getMetadataInput := input.Into()

	result, err := dvh.uc.GetMetadata(c.Context(), getMetadataInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// IsCredentialValid godoc
// @Summary      Check Credential Validity
// @Description  Verifies if a credential is valid and not expired
// @Tags         DID VC - NFT Base
// @Accept       json
// @Produce      json
// @Param        token_id query string true "Token ID (hex-encoded big integer with 0x prefix)" example(0x1)
// @Param        options query string true "Options flag (0 for identity, 1 for certificate)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Validity check completed"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - validity check failed"
// @Router       /did-vc/is-credential-valid [get]
func (dvh *DIDVCHandler) IsCredentialValid(c *fiber.Ctx) error {
	var input didvcdto.IsCredentialValidDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	isCredentialValidInput := input.Into()

	result, err := dvh.uc.IsCredentialValid(c.Context(), isCredentialValidInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetCredentialsWithMetadata godoc
// @Summary      Get Multiple Credentials with Metadata
// @Description  Retrieves multiple credentials and their metadata in a single call
// @Tags         DID VC - NFT Base
// @Accept       json
// @Produce      json
// @Param        token_ids query string true "Comma-separated token IDs" example(1,2,3)
// @Param        options query string true "Options flag (0 for identity, 1 for certificate)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Credentials retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - credentials retrieval failed"
// @Router       /did-vc/get-credentials-with-metadata [get]
func (dvh *DIDVCHandler) GetCredentialsWithMetadata(c *fiber.Ctx) error {
	var input didvcdto.GetCredentialsWithMetadataDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getCredentialsWithMetadataInput := input.Into()

	result, err := dvh.uc.GetCredentialsWithMetadata(c.Context(), getCredentialsWithMetadataInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetUnclaimedTokenId godoc
// @Summary      Get Unclaimed Token DID
// @Description  Retrieves the DID hash of an unclaimed credential token
// @Tags         DID VC - NFT Base
// @Accept       json
// @Produce      json
// @Param        token_id query string true "Token ID (hex-encoded big integer with 0x prefix)" example(0x1)
// @Param        options query string true "Options flag (0 for identity, 1 for certificate)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Unclaimed token DID retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - retrieval failed"
// @Router       /did-vc/get-unclaimed-token-id [get]
func (dvh *DIDVCHandler) GetUnclaimedTokenId(c *fiber.Ctx) error {
	var input didvcdto.GetUnclaimedTokenIdDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getUnclaimedTokenIdInput := input.Into()

	result, err := dvh.uc.GetUnclaimedTokenId(c.Context(), getUnclaimedTokenIdInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// TotalTokensToBeClaimedByDid godoc
// @Summary      Get Total Unclaimed Tokens Count
// @Description  Retrieves the total number of unclaimed credentials for a specific DID
// @Tags         DID VC - NFT Base
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        options query string true "Options flag (0 for identity, 1 for certificate)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Count retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - count retrieval failed"
// @Router       /did-vc/total-tokens-to-be-claimed-by-did [get]
func (dvh *DIDVCHandler) TotalTokensToBeClaimedByDid(c *fiber.Ctx) error {
	var input didvcdto.TotalTokensToBeClaimedByDidDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	totalTokensToBeClaimedByDidInput := input.Into()

	result, err := dvh.uc.TotalTokensToBeClaimedByDid(c.Context(), totalTokensToBeClaimedByDidInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetToBeClaimedTokensByDid godoc
// @Summary      Get Unclaimed Tokens by DID
// @Description  Retrieves a paginated list of unclaimed credential token IDs for a specific DID
// @Tags         DID VC - NFT Base
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        offset query string true "Pagination offset" example(0)
// @Param        limit query string true "Pagination limit" example(10)
// @Param        options query string true "Options flag (0 for identity, 1 for certificate)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Token IDs retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - retrieval failed"
// @Router       /did-vc/get-to-be-claimed-tokens-by-did [get]
func (dvh *DIDVCHandler) GetToBeClaimedTokensByDid(c *fiber.Ctx) error {
	var input didvcdto.GetToBeClaimedTokensByDidDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getToBeClaimedTokensByDidInput := input.Into()

	result, err := dvh.uc.GetToBeClaimedTokensByDid(c.Context(), getToBeClaimedTokensByDidInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// IsApprovedForAll godoc
// @Summary      Check Operator Approval Status
// @Description  Verifies if an operator is approved to manage all credentials for an owner
// @Tags         DID VC - NFT Base
// @Accept       json
// @Produce      json
// @Param        owner query string true "Owner's Ethereum address (with 0x prefix)" example(0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb)
// @Param        operator query string true "Operator's Ethereum address (with 0x prefix)" example(0x1234567890abcdef1234567890abcdef12345678)
// @Param        options query string true "Options flag (0 for identity, 1 for certificate)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Approval status retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid address format"
// @Failure      500 {object} response.Response "Internal server error - approval check failed"
// @Router       /did-vc/is-approved-for-all [get]
func (dvh *DIDVCHandler) IsApprovedForAll(c *fiber.Ctx) error {
	var input didvcdto.IsApprovedForAllDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	isApprovedForAllInput := input.Into()

	result, err := dvh.uc.IsApprovedForAll(c.Context(), isApprovedForAllInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}
