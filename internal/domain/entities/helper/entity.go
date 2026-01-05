package helperdomain

import (
	"math/big"

	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
)

type StringToHex32Input struct {
	Input string
}

type EncodeCreateDIDParamInput struct {
	Address       string
	DIDParam      rootfactory.CreateDIDParam
	KeyIdentifier string
}

type EncodeUpdateDIDParamInput struct {
	Address       string
	DIDParam      rootfactory.UpdateDIDParams
	KeyIdentifier string
}

type EncodeDeactiveDIDParamInput struct {
	Address       string
	DIDIndex      *big.Int
	KeyIdentifier string
}

type EncodeReactiveDIDParamInput struct {
	Address       string
	DIDIndex      *big.Int
	KeyIdentifier string
}

type EncodeTransferDIDOwnerParamInput struct {
	Address       string
	DIDParam      rootfactory.TransferDIDOwnershipParams
	KeyIdentifier string
}

type EncodeStoreDataParamInput struct {
	Address       string
	DIDParam      rootfactory.StoreDataParams
	KeyIdentifier string
}

type EncodeDeleteDataParamInput struct {
	Address       string
	DIDParam      rootfactory.DeleteDataParams
	KeyIdentifier string
}

type EncodeAddKeyParamInput struct {
	Address       string
	DIDParam      rootfactory.StoreKeyParams
	KeyIdentifier string
}

type EncodeRemoveKeyParamInput struct {
	Address       string
	DIDParam      rootfactory.RemoveKeyParams
	KeyIdentifier string
}

type EncodeAddClaimParamInput struct {
	Address       string
	DIDParam      rootfactory.StoreClaimParams
	KeyIdentifier string
}

type EncodeRemoveClaimParamInput struct {
	Address       string
	DIDParam      rootfactory.RemoveClaimParams
	KeyIdentifier string
}
