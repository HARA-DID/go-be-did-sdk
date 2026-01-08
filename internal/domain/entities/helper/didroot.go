package helperdomain

import (
	"math/big"

	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type EncodeCreateDIDParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.CreateDIDParam
	KeyIdentifier string
}

type EncodeUpdateDIDParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.UpdateDIDParams
	KeyIdentifier string
}

type EncodeDeactiveDIDParamInput struct {
	Address       utils.Address
	DIDIndex      *big.Int
	KeyIdentifier string
}

type EncodeReactiveDIDParamInput struct {
	Address       utils.Address
	DIDIndex      *big.Int
	KeyIdentifier string
}

type EncodeTransferDIDOwnerParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.TransferDIDOwnershipParams
	KeyIdentifier string
}

type EncodeStoreDataParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.StoreDataParams
	KeyIdentifier string
}

type EncodeDeleteDataParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.DeleteDataParams
	KeyIdentifier string
}

type EncodeAddKeyParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.StoreKeyParams
	KeyIdentifier string
}

type EncodeRemoveKeyParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.RemoveKeyParams
	KeyIdentifier string
}

type EncodeAddClaimParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.StoreClaimParams
	KeyIdentifier string
}

type EncodeRemoveClaimParamInput struct {
	Address       utils.Address
	DIDParam      rootfactory.RemoveClaimParams
	KeyIdentifier string
}
