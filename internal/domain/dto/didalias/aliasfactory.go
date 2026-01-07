package didaliasdto

import (
	"strconv"

	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// ResolveDTO represents the request payload for resolving a node to DID
// @Description Data Transfer Object for resolving a node hash to its associated DID
type ResolveDTO struct {
	Node string `query:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *ResolveDTO) Into() didaliasdomain.ResolveInput {
	return didaliasdomain.ResolveInput{
		Node: utils.HexToHash(dto.Node),
	}
}

// ResolveFromStringDTO represents the request payload for resolving a domain name to DID
// @Description Data Transfer Object for resolving a domain name string to its associated DID
type ResolveFromStringDTO struct {
	Name string `query:"name" validate:"required" example:"example.tld"`
}

func (dto *ResolveFromStringDTO) Into() didaliasdomain.ResolveFromStringInput {
	return didaliasdomain.ResolveFromStringInput{
		Name: dto.Name,
	}
}

// GetAliasStatusDTO represents the request payload for getting alias status by node
// @Description Data Transfer Object for retrieving alias status (expiry, revoked, valid) by node hash
type GetAliasStatusDTO struct {
	Node string `query:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetAliasStatusDTO) Into() didaliasdomain.GetAliasStatusInput {
	return didaliasdomain.GetAliasStatusInput{
		Node: utils.HexToHash(dto.Node),
	}
}

// GetAliasStatusFromStringDTO represents the request payload for getting alias status by name
// @Description Data Transfer Object for retrieving alias status by domain name string
type GetAliasStatusFromStringDTO struct {
	Name string `query:"name" validate:"required" example:"example.tld"`
}

func (dto *GetAliasStatusFromStringDTO) Into() didaliasdomain.GetAliasStatusFromStringInput {
	return didaliasdomain.GetAliasStatusFromStringInput{
		Name: dto.Name,
	}
}

// GetOwnerDTO represents the request payload for getting owner by node
// @Description Data Transfer Object for retrieving the owner address of a node
type GetOwnerDTO struct {
	Node string `query:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetOwnerDTO) Into() didaliasdomain.GetOwnerInput {
	return didaliasdomain.GetOwnerInput{
		Node: utils.HexToHash(dto.Node),
	}
}

// GetOwnerFromStringDTO represents the request payload for getting owner by name
// @Description Data Transfer Object for retrieving the owner address by domain name
type GetOwnerFromStringDTO struct {
	Name string `query:"name" validate:"required" example:"example.tld"`
}

func (dto *GetOwnerFromStringDTO) Into() didaliasdomain.GetOwnerFromStringInput {
	return didaliasdomain.GetOwnerFromStringInput{
		Name: dto.Name,
	}
}

// GetDIDDTO represents the request payload for getting DID by node
// @Description Data Transfer Object for retrieving DID hash associated with a node
type GetDIDDTO struct {
	Node string `query:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetDIDDTO) Into() didaliasdomain.GetDIDInput {
	return didaliasdomain.GetDIDInput{
		Node: utils.HexToHash(dto.Node),
	}
}

// GetDIDFromStringDTO represents the request payload for getting DID by name
// @Description Data Transfer Object for retrieving DID hash by domain name
type GetDIDFromStringDTO struct {
	Name string `query:"name" validate:"required" example:"example.tld"`
}

func (dto *GetDIDFromStringDTO) Into() didaliasdomain.GetDIDFromStringInput {
	return didaliasdomain.GetDIDFromStringInput{
		Name: dto.Name,
	}
}

// NamehashDTO represents the request payload for computing namehash
// @Description Data Transfer Object for computing the namehash of a domain name
type NamehashDTO struct {
	Name string `query:"name" validate:"required" example:"example.tld"`
}

func (dto *NamehashDTO) Into() didaliasdomain.NamehashInput {
	return didaliasdomain.NamehashInput{
		Name: dto.Name,
	}
}

// GetRegistrationPeriodDTO represents the request payload for getting registration period duration
// @Description Data Transfer Object for retrieving the duration of a registration period
type GetRegistrationPeriodDTO struct {
	Period string `query:"period" validate:"required,registration_period" example:"0"`
}

func (dto *GetRegistrationPeriodDTO) Into() didaliasdomain.GetRegistrationPeriodInput {
	period, _ := strconv.ParseUint(dto.Period, 10, 8)
	return didaliasdomain.GetRegistrationPeriodInput{
		Period: aliasfactory.RegistrationPeriod(period),
	}
}

// SetDIDRootStorageDTO represents the request payload for setting DID root storage
// @Description Data Transfer Object for setting the DID root storage contract address
type SetDIDRootStorageDTO struct {
	DIDRootStorage string `json:"did_root_storage" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

func (dto *SetDIDRootStorageDTO) Into() didaliasdomain.SetDIDRootStorageInput {
	return didaliasdomain.SetDIDRootStorageInput{
		DIDRootStorage: utils.HexToAddress(dto.DIDRootStorage),
	}
}

// RegisterTLDDTO represents the request payload for registering a TLD
// @Description Data Transfer Object for registering a top-level domain
type RegisterTLDDTO struct {
	TLD   string `json:"tld" validate:"required" example:"tld"`
	Owner string `json:"owner" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

func (dto *RegisterTLDDTO) Into() didaliasdomain.RegisterTLDInput {
	return didaliasdomain.RegisterTLDInput{
		TLD:   dto.TLD,
		Owner: utils.HexToAddress(dto.Owner),
	}
}

// RegisterDomainDTO represents the request payload for registering a domain
// @Description Data Transfer Object for registering a domain under a TLD
type RegisterDomainDTO struct {
	Label  string `json:"label" validate:"required" example:"example"`
	TLD    string `json:"tld" validate:"required" example:"tld"`
	Period string `json:"period" validate:"required,registration_period" example:"0"`
}

func (dto *RegisterDomainDTO) Into() didaliasdomain.RegisterDomainInput {
	period, _ := strconv.ParseUint(dto.Period, 10, 8)
	return didaliasdomain.RegisterDomainInput{
		Label:  dto.Label,
		TLD:    dto.TLD,
		Period: aliasfactory.RegistrationPeriod(period),
	}
}

// RegisterSubdomainDTO represents the request payload for registering a subdomain
// @Description Data Transfer Object for registering a subdomain under a parent domain
type RegisterSubdomainDTO struct {
	Label        string `json:"label" validate:"required" example:"sub"`
	ParentDomain string `json:"parent_domain" validate:"required" example:"example.tld"`
	Period       string `json:"period" validate:"required,registration_period" example:"0"`
}

func (dto *RegisterSubdomainDTO) Into() didaliasdomain.RegisterSubdomainInput {
	period, _ := strconv.ParseUint(dto.Period, 10, 8)
	return didaliasdomain.RegisterSubdomainInput{
		Label:        dto.Label,
		ParentDomain: dto.ParentDomain,
		Period:       aliasfactory.RegistrationPeriod(period),
	}
}

// SetDIDDTO represents the request payload for setting a DID for a node
// @Description Data Transfer Object for associating a DID with a node/domain
type SetDIDDTO struct {
	Node string `json:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	DID  string `json:"did" validate:"required,hex32" example:"0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"`
}

func (dto *SetDIDDTO) Into() didaliasdomain.SetDIDInput {
	return didaliasdomain.SetDIDInput{
		Node: utils.HexToHash(dto.Node),
		DID:  utils.HexToHash(dto.DID),
	}
}

// ExtendRegistrationDTO represents the request payload for extending registration
// @Description Data Transfer Object for extending the registration period of a node
type ExtendRegistrationDTO struct {
	Node   string `json:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Period string `json:"period" validate:"required,registration_period" example:"0"`
}

func (dto *ExtendRegistrationDTO) Into() didaliasdomain.ExtendRegistrationInput {
	period, _ := strconv.ParseUint(dto.Period, 10, 8)
	return didaliasdomain.ExtendRegistrationInput{
		Node:   utils.HexToHash(dto.Node),
		Period: aliasfactory.RegistrationPeriod(period),
	}
}

// RevokeAliasDTO represents the request payload for revoking an alias
// @Description Data Transfer Object for revoking a node/alias
type RevokeAliasDTO struct {
	Node string `json:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *RevokeAliasDTO) Into() didaliasdomain.RevokeAliasInput {
	return didaliasdomain.RevokeAliasInput{
		Node: utils.HexToHash(dto.Node),
	}
}

// UnrevokeAliasDTO represents the request payload for unrevoking an alias
// @Description Data Transfer Object for unrevoking a previously revoked node/alias
type UnrevokeAliasDTO struct {
	Node string `json:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *UnrevokeAliasDTO) Into() didaliasdomain.UnrevokeAliasInput {
	return didaliasdomain.UnrevokeAliasInput{
		Node: utils.HexToHash(dto.Node),
	}
}

// TransferAliasOwnershipDTO represents the request payload for transferring alias ownership
// @Description Data Transfer Object for transferring ownership of a node/alias to a new owner
type TransferAliasOwnershipDTO struct {
	Node     string `json:"node" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	NewOwner string `json:"new_owner" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

func (dto *TransferAliasOwnershipDTO) Into() didaliasdomain.TransferAliasOwnershipInput {
	return didaliasdomain.TransferAliasOwnershipInput{
		Node:     utils.HexToHash(dto.Node),
		NewOwner: utils.HexToAddress(dto.NewOwner),
	}
}

// TransferOwnershipDTO represents the request payload for transferring contract ownership
// @Description Data Transfer Object for transferring ownership of the entire contract
type TransferOwnershipDTO struct {
	NewOwner string `json:"new_owner" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

func (dto *TransferOwnershipDTO) Into() didaliasdomain.TransferOwnershipInput {
	return didaliasdomain.TransferOwnershipInput{
		NewOwner: utils.HexToAddress(dto.NewOwner),
	}
}
