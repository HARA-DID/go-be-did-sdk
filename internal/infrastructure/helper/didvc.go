package internalhelper

import (
	"fmt"

	"github.com/meQlause/hara-core-blockchain-lib/utils"

	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
	vcsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didvc"
)

func (huc *InternalHelper) EncodeIssueCredentialParam(input helperdo.EncodeIssueCredentialParamInput) (string, error) {
	contractABI := vcsdk.GetDIDVCSDK().VCFactory.ContractABI

	callData, err := contractABI.Pack(
		"issueCredential",
		input.VCParam.Option,
		input.VCParam.DIDRecipient,
		input.VCParam.ExpiredAt,
		input.VCParam.OffchainHash,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack issueCredential: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeBurnCredentialParam(input helperdo.EncodeBurnCredentialParamInput) (string, error) {
	contractABI := vcsdk.GetDIDVCSDK().VCFactory.ContractABI

	callData, err := contractABI.Pack(
		"burnCredential",
		input.VCParam.Option,
		input.VCParam.DID,
		input.VCParam.TokenID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack burnCredential: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeUpdateMetadataParam(input helperdo.EncodeUpdateMetadataParamInput) (string, error) {
	contractABI := vcsdk.GetDIDVCSDK().VCFactory.ContractABI

	callData, err := contractABI.Pack(
		"updateMetadata",
		input.VCParam.Option,
		input.VCParam.TokenID,
		input.VCParam.IsValid,
		input.VCParam.ExpiredAt,
		input.VCParam.OffchainHash,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack updateMetadata: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeRevokeCredentialParam(input helperdo.EncodeRevokeCredentialParamInput) (string, error) {
	contractABI := vcsdk.GetDIDVCSDK().VCFactory.ContractABI

	callData, err := contractABI.Pack(
		"revokeCredential",
		input.VCParam.Option,
		input.VCParam.TokenID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack revokeCredential: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeClaimCredentialParam(input helperdo.EncodeClaimCredentialParamInput) (string, error) {
	contractABI := vcsdk.GetDIDVCSDK().VCFactory.ContractABI

	callData, err := contractABI.Pack(
		"claimCredential",
		input.VCParam.Option,
		input.VCParam.TokenID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack claimCredential: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}
