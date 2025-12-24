package internalhelper

import (
	"encoding/hex"
	"fmt"

	"github.com/meQlause/go-be-did/internal/config"
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	"github.com/meQlause/go-be-did/internal/repository"
	backendutils "github.com/meQlause/go-be-did/utils"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type InternalHelper struct {
}

func NewInternalHelper() (repository.HelperRepository, error) {
	return &InternalHelper{}, nil
}

func (h *InternalHelper) StringToByte32(input helperdo.StringToByte32Input) [32]byte {
	return utils.StringToByte32(input.Input)
}

func (huc *InternalHelper) EncodeCreateDIDParam(createDIDParam helperdo.EncodeCreateDIDParamInput) string {
	encodedData := utils.EncodeArgs(config.Blockchain().Network.ArgBuilder().
		Type("string").Value(createDIDParam.DIDParam.DID))

	contractABI := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI

	callData, err := contractABI.Pack(
		"callExternal",
		backendutils.TypeCreateDID,
		encodedData,
		createDIDParam.KeyIdentifier,
	)

	if err != nil {
		panic(fmt.Errorf("failed to pack: %w", err))
	}

	return "0x" + hex.EncodeToString(callData)
}
