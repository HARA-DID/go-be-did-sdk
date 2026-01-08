package didaliashandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/config"
	didaliasevent "github.com/meQlause/go-be-did/internal/delivery/event/didalias"
	didaliasdto "github.com/meQlause/go-be-did/internal/domain/dtos/didalias"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"
)

// Resolve godoc
// @Summary      Resolve Node to DID
// @Description  Retrieves the DID associated with a specific node hash.
// @Description
// @Description  This endpoint returns:
// @Description  - The DID hash associated with the given node
// @Description  - Returns empty hash if node doesn't exist or has no DID set
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with the DID hash
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        node query string true "Node hash (32 bytes hex)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "DID resolution result"
// @Failure      400 {object} response.Response "Invalid request - malformed node hash or validation error"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/resolve [get]
func (h *DIDAliasHandler) Resolve(c *fiber.Ctx) error {
	var input didaliasdto.ResolveDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.Resolve(c.Context(), input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result.Hex(),
	}

	return response.Success(c, resp)
}

// ResolveFromString godoc
// @Summary      Resolve Domain Name to DID
// @Description  Retrieves the DID associated with a domain name string.
// @Description
// @Description  This endpoint:
// @Description  - Computes the namehash of the domain name
// @Description  - Returns the DID associated with that domain
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with the DID hash
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        name query string true "Domain name" example(example.tld)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "DID resolution result"
// @Failure      400 {object} response.Response "Invalid request - empty name or validation error"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/resolve-from-string [get]
func (h *DIDAliasHandler) ResolveFromString(c *fiber.Ctx) error {
	var input didaliasdto.ResolveFromStringDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.ResolveFromString(c.Context(), input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result.Hex(),
	}

	return response.Success(c, resp)
}

// GetAliasStatus godoc
// @Summary      Get Alias Status by Node
// @Description  Retrieves the status of an alias including expiry time, revocation status, and validity.
// @Description
// @Description  This endpoint returns:
// @Description  - Expiry timestamp (Unix timestamp)
// @Description  - Whether the alias is revoked
// @Description  - Whether the alias is currently valid (not expired and not revoked)
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with status details
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        node query string true "Node hash (32 bytes hex)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Alias status"
// @Failure      400 {object} response.Response "Invalid request - malformed node hash"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/status [get]
func (h *DIDAliasHandler) GetAliasStatus(c *fiber.Ctx) error {
	var input didaliasdto.GetAliasStatusDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.GetAliasStatus(c.Context(), input.Into())
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

// GetAliasStatusFromString godoc
// @Summary      Get Alias Status by Domain Name
// @Description  Retrieves the status of an alias using its domain name.
// @Description
// @Description  This endpoint:
// @Description  - Computes the namehash of the domain name
// @Description  - Returns expiry, revocation status, and validity
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with status details
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        name query string true "Domain name" example(example.tld)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Alias status"
// @Failure      400 {object} response.Response "Invalid request - empty name"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/status-from-string [get]
func (h *DIDAliasHandler) GetAliasStatusFromString(c *fiber.Ctx) error {
	var input didaliasdto.GetAliasStatusFromStringDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.GetAliasStatusFromString(c.Context(), input.Into())
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

// GetOwner godoc
// @Summary      Get Node Owner
// @Description  Retrieves the owner address of a node.
// @Description
// @Description  This endpoint returns:
// @Description  - The Ethereum address that owns the node/alias
// @Description  - Returns zero address if node doesn't exist
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with owner address
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        node query string true "Node hash (32 bytes hex)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Owner address"
// @Failure      400 {object} response.Response "Invalid request - malformed node hash"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/owner [get]
func (h *DIDAliasHandler) GetOwner(c *fiber.Ctx) error {
	var input didaliasdto.GetOwnerDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.GetOwner(c.Context(), input.Into())
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

// GetOwnerFromString godoc
// @Summary      Get Owner by Domain Name
// @Description  Retrieves the owner address using a domain name.
// @Description
// @Description  This endpoint:
// @Description  - Computes the namehash of the domain name
// @Description  - Returns the owner's Ethereum address
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with owner address
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        name query string true "Domain name" example(example.tld)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Owner address"
// @Failure      400 {object} response.Response "Invalid request - empty name"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/owner-from-string [get]
func (h *DIDAliasHandler) GetOwnerFromString(c *fiber.Ctx) error {
	var input didaliasdto.GetOwnerFromStringDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.GetOwnerFromString(c.Context(), input.Into())
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

// GetDID godoc
// @Summary      Get DID by Node
// @Description  Retrieves the DID hash associated with a node.
// @Description
// @Description  This endpoint returns:
// @Description  - The DID hash linked to the node
// @Description  - Returns empty hash if no DID is set
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with DID hash
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        node query string true "Node hash (32 bytes hex)" example(0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "DID hash"
// @Failure      400 {object} response.Response "Invalid request - malformed node hash"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/did [get]
func (h *DIDAliasHandler) GetDID(c *fiber.Ctx) error {
	var input didaliasdto.GetDIDDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.GetDID(c.Context(), input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result.Hex(),
	}

	return response.Success(c, resp)
}

// GetDIDFromString godoc
// @Summary      Get DID by Domain Name
// @Description  Retrieves the DID hash using a domain name.
// @Description
// @Description  This endpoint:
// @Description  - Computes the namehash of the domain name
// @Description  - Returns the associated DID hash
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with DID hash
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        name query string true "Domain name" example(example.tld)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "DID hash"
// @Failure      400 {object} response.Response "Invalid request - empty name"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/did-from-string [get]
func (h *DIDAliasHandler) GetDIDFromString(c *fiber.Ctx) error {
	var input didaliasdto.GetDIDFromStringDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.GetDIDFromString(c.Context(), input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result.Hex(),
	}

	return response.Success(c, resp)
}

// Namehash godoc
// @Summary      Compute Namehash
// @Description  Computes the namehash of a domain name according to ENS specification.
// @Description
// @Description  This endpoint:
// @Description  - Takes a domain name string
// @Description  - Returns its namehash (used as node identifier)
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with namehash
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        name query string true "Domain name" example(example.tld)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Namehash result"
// @Failure      400 {object} response.Response "Invalid request - empty name"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/namehash [get]
func (h *DIDAliasHandler) Namehash(c *fiber.Ctx) error {
	var input didaliasdto.NamehashDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.Namehash(c.Context(), input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result.Hex(),
	}

	return response.Success(c, resp)
}

// GetRegistrationPeriod godoc
// @Summary      Get Registration Period Duration
// @Description  Retrieves the duration (in seconds) for a given registration period type.
// @Description
// @Description  This endpoint returns:
// @Description  - Duration in seconds for the specified period
// @Description  - Period types: 0 (ONE_YEAR), 1 (TWO_YEARS), 2 (THREE_YEARS)
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true if request is processed correctly
// @Description  - `data` (object): BlockchainResponse with duration in seconds
// @Description  - `meta` (object): Contains timestamp and API version
// @Tags         did-alias
// @Accept       json
// @Produce      json
// @Param        period query string true "Registration period (0, 1, or 2)" example(0)
// @Success      200 {object} response.Response{data=response.BlockchainResponse} "Registration period duration"
// @Failure      400 {object} response.Response "Invalid request - period must be 0, 1, or 2"
// @Failure      500 {object} response.Response "Internal server error - blockchain call failed"
// @Router       /did-alias/registration-period [get]
func (h *DIDAliasHandler) GetRegistrationPeriod(c *fiber.Ctx) error {
	var input didaliasdto.GetRegistrationPeriodDTO
	if err := c.QueryParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.GetRegistrationPeriod(c.Context(), input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	resp := response.BlockchainResponse{
		Success:  true,
		Errors:   "No Error Message",
		Returned: result.String(),
	}

	return response.Success(c, resp)
}

// RegisterTLD godoc
// @Summary Register a new Top-Level Domain (TLD)
// @Description Register a new top-level domain in the alias system. Only admins can register TLDs.
// @Tags DID Alias
// @Accept json
// @Produce json
// @Param request body didaliasdto.RegisterTLDDTO true "Register TLD Request"
// @Success 200 {object} response.Response{data=map[string]response.BlockchainResponse} "Successfully registered TLD"
// @Failure 400 {object} response.Response "Invalid request body or validation error"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /did-alias/register-tld [post]
func (h *DIDAliasHandler) RegisterTLD(c *fiber.Ctx) error {
	var input didaliasdto.RegisterTLDDTO
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	result, err := h.uc.RegisterTLD(c.Context(), input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	txSuccess, txErrors := config.Blockchain().CheckTxs(c.Context(), result.TxHash)

	resp := make(map[string]response.BlockchainResponse)

	for txHash, ok := range txSuccess {
		eventData, err := didaliasevent.DecodeRegisterTLDEvents(c.Context(), txHash)
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
