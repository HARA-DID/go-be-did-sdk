package repository

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

type HelperRepository interface {
	StringToHex32(input helperdo.StringToHex32Input) string
	EncodeCreateDIDParam(input helperdo.EncodeCreateDIDParamInput) (string, error)
	EncodeUpdateDIDParam(input helperdo.EncodeUpdateDIDParamInput) (string, error)
	EncodeDeactiveDIDParam(input helperdo.EncodeDeactiveDIDParamInput) (string, error)
	EncodeReactiveDIDParam(input helperdo.EncodeReactiveDIDParamInput) (string, error)
	EncodeTransferDIDOwnerParam(input helperdo.EncodeTransferDIDOwnerParamInput) (string, error)
	EncodeStoreDataParam(input helperdo.EncodeStoreDataParamInput) (string, error)
	EncodeDeleteDataParam(input helperdo.EncodeDeleteDataParamInput) (string, error)
	EncodeAddKeyParam(input helperdo.EncodeAddKeyParamInput) (string, error)
	EncodeRemoveKeyParam(input helperdo.EncodeRemoveKeyParamInput) (string, error)
	EncodeAddClaimParam(input helperdo.EncodeAddClaimParamInput) (string, error)
	EncodeRemoveClaimParam(input helperdo.EncodeRemoveClaimParamInput) (string, error)
}
