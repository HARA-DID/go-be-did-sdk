package didaliasdomain

import (
	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type ResolveInput struct {
	Node utils.Hash
}

type ResolveFromStringInput struct {
	Name string
}

type GetAliasStatusInput struct {
	Node utils.Hash
}

type GetAliasStatusFromStringInput struct {
	Name string
}

type GetOwnerInput struct {
	Node utils.Hash
}

type GetOwnerFromStringInput struct {
	Name string
}

type GetDIDInput struct {
	Node utils.Hash
}

type GetDIDFromStringInput struct {
	Name string
}

type NamehashInput struct {
	Name string
}

type GetRegistrationPeriodInput struct {
	Period aliasfactory.RegistrationPeriod
}

type SetDIDRootStorageInput struct {
	DIDRootStorage utils.Address
}

type RegisterTLDInput struct {
	PrivKey string
	Input   aliasfactory.RegisterTLDParams
}

type RegisterDomainInput struct {
	Label  string
	TLD    string
	Period aliasfactory.RegistrationPeriod
}

type RegisterSubdomainInput struct {
	Label        string
	ParentDomain string
	Period       aliasfactory.RegistrationPeriod
}

type SetDIDInput struct {
	Node utils.Hash
	DID  utils.Hash
}

type ExtendRegistrationInput struct {
	Node   utils.Hash
	Period aliasfactory.RegistrationPeriod
}

type RevokeAliasInput struct {
	Node utils.Hash
}

type UnrevokeAliasInput struct {
	Node utils.Hash
}

type TransferAliasOwnershipInput struct {
	Node     utils.Hash
	NewOwner utils.Address
}

type TransferOwnershipInput struct {
	NewOwner utils.Address
}
