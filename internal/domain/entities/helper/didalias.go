package helperdomain

import (
	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type EncodeSetDIDRootStorageParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.SetDIDRootStorageParams
}

type EncodeRegisterDomainParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.RegisterDomainParams
}

type EncodeRegisterSubdomainParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.RegisterSubdomainParams
}

type EncodeSetDIDParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.SetDIDParams
}

type EncodeExtendRegistrationParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.ExtendRegistrationParams
}

type EncodeRevokeAliasParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.NodeOnlyParams
}

type EncodeUnrevokeAliasParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.NodeOnlyParams
}

type EncodeTransferAliasOwnershipParamInput struct {
	Address  utils.Address
	DIDParam aliasfactory.TransferAliasOwnershipParams
}
