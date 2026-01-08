package helperdto

import (
	"strconv"

	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
	helperdomain "github.com/meQlause/go-be-did/internal/domain/entities/helper"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type StringToHex32DTO struct {
	Input string `json:"input" validate:"required" example:"example_string"`
}

func (dto *StringToHex32DTO) Into() helperdomain.StringToHex32Input {
	return helperdomain.StringToHex32Input{
		Input: dto.Input,
	}
}

type EncodeCreateDIDDTO struct {
	Address       string       `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      CreateDIDDTO `json:"did_param"`
	KeyIdentifier string       `json:"key_identifier" example:"key1"`
}

type CreateDIDDTO struct {
	DID string `json:"did" example:"did:example:123456789"`
}

func (dto EncodeCreateDIDDTO) Into() helperdomain.EncodeCreateDIDParamInput {
	return helperdomain.EncodeCreateDIDParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.CreateDIDParam{
			DID: dto.DIDParam.DID,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeUpdateDIDDTO struct {
	Address       string       `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      UpdateDIDDTO `json:"did_param"`
	KeyIdentifier string       `json:"key_identifier" example:"key1"`
}

type UpdateDIDDTO struct {
	DIDIndex string `json:"did_index" validate:"required,uint64" example:"1"`
	URI      string `json:"uri" validate:"required" example:"https://example.com/did/metadata"`
}

func (dto EncodeUpdateDIDDTO) Into() helperdomain.EncodeUpdateDIDParamInput {
	return helperdomain.EncodeUpdateDIDParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.UpdateDIDParams{
			DIDIndex: utils.StringToBigInt(dto.DIDParam.DIDIndex),
			URI:      dto.DIDParam.URI,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeDeactiveDIDDTO struct {
	Address       string `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDIndex      string `json:"did_index" validate:"required,uint64" example:"1"`
	KeyIdentifier string `json:"key_identifier" example:"key1"`
}

func (dto EncodeDeactiveDIDDTO) Into() helperdomain.EncodeDeactiveDIDParamInput {
	return helperdomain.EncodeDeactiveDIDParamInput{
		Address:       utils.HexToAddress(dto.Address),
		DIDIndex:      utils.StringToBigInt(dto.DIDIndex),
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeReactiveDIDDTO struct {
	Address       string `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDIndex      string `json:"did_index" validate:"required,uint64" example:"1"`
	KeyIdentifier string `json:"key_identifier" example:"key1"`
}

func (dto EncodeReactiveDIDDTO) Into() helperdomain.EncodeReactiveDIDParamInput {
	return helperdomain.EncodeReactiveDIDParamInput{
		Address:       utils.HexToAddress(dto.Address),
		DIDIndex:      utils.StringToBigInt(dto.DIDIndex),
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeTransferDIDOwnerDTO struct {
	Address       string                  `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      TransferDIDOwnershipDTO `json:"did_param"`
	KeyIdentifier string                  `json:"key_identifier" example:"key1"`
}

type TransferDIDOwnershipDTO struct {
	DIDIndex string `json:"did_index" validate:"required,uint64" example:"1"`
	OldOwner string `json:"old_owner" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	NewOwner string `json:"new_owner" validate:"required,eth_address" example:"0x8ba1f109551bD432803012645Ac136ddd64DBA72"`
}

func (dto EncodeTransferDIDOwnerDTO) Into() helperdomain.EncodeTransferDIDOwnerParamInput {
	return helperdomain.EncodeTransferDIDOwnerParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.TransferDIDOwnershipParams{
			DIDIndex: utils.StringToBigInt(dto.DIDParam.DIDIndex),
			NewOwner: utils.HexToAddress(dto.DIDParam.NewOwner),
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeStoreDataDTO struct {
	Address       string       `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      StoreDataDTO `json:"did_param"`
	KeyIdentifier string       `json:"key_identifier" example:"key1"`
}

type StoreDataDTO struct {
	DIDIndex string `json:"did_index" validate:"required,uint64" example:"1"`
	Key      string `json:"key" validate:"required" example:"email"`
	Value    string `json:"value" validate:"required" example:"user@example.com"`
}

func (dto EncodeStoreDataDTO) Into() helperdomain.EncodeStoreDataParamInput {
	return helperdomain.EncodeStoreDataParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.StoreDataParams{
			DIDIndex: utils.StringToBigInt(dto.DIDParam.DIDIndex),
			Key:      dto.DIDParam.Key,
			Value:    dto.DIDParam.Value,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeDeleteDataDTO struct {
	Address       string        `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      DeleteDataDTO `json:"did_param"`
	KeyIdentifier string        `json:"key_identifier" example:"key1"`
}

type DeleteDataDTO struct {
	DIDIndex string `json:"did_index" validate:"required,uint64" example:"1"`
	Key      string `json:"key" validate:"required" example:"email"`
}

func (dto EncodeDeleteDataDTO) Into() helperdomain.EncodeDeleteDataParamInput {
	return helperdomain.EncodeDeleteDataParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.DeleteDataParams{
			DIDIndex: utils.StringToBigInt(dto.DIDParam.DIDIndex),
			Key:      dto.DIDParam.Key,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeAddKeyDTO struct {
	Address       string      `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      StoreKeyDTO `json:"did_param"`
	KeyIdentifier string      `json:"key_identifier" example:"key1"`
}

type StoreKeyDTO struct {
	DIDIndex      string `json:"did_index" validate:"required,uint64" example:"1"`
	KeyDataHashed string `json:"key_data_hashed" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Purpose       string `json:"purpose" validate:"required,uint8" example:"1"`
	KeyType       string `json:"key_type" validate:"required,uint8" example:"1"`
}

func (dto EncodeAddKeyDTO) Into() helperdomain.EncodeAddKeyParamInput {
	var keyDataHashed [32]byte
	copy(keyDataHashed[:], dto.DIDParam.KeyDataHashed)
	keyType, _ := strconv.ParseUint(dto.DIDParam.KeyType, 10, 8)
	purpose, _ := strconv.ParseUint(dto.DIDParam.Purpose, 10, 8)

	return helperdomain.EncodeAddKeyParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.StoreKeyParams{
			DIDIndex:      utils.StringToBigInt(dto.DIDParam.DIDIndex),
			KeyDataHashed: keyDataHashed,
			Purpose:       uint8(purpose),
			KeyType:       uint8(keyType),
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeRemoveKeyDTO struct {
	Address       string       `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      RemoveKeyDTO `json:"did_param"`
	KeyIdentifier string       `json:"key_identifier" example:"key1"`
}

type RemoveKeyDTO struct {
	DIDIndex string `json:"did_index" validate:"required,uint64" example:"1"`
	KeyData  string `json:"key_data" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto EncodeRemoveKeyDTO) Into() helperdomain.EncodeRemoveKeyParamInput {
	var keyData [32]byte
	copy(keyData[:], dto.DIDParam.KeyData)

	return helperdomain.EncodeRemoveKeyParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.RemoveKeyParams{
			DIDIndex:      utils.StringToBigInt(dto.DIDParam.DIDIndex),
			KeyDataHashed: keyData,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeAddClaimDTO struct {
	Address       string        `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      StoreClaimDTO `json:"did_param"`
	KeyIdentifier string        `json:"key_identifier" example:"key1"`
}

type StoreClaimDTO struct {
	DIDIndex  string `json:"did_index" validate:"required" example:"1"`
	ClaimID   string `json:"claim_id" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Topic     string `json:"topic" validate:"required,uint8" example:"1"`
	Issuer    string `json:"issuer" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Signature string `json:"signature" validate:"required,hex_data" example:"0xabcdef"`
	Data      string `json:"data" validate:"required,hex_data" example:"0x1234567890"`
	URI       string `json:"uri" validate:"required" example:"https://example.com/claims/1"`
}

func (dto EncodeAddClaimDTO) Into() helperdomain.EncodeAddClaimParamInput {
	var claimID [32]byte
	copy(claimID[:], dto.DIDParam.ClaimID)
	topic, _ := strconv.ParseUint(dto.DIDParam.Topic, 10, 8)

	return helperdomain.EncodeAddClaimParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.StoreClaimParams{
			DIDIndex:  utils.StringToBigInt(dto.DIDParam.DIDIndex),
			ClaimID:   claimID,
			Topic:     uint8(topic),
			Issuer:    utils.HexToAddress(dto.DIDParam.Issuer),
			Signature: []byte(dto.DIDParam.Signature),
			Data:      []byte(dto.DIDParam.Data),
			URI:       dto.DIDParam.URI,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeRemoveClaimDTO struct {
	Address       string         `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      RemoveClaimDTO `json:"did_param"`
	KeyIdentifier string         `json:"key_identifier" example:"key1"`
}

type RemoveClaimDTO struct {
	DIDIndex string `json:"did_index" validate:"required,uint64" example:"1"`
	ClaimID  string `json:"claim_id" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto EncodeRemoveClaimDTO) Into() helperdomain.EncodeRemoveClaimParamInput {
	var claimID [32]byte
	copy(claimID[:], dto.DIDParam.ClaimID)

	return helperdomain.EncodeRemoveClaimParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: rootfactory.RemoveClaimParams{
			DIDIndex: utils.StringToBigInt(dto.DIDParam.DIDIndex),
			ClaimID:  claimID,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}

// SetDIDRootStorageDTO represents the request to set DID root storage
// @Description Data Transfer Object for setting the DID root storage address
type SetDIDRootStorageDIDParam struct {
	DIDRootStorage string `json:"did_root_storage" validate:"required,eth_address"`
}

type EncodeSetDIDRootStorageDTO struct {
	Address  string                    `json:"address" validate:"required,eth_address"`
	DIDParam SetDIDRootStorageDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeSetDIDRootStorageDTO) Into() helperdomain.EncodeSetDIDRootStorageParamInput {
	return helperdomain.EncodeSetDIDRootStorageParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.SetDIDRootStorageParams{
			DIDRootStorage: dto.DIDParam.DIDRootStorage,
		},
	}
}

// RegisterDomainDTO represents the request to register a domain
// @Description Data Transfer Object for registering a new domain under a TLD
type RegisterDomainDIDParam struct {
	Label  string `json:"label" validate:"required,domain_label"`
	TLD    string `json:"tld" validate:"required,domain_label"`
	Owner  string `json:"owner" validate:"required,eth_address"`
	Period string `json:"period" validate:"required,registration_period"`
}

type EncodeRegisterDomainDTO struct {
	Address  string                 `json:"address" validate:"required,eth_address"`
	DIDParam RegisterDomainDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeRegisterDomainDTO) Into() helperdomain.EncodeRegisterDomainParamInput {
	period, _ := strconv.ParseUint(dto.DIDParam.Period, 10, 8)
	return helperdomain.EncodeRegisterDomainParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.RegisterDomainParams{
			Label:  dto.DIDParam.Label,
			TLD:    dto.DIDParam.TLD,
			Owner:  dto.DIDParam.Owner,
			Period: aliasfactory.RegistrationPeriod(period),
		},
	}
}

// RegisterSubdomainDTO represents the request to register a subdomain
// @Description Data Transfer Object for registering a new subdomain under a parent domain
type RegisterSubdomainDIDParam struct {
	Label        string `json:"label" validate:"required,domain_label"`
	ParentDomain string `json:"parent_domain" validate:"required,domain_name"`
	Period       string `json:"period" validate:"required,registration_period"`
}

type EncodeRegisterSubdomainDTO struct {
	Address  string                    `json:"address" validate:"required,eth_address"`
	DIDParam RegisterSubdomainDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeRegisterSubdomainDTO) Into() helperdomain.EncodeRegisterSubdomainParamInput {
	period, _ := strconv.ParseUint(dto.DIDParam.Period, 10, 8)
	return helperdomain.EncodeRegisterSubdomainParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.RegisterSubdomainParams{
			Label:        dto.DIDParam.Label,
			ParentDomain: dto.DIDParam.ParentDomain,
			Period:       aliasfactory.RegistrationPeriod(period),
		},
	}
}

// SetDIDDTO represents the request to set a DID for a node
// @Description Data Transfer Object for setting a DID (Decentralized Identifier) to a node
type SetDIDDIDParam struct {
	Node string `json:"node" validate:"required,hex32"`
	DID  string `json:"did" validate:"required,hex32"`
}

type EncodeSetDIDDTO struct {
	Address  string         `json:"address" validate:"required,eth_address"`
	DIDParam SetDIDDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeSetDIDDTO) Into() helperdomain.EncodeSetDIDParamInput {
	return helperdomain.EncodeSetDIDParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.SetDIDParams{
			Node: utils.HexToHash(dto.DIDParam.Node),
			DID:  utils.HexToHash(dto.DIDParam.DID),
		},
	}
}

// ExtendRegistrationDTO represents the request to extend registration
// @Description Data Transfer Object for extending the registration period of an alias
type ExtendRegistrationDIDParam struct {
	Node   string `json:"node" validate:"required,hex32"`
	Period string `json:"period" validate:"required,registration_period"`
}

type EncodeExtendRegistrationDTO struct {
	Address  string                     `json:"address" validate:"required,eth_address"`
	DIDParam ExtendRegistrationDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeExtendRegistrationDTO) Into() helperdomain.EncodeExtendRegistrationParamInput {
	period, _ := strconv.ParseUint(dto.DIDParam.Period, 10, 8)
	return helperdomain.EncodeExtendRegistrationParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.ExtendRegistrationParams{
			Node:   utils.HexToHash(dto.DIDParam.Node),
			Period: aliasfactory.RegistrationPeriod(period),
		},
	}
}

// RevokeAliasDTO represents the request to revoke an alias
// @Description Data Transfer Object for revoking an alias
type RevokeAliasDIDParam struct {
	Node string `json:"node" validate:"required,hex32"`
}

type EncodeRevokeAliasDTO struct {
	Address  string              `json:"address" validate:"required,eth_address"`
	DIDParam RevokeAliasDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeRevokeAliasDTO) Into() helperdomain.EncodeRevokeAliasParamInput {
	return helperdomain.EncodeRevokeAliasParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.NodeOnlyParams{
			Node: utils.HexToHash(dto.DIDParam.Node),
		},
	}
}

// UnrevokeAliasDTO represents the request to unrevoke an alias
// @Description Data Transfer Object for unrevoking a previously revoked alias
type UnrevokeAliasDIDParam struct {
	Node string `json:"node" validate:"required,hex32"`
}

type EncodeUnrevokeAliasDTO struct {
	Address  string                `json:"address" validate:"required,eth_address"`
	DIDParam UnrevokeAliasDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeUnrevokeAliasDTO) Into() helperdomain.EncodeUnrevokeAliasParamInput {
	return helperdomain.EncodeUnrevokeAliasParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.NodeOnlyParams{
			Node: utils.HexToHash(dto.DIDParam.Node),
		},
	}
}

// TransferAliasOwnershipDTO represents the request to transfer alias ownership
// @Description Data Transfer Object for transferring ownership of an alias to a new owner
type TransferAliasOwnershipDIDParam struct {
	Node     string `json:"node" validate:"required,hex32"`
	NewOwner string `json:"new_owner" validate:"required,eth_address"`
}

type EncodeTransferAliasOwnershipDTO struct {
	Address  string                         `json:"address" validate:"required,eth_address"`
	DIDParam TransferAliasOwnershipDIDParam `json:"did_param" validate:"required"`
}

func (dto *EncodeTransferAliasOwnershipDTO) Into() helperdomain.EncodeTransferAliasOwnershipParamInput {
	return helperdomain.EncodeTransferAliasOwnershipParamInput{
		Address: utils.HexToAddress(dto.Address),
		DIDParam: aliasfactory.TransferAliasOwnershipParams{
			Node:     utils.HexToHash(dto.DIDParam.Node),
			NewOwner: dto.DIDParam.NewOwner,
		},
	}
}
