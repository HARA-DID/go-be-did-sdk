package internalhelper

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	backendutils "github.com/meQlause/go-be-did/utils"
)

type InternalHelper struct {
}

func NewInternalHelper() (repository.HelperRepository, error) {
	return &InternalHelper{}, nil
}

func (h *InternalHelper) StringToByte32(input helperdo.StringToByte32Input) [32]byte {
	return utils.StringToByte32(input.Input)
}

func (huc *InternalHelper) EncodeCreateDIDParam(createDIDParam helperdo.EncodeCreateDIDParamInput) (string, error) {
	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI
	stringType, err := abi.NewType("string", "", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create string type: %w", err)
	}

	arguments := abi.Arguments{
		{Type: stringType},
	}

	encodedData, err := arguments.Pack(createDIDParam.DIDParam.DID)
	if err != nil {
		return "", fmt.Errorf("failed to encode DID: %w", err)
	}
	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeCreateDID,
		encodedData,
		createDIDParam.KeyIdentifier,
	)

	if err != nil {
		return "", fmt.Errorf("failed to pack: %w", err)
	}

	return "0x" + hex.EncodeToString(callData), nil
}
