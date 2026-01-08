package helperhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/pkg/response"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	helperdto "github.com/meQlause/go-be-did/internal/domain/dtos/helper"
	backendutils "github.com/meQlause/go-be-did/utils"
)

// EncodeSetDIDRootStorageParam godoc
// @Summary Encode Set DID Root Storage parameters
// @Description Encodes parameters for setting the DID root storage address
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeSetDIDRootStorageDTO true "Set DID Root Storage parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-set-did-root-storage-param [post]
func (hh *HelperHandler) EncodeSetDIDRootStorageParam(c *fiber.Ctx) error {
	var input helperdto.EncodeSetDIDRootStorageDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeSetDIDRootStorageParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeSetDIDRootStorage,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeRegisterDomainParam godoc
// @Summary Encode Register Domain parameters
// @Description Encodes parameters for registering a new domain under a TLD
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeRegisterDomainDTO true "Register Domain parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-register-domain-param [post]
func (hh *HelperHandler) EncodeRegisterDomainParam(c *fiber.Ctx) error {
	var input helperdto.EncodeRegisterDomainDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRegisterDomainParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeRegisterDomain,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeRegisterSubdomainParam godoc
// @Summary Encode Register Subdomain parameters
// @Description Encodes parameters for registering a new subdomain under a parent domain
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeRegisterSubdomainDTO true "Register Subdomain parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-register-subdomain-param [post]
func (hh *HelperHandler) EncodeRegisterSubdomainParam(c *fiber.Ctx) error {
	var input helperdto.EncodeRegisterSubdomainDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRegisterSubdomainParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeRegisterSubdomain,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeSetDIDParam godoc
// @Summary Encode Set DID parameters
// @Description Encodes parameters for setting a DID (Decentralized Identifier) to a node
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeSetDIDDTO true "Set DID parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-set-did-param [post]
func (hh *HelperHandler) EncodeSetDIDParam(c *fiber.Ctx) error {
	var input helperdto.EncodeSetDIDDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeSetDIDParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeSetDID,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeExtendRegistrationParam godoc
// @Summary Encode Extend Registration parameters
// @Description Encodes parameters for extending the registration period of an alias
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeExtendRegistrationDTO true "Extend Registration parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-extend-registration-param [post]
func (hh *HelperHandler) EncodeExtendRegistrationParam(c *fiber.Ctx) error {
	var input helperdto.EncodeExtendRegistrationDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeExtendRegistrationParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeExtendRegistration,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeRevokeAliasParam godoc
// @Summary Encode Revoke Alias parameters
// @Description Encodes parameters for revoking an alias
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeRevokeAliasDTO true "Revoke Alias parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-revoke-alias-param [post]
func (hh *HelperHandler) EncodeRevokeAliasParam(c *fiber.Ctx) error {
	var input helperdto.EncodeRevokeAliasDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRevokeAliasParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeRevokeAlias,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeUnrevokeAliasParam godoc
// @Summary Encode Unrevoke Alias parameters
// @Description Encodes parameters for unrevoking a previously revoked alias
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeUnrevokeAliasDTO true "Unrevoke Alias parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-unrevoke-alias-param [post]
func (hh *HelperHandler) EncodeUnrevokeAliasParam(c *fiber.Ctx) error {
	var input helperdto.EncodeUnrevokeAliasDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeUnrevokeAliasParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeUnrevokeAlias,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}

// EncodeTransferAliasOwnershipParam godoc
// @Summary Encode Transfer Alias Ownership parameters
// @Description Encodes parameters for transferring ownership of an alias to a new owner
// @Tags Helper - DID Alias
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeTransferAliasOwnershipDTO true "Transfer Alias Ownership parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-transfer-alias-ownership-param [post]
func (hh *HelperHandler) EncodeTransferAliasOwnershipParam(c *fiber.Ctx) error {
	var input helperdto.EncodeTransferAliasOwnershipDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeTransferAliasOwnershipParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	details := response.Details{
		Service: backendutils.ServiceDIDAlias,
		TxType:  backendutils.TypeTransferAliasOwnership,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, resp)
}
