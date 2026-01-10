package didvchandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"

	didvcdto "github.com/meQlause/go-be-did/internal/domain/dtos/didvc"
)

// GetIdentityTokenCount godoc
// @Summary      Get Identity Token Count
// @Description  Retrieves the total number of identity credentials for a specific DID
// @Tags         DID VC - VC Storage
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Identity token count retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - count retrieval failed"
// @Router       /did-vc/get-identity-token-count [get]
func (dvh *DIDVCHandler) GetIdentityTokenCount(c *fiber.Ctx) error {
	var input didvcdto.GetIdentityTokenCountDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getIdentityTokenCountInput := input.Into()

	result, err := dvh.uc.GetIdentityTokenCount(c.Context(), getIdentityTokenCountInput)
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

// GetCertificateTokenCount godoc
// @Summary      Get Certificate Token Count
// @Description  Retrieves the total number of certificate credentials for a specific DID
// @Tags         DID VC - VC Storage
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Certificate token count retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - count retrieval failed"
// @Router       /did-vc/get-certificate-token-count [get]
func (dvh *DIDVCHandler) GetCertificateTokenCount(c *fiber.Ctx) error {
	var input didvcdto.GetCertificateTokenCountDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getCertificateTokenCountInput := input.Into()

	result, err := dvh.uc.GetCertificateTokenCount(c.Context(), getCertificateTokenCountInput)
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

// GetIdentityTokenIds godoc
// @Summary      Get Identity Token IDs with Pagination
// @Description  Retrieves a paginated list of identity credential token IDs for a specific DID
// @Tags         DID VC - VC Storage
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        offset query string true "Pagination offset" example(0)
// @Param        limit query string true "Pagination limit" example(10)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Identity token IDs retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - retrieval failed"
// @Router       /did-vc/get-identity-token-ids [get]
func (dvh *DIDVCHandler) GetIdentityTokenIds(c *fiber.Ctx) error {
	var input didvcdto.GetIdentityTokenIdsDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getIdentityTokenIdsInput := input.Into()

	result, err := dvh.uc.GetIdentityTokenIds(c.Context(), getIdentityTokenIdsInput)
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

// GetCertificateTokenIds godoc
// @Summary      Get Certificate Token IDs with Pagination
// @Description  Retrieves a paginated list of certificate credential token IDs for a specific DID
// @Tags         DID VC - VC Storage
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        offset query string true "Pagination offset" example(0)
// @Param        limit query string true "Pagination limit" example(10)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Certificate token IDs retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid format"
// @Failure      500 {object} response.Response "Internal server error - retrieval failed"
// @Router       /did-vc/get-certificate-token-ids [get]
func (dvh *DIDVCHandler) GetCertificateTokenIds(c *fiber.Ctx) error {
	var input didvcdto.GetCertificateTokenIdsDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getCertificateTokenIdsInput := input.Into()

	result, err := dvh.uc.GetCertificateTokenIds(c.Context(), getCertificateTokenIdsInput)
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

// GetAllIdentityTokenIds godoc
// @Summary      Get All Identity Token IDs
// @Description  Retrieves all identity credential token IDs for a specific DID without pagination
// @Tags         DID VC - VC Storage
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "All identity token IDs retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - retrieval failed"
// @Router       /did-vc/get-all-identity-token-ids [get]
func (dvh *DIDVCHandler) GetAllIdentityTokenIds(c *fiber.Ctx) error {
	var input didvcdto.GetAllIdentityTokenIdsDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getAllIdentityTokenIdsInput := input.Into()

	result, err := dvh.uc.GetAllIdentityTokenIds(c.Context(), getAllIdentityTokenIdsInput)
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

// GetAllCertificateTokenIds godoc
// @Summary      Get All Certificate Token IDs
// @Description  Retrieves all certificate credential token IDs for a specific DID without pagination
// @Tags         DID VC - VC Storage
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "All certificate token IDs retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - retrieval failed"
// @Router       /did-vc/get-all-certificate-token-ids [get]
func (dvh *DIDVCHandler) GetAllCertificateTokenIds(c *fiber.Ctx) error {
	var input didvcdto.GetAllCertificateTokenIdsDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getAllCertificateTokenIdsInput := input.Into()

	result, err := dvh.uc.GetAllCertificateTokenIds(c.Context(), getAllCertificateTokenIdsInput)
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

// GetDIDRootStorage godoc
// @Summary      Get DID Root Storage Address
// @Description  Retrieves the contract address of the DID root storage
// @Tags         DID VC - VC Storage
// @Accept       json
// @Produce      json
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "DID root storage address retrieved successfully"
// @Failure      500 {object} response.Response "Internal server error - address retrieval failed"
// @Router       /did-vc/get-did-root-storage [get]
func (dvh *DIDVCHandler) GetDIDRootStorage(c *fiber.Ctx) error {
	var input didvcdto.GetDIDRootStorageDTO

	getDIDRootStorageInput := input.Into()

	result, err := dvh.uc.GetDIDRootStorage(c.Context(), getDIDRootStorageInput)
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
