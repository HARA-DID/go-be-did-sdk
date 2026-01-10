package helperhandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/pkg/response"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	helperdto "github.com/meQlause/go-be-did/internal/domain/dtos/helper"
	backendutils "github.com/meQlause/go-be-did/utils"
)

// EncodeIssueCredentialParam godoc
// @Summary Encode Issue Credential parameters
// @Description Encodes parameters for issuing a new verifiable credential
// @Tags Helper - DID VC
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeIssueCredentialDTO true "Issue Credential parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-issue-credential-param [post]
func (hh *HelperHandler) EncodeIssueCredentialParam(c *fiber.Ctx) error {
	var input helperdto.EncodeIssueCredentialDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeIssueCredentialParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	fmt.Println(encodedData)

	details := response.Details{
		Service: backendutils.ServiceDIDVC,
		TxType:  backendutils.TypeIssueCredential,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeBurnCredentialParam godoc
// @Summary Encode Burn Credential parameters
// @Description Encodes parameters for burning/destroying a verifiable credential
// @Tags Helper - DID VC
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeBurnCredentialDTO true "Burn Credential parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-burn-credential-param [post]
func (hh *HelperHandler) EncodeBurnCredentialParam(c *fiber.Ctx) error {
	var input helperdto.EncodeBurnCredentialDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeBurnCredentialParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDVC,
		TxType:  backendutils.TypeBurnCredential,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeUpdateMetadataParam godoc
// @Summary Encode Update Metadata parameters
// @Description Encodes parameters for updating credential metadata
// @Tags Helper - DID VC
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeUpdateMetadataDTO true "Update Metadata parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-update-metadata-param [post]
func (hh *HelperHandler) EncodeUpdateMetadataParam(c *fiber.Ctx) error {
	var input helperdto.EncodeUpdateMetadataDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeUpdateMetadataParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDVC,
		TxType:  backendutils.TypeUpdateMetadata,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeRevokeCredentialParam godoc
// @Summary Encode Revoke Credential parameters
// @Description Encodes parameters for revoking a verifiable credential
// @Tags Helper - DID VC
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeRevokeCredentialDTO true "Revoke Credential parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-revoke-credential-param [post]
func (hh *HelperHandler) EncodeRevokeCredentialParam(c *fiber.Ctx) error {
	var input helperdto.EncodeRevokeCredentialDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeRevokeCredentialParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDVC,
		TxType:  backendutils.TypeRevokeCredential,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}

// EncodeClaimCredentialParam godoc
// @Summary Encode Claim Credential parameters
// @Description Encodes parameters for claiming an unclaimed verifiable credential
// @Tags Helper - DID VC
// @Accept json
// @Produce json
// @Param request body helperdto.EncodeClaimCredentialDTO true "Claim Credential parameters"
// @Success 200 {object} response.Response{data=HelperResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /helper/encode-claim-credential-param [post]
func (hh *HelperHandler) EncodeClaimCredentialParam(c *fiber.Ctx) error {
	var input helperdto.EncodeClaimCredentialDTO
	if err := hh.parseAndValidate(c, &input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err)
	}

	encodedData, err := hh.uc.EncodeClaimCredentialParam(input.Into())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	details := response.Details{
		Service: backendutils.ServiceDIDVC,
		TxType:  backendutils.TypeClaimCredential,
	}

	resp, err := hh.buildHelperResponse(utils.HexToAddress(input.Address), encodedData, details)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, resp)
}
