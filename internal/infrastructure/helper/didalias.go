package internalhelper

import (
	"fmt"

	"github.com/meQlause/hara-core-blockchain-lib/utils"

	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
	didaliassdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didalias"
)

func (huc *InternalHelper) EncodeSetDIDRootStorageParam(input helperdo.EncodeSetDIDRootStorageParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"setDIDRootStorage",
		input.DIDParam.DIDRootStorage,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeRegisterDomainParam(input helperdo.EncodeRegisterDomainParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"registerDomain",
		input.DIDParam.Label,
		input.DIDParam.TLD,
		input.DIDParam.Period,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeRegisterSubdomainParam(input helperdo.EncodeRegisterSubdomainParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"registerSubdomain",
		input.DIDParam.Label,
		input.DIDParam.ParentDomain,
		input.DIDParam.Period,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeSetDIDParam(input helperdo.EncodeSetDIDParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"setDID",
		input.DIDParam.Node,
		input.DIDParam.DID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeExtendRegistrationParam(input helperdo.EncodeExtendRegistrationParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"extendRegistration",
		input.DIDParam.Node,
		input.DIDParam.Period,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeRevokeAliasParam(input helperdo.EncodeRevokeAliasParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"revokeAlias",
		input.DIDParam.Node,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeUnrevokeAliasParam(input helperdo.EncodeUnrevokeAliasParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"unrevokeAlias",
		input.DIDParam.Node,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeTransferAliasOwnershipParam(input helperdo.EncodeTransferAliasOwnershipParamInput) (string, error) {
	contractABI := didaliassdk.GetDIDAliasSDK().AliasFactory.ContractABI

	callData, err := contractABI.Pack(
		"transferAliasOwnership",
		input.DIDParam.Node,
		input.DIDParam.NewOwner,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}
