package internalhelper

import (
	"fmt"

	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	backendutils "github.com/meQlause/go-be-did/utils"

	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type InternalHelper struct {
}

func NewInternalHelper() (repository.HelperRepository, error) {
	return &InternalHelper{}, nil
}

func (h *InternalHelper) StringToHex32(input helperdo.StringToHex32Input) string {
	return utils.StringToHex32(input.Input)
}

func (huc *InternalHelper) EncodeCreateDIDParam(input helperdo.EncodeCreateDIDParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI
	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("string").Value(input.DIDParam.DID))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeCreateDID,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeUpdateDIDParam(input helperdo.EncodeUpdateDIDParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("string").Value(input.DIDParam.URI))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeUpdateDID,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeDeactiveDIDParam(input helperdo.EncodeDeactiveDIDParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDIndex))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeDeactivateDID,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeReactiveDIDParam(input helperdo.EncodeReactiveDIDParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDIndex))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeReactivateDID,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeTransferDIDOwnerParam(input helperdo.EncodeTransferDIDOwnerParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("address").Value(input.DIDParam.NewOwner))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeTransferDID,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeStoreDataParam(input helperdo.EncodeStoreDataParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("string").Value(input.DIDParam.Key).
		Type("string").Value(input.DIDParam.Value))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeStoreData,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeDeleteDataParam(input helperdo.EncodeDeleteDataParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("string").Value(input.DIDParam.Key))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeDeleteData,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeAddKeyParam(input helperdo.EncodeAddKeyParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("bytes32").Value(input.DIDParam.KeyDataHashed).
		Type("string").Value(input.DIDParam.KeyIdentifierDst).
		Type("uint8").Value(input.DIDParam.Purpose).
		Type("uint8").Value(input.DIDParam.KeyType))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeAddKey,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeRemoveKeyParam(input helperdo.EncodeRemoveKeyParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("bytes32").Value(input.DIDParam.KeyDataHashed))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeRemoveKey,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeAddClaimParam(input helperdo.EncodeAddClaimParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("uint8").Value(input.DIDParam.Topic).
		Type("bytes").Value(input.DIDParam.ClaimID).
		Type("string").Value(input.DIDParam.URI).
		Type("bytes").Value(input.DIDParam.Signature))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeAddClaim,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}

func (huc *InternalHelper) EncodeRemoveClaimParam(input helperdo.EncodeRemoveClaimParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	encodedData := utils.EncodeArgs(config.Network().ArgBuilder().
		Type("uint256").Value(input.DIDParam.DIDIndex).
		Type("bytes").Value(input.DIDParam.ClaimID))

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeRemoveClaim,
		encodedData,
		input.KeyIdentifier,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + utils.EncodeToString(callData), nil
}
