package didaliasdto

import (
	"strconv"

	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
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

// RegisterTLDDTO represents the request to register a TLD
// @Description Data Transfer Object for registering a new top-level domain
type RegisterTLDDIDParam struct {
	TLD   string `json:"tld" validate:"required,domain_label"`
	Owner string `json:"owner" validate:"required,eth_address"`
}

type RegisterTLDDTO struct {
	PrivKey  string              `json:"priv_key" validate:"required,eth_private_key"`
	DIDParam RegisterTLDDIDParam `json:"did_param" validate:"required"`
}

func (dto *RegisterTLDDTO) Into() didaliasdomain.RegisterTLDInput {
	return didaliasdomain.RegisterTLDInput{
		PrivKey: dto.PrivKey,
		Input: aliasfactory.RegisterTLDParams{
			TLD:   dto.DIDParam.TLD,
			Owner: dto.DIDParam.Owner,
		},
	}
}
