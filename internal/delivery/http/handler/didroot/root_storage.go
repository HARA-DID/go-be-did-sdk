package didroothandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"

	didrootdto "github.com/meQlause/go-be-did/internal/domain/dto/didroot"
	backendutils "github.com/meQlause/go-be-did/utils"
)

// GetData godoc
// @Summary      Get Data by Hash
// @Description  Retrieves data associated with a specific hash from the DID Root contract
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "Data hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        key query string true "key string" example(email)
// @Success      200 {object} response.Response{data=backendutils.Response} "Data retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed hash parameter"
// @Failure      500 {object} response.Response "Internal server error - contract call failed or network issues"
// @Router       /did-root/get-data [get]
func (drh *DIDRootHandler) GetData(c *fiber.Ctx) error {
	var input didrootdto.GetDataDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getDataInput := input.Into()

	result, err := drh.uc.GetData(c.Context(), getDataInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// ResolveDID godoc
// @Summary      Resolve DID
// @Description  Resolves a Decentralized Identifier (DID) to retrieve its associated document and metadata
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=backendutils.Response} "DID resolved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - DID resolution failed or not found"
// @Router       /did-root/resolve-did [get]
func (drh *DIDRootHandler) ResolveDID(c *fiber.Ctx) error {
	var input didrootdto.ResolveDIDDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	resolveDIDInput := input.Into()

	result, err := drh.uc.ResolveDID(c.Context(), resolveDIDInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// VerifyDIDOwnership godoc
// @Summary      Verify DID Ownership
// @Description  Verifies if a specific address is the owner of a given DID
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        owner query string true "Ethereum address to verify ownership (with 0x prefix)" example(0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb)
// @Success      200 {object} response.Response{data=backendutils.Response} "Ownership verification completed - check returned data for result"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid address format"
// @Failure      500 {object} response.Response "Internal server error - verification failed"
// @Router       /did-root/verify-did-ownership [get]
func (drh *DIDRootHandler) VerifyDIDOwnership(c *fiber.Ctx) error {
	var input didrootdto.VerifyDIDOwnershipDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	verifyDIDOwnershipInput := input.Into()

	result, err := drh.uc.VerifyDIDOwnership(c.Context(), verifyDIDOwnershipInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetKey godoc
// @Summary      Get Key by DID and Key Hash
// @Description  Retrieves a specific cryptographic key associated with a DID using the key's hashed data
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        key_data_hashed query string true "Hashed key data (32-byte hex string with 0x prefix)" example(0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890)
// @Success      200 {object} response.Response{data=backendutils.Response} "Key retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid hash format"
// @Failure      500 {object} response.Response "Internal server error - key retrieval failed or not found"
// @Router       /did-root/get-key [get]
func (drh *DIDRootHandler) GetKey(c *fiber.Ctx) error {
	var input didrootdto.GetKeyDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getKeyInput := input.Into()

	result, err := drh.uc.GetKey(c.Context(), getKeyInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetKeysByDID godoc
// @Summary      Get All Keys for a DID
// @Description  Retrieves all cryptographic keys associated with a specific DID
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=backendutils.Response} "Keys retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - keys retrieval failed"
// @Router       /did-root/get-keys-by-did [get]
func (drh *DIDRootHandler) GetKeysByDID(c *fiber.Ctx) error {
	var input didrootdto.GetKeysByDIDDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getKeysByDIDInput := input.Into()

	result, err := drh.uc.GetKeysByDID(c.Context(), getKeysByDIDInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetClaim godoc
// @Summary      Get Claim by DID and Claim ID
// @Description  Retrieves a specific verifiable claim associated with a DID
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        claim_id query string true "Claim identifier (32-byte hex string with 0x prefix)" example(0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890)
// @Success      200 {object} response.Response{data=backendutils.Response} "Claim retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid hash format"
// @Failure      500 {object} response.Response "Internal server error - claim retrieval failed or not found"
// @Router       /did-root/get-claim [get]
func (drh *DIDRootHandler) GetClaim(c *fiber.Ctx) error {
	var input didrootdto.GetClaimDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getClaimInput := input.Into()

	result, err := drh.uc.GetClaim(c.Context(), getClaimInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetClaimsByDID godoc
// @Summary      Get All Claims for a DID
// @Description  Retrieves all verifiable claims associated with a specific DID
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=backendutils.Response} "Claims retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - claims retrieval failed"
// @Router       /did-root/get-claims-by-did [get]
func (drh *DIDRootHandler) GetClaimsByDID(c *fiber.Ctx) error {
	var input didrootdto.GetClaimsByDIDDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getClaimsByDIDInput := input.Into()

	result, err := drh.uc.GetClaimsByDID(c.Context(), getClaimsByDIDInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// VerifyClaim godoc
// @Summary      Verify Claim
// @Description  Verifies if a specific address has authority over a claim for a given DID
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        claim_id query string true "Claim identifier (32-byte hex string with 0x prefix)" example(0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890)
// @Param        to_verify query string true "Ethereum address to verify claim authority (with 0x prefix)" example(0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb)
// @Success      200 {object} response.Response{data=backendutils.Response} "Claim verification completed - check returned data for result"
// @Failure      400 {object} response.Response "Invalid request - missing parameters or invalid address/hash format"
// @Failure      500 {object} response.Response "Internal server error - verification failed"
// @Router       /did-root/verify-claim [get]
func (drh *DIDRootHandler) VerifyClaim(c *fiber.Ctx) error {
	var input didrootdto.VerifyClaimDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	verifyClaimInput := input.Into()

	result, err := drh.uc.VerifyClaim(c.Context(), verifyClaimInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetDIDKeyDataCount godoc
// @Summary      Get DID Key Data Count
// @Description  Returns the total number of keys associated with a specific DID
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=backendutils.Response} "Key count retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - count retrieval failed"
// @Router       /did-root/get-did-key-data-count [get]
func (drh *DIDRootHandler) GetDIDKeyDataCount(c *fiber.Ctx) error {
	var input didrootdto.GetDIDKeyDataCountDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getDIDKeyDataCountInput := input.Into()

	result, err := drh.uc.GetDIDKeyDataCount(c.Context(), getDIDKeyDataCountInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// getDIDKeyDataByIndex godoc
// @Summary      Get DID Key by Index
// @Description  Retrieves a specific key associated with a DID using its index position
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Param        index query string true "Key index (unsigned integer as string)" example(0)
// @Success      200 {object} response.Response{data=backendutils.Response} "Key retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing parameters, invalid hash format, or invalid index"
// @Failure      500 {object} response.Response "Internal server error - key retrieval failed or index out of bounds"
// @Router       /did-root/get-did-key-data-by-index [get]
func (drh *DIDRootHandler) GetDIDKeyDataByIndex(c *fiber.Ctx) error {
	var input didrootdto.GetDIDKeyDataByIndexDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getDIDKeyDataByIndexInput := input.Into()

	result, err := drh.uc.GetDIDKeyDataByIndex(c.Context(), getDIDKeyDataByIndexInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// GetOriginalKey godoc
// @Summary      Get Original Key by Key Code
// @Description  Retrieves the original key data using a key code identifier
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        key_code query string true "Key code identifier (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=backendutils.Response} "Original key retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed key_code parameter"
// @Failure      500 {object} response.Response "Internal server error - key retrieval failed or not found"
// @Router       /did-root/get-original-key [get]
func (drh *DIDRootHandler) GetOriginalKey(c *fiber.Ctx) error {
	var input didrootdto.GetOriginalKeyDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	getOriginalKeyInput := input.Into()

	result, err := drh.uc.GetOriginalKey(c.Context(), getOriginalKeyInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// DIDIndexMap godoc
// @Summary      Get DID Hash by Index
// @Description  Maps a DID index to its corresponding DID hash
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_index query string true "DID index (hex-encoded big integer with 0x prefix)" example(0x1)
// @Success      200 {object} response.Response{data=backendutils.Response} "DID hash retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_index parameter"
// @Failure      500 {object} response.Response "Internal server error - mapping retrieval failed or index not found"
// @Router       /did-root/did-index-map [get]
func (drh *DIDRootHandler) DIDIndexMap(c *fiber.Ctx) error {
	var input didrootdto.DIDIndexMapDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	didIndexMapInput := input.Into()

	result, err := drh.uc.DIDIndexMap(c.Context(), didIndexMapInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}

// DIDIndexMapReverse godoc
// @Summary      Get DID Index by Hash
// @Description  Reverse maps a DID hash to its corresponding DID index
// @Tags         did-root
// @Accept       json
// @Produce      json
// @Param        did_hash query string true "DID hash (32-byte hex string with 0x prefix)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=backendutils.Response} "DID index retrieved successfully"
// @Failure      400 {object} response.Response "Invalid request - missing or malformed did_hash parameter"
// @Failure      500 {object} response.Response "Internal server error - reverse mapping retrieval failed or hash not found"
// @Router       /did-root/did-index-map-reverse [get]
func (drh *DIDRootHandler) DIDIndexMapReverse(c *fiber.Ctx) error {
	var input didrootdto.DIDIndexMapReverseDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	didIndexMapReverseInput := input.Into()

	result, err := drh.uc.DIDIndexMapReverse(c.Context(), didIndexMapReverseInput)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := backendutils.Response{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result,
	}

	return response.Success(c, resp)
}
