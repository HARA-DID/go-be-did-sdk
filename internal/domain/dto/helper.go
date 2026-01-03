package dto

import (
	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
	helperdomain "github.com/meQlause/go-be-did/internal/domain/helper"
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
	DIDParam      CreateDIDDTO `json:"did_param" validate:"required"`
	KeyIdentifier string       `json:"key_identifier" validate:"required" example:"key1"`
}

type CreateDIDDTO struct {
	DID string `json:"did" validate:"required" example:"did:example:123456789"`
}

func (dto EncodeCreateDIDDTO) Into() helperdomain.EncodeCreateDIDParamInput {
	return helperdomain.EncodeCreateDIDParamInput{
		Address: dto.Address,
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
	DIDIndex uint64 `json:"did_index" validate:"required" example:"1"`
	URI      string `json:"uri" validate:"required" example:"https://example.com/did/metadata"`
}

func (dto EncodeUpdateDIDDTO) Into() helperdomain.EncodeUpdateDIDParamInput {
	return helperdomain.EncodeUpdateDIDParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.UpdateDIDParams{
			DIDIndex: dto.DIDParam.DIDIndex,
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
		Address:       dto.Address,
		DIDIndex:      utils.StringToBigInt(dto.DIDIndex),
		KeyIdentifier: dto.KeyIdentifier,
	}
}

type EncodeReactiveDIDDTO struct {
	Address       string `json:"address" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDIndex      string `json:"did_index" validate:"required" example:"1"`
	KeyIdentifier string `json:"key_identifier" example:"key1"`
}

func (dto EncodeReactiveDIDDTO) Into() helperdomain.EncodeReactiveDIDParamInput {
	return helperdomain.EncodeReactiveDIDParamInput{
		Address:       dto.Address,
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
	DIDIndex uint64 `json:"did_index" validate:"required" example:"1"`
	OldOwner string `json:"old_owner" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	NewOwner string `json:"new_owner" validate:"required,eth_address" example:"0x8ba1f109551bD432803012645Ac136ddd64DBA72"`
}

func (dto EncodeTransferDIDOwnerDTO) Into() helperdomain.EncodeTransferDIDOwnerParamInput {
	return helperdomain.EncodeTransferDIDOwnerParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.TransferDIDOwnershipParams{
			DIDIndex: dto.DIDParam.DIDIndex,
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
	DIDIndex uint64 `json:"did_index" validate:"required" example:"1"`
	Key      string `json:"key" validate:"required" example:"email"`
	Value    string `json:"value" validate:"required" example:"user@example.com"`
}

func (dto EncodeStoreDataDTO) Into() helperdomain.EncodeStoreDataParamInput {
	return helperdomain.EncodeStoreDataParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.StoreDataParams{
			DIDIndex: dto.DIDParam.DIDIndex,
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
	DIDIndex uint64 `json:"did_index" validate:"required" example:"1"`
	Key      string `json:"key" validate:"required" example:"email"`
}

func (dto EncodeDeleteDataDTO) Into() helperdomain.EncodeDeleteDataParamInput {
	return helperdomain.EncodeDeleteDataParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.DeleteDataParams{
			DIDIndex: dto.DIDParam.DIDIndex,
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
	DIDIndex      uint64 `json:"did_index" validate:"required" example:"1"`
	KeyDataHashed string `json:"key_data_hashed" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Purpose       uint8  `json:"purpose" validate:"required" example:"1"`
	KeyType       uint8  `json:"key_type" validate:"required" example:"1"`
}

func (dto EncodeAddKeyDTO) Into() helperdomain.EncodeAddKeyParamInput {
	var keyDataHashed [32]byte
	copy(keyDataHashed[:], dto.DIDParam.KeyDataHashed)

	return helperdomain.EncodeAddKeyParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.StoreKeyParams{
			DIDIndex:      dto.DIDParam.DIDIndex,
			KeyDataHashed: keyDataHashed,
			Purpose:       dto.DIDParam.Purpose,
			KeyType:       dto.DIDParam.KeyType,
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
	DIDIndex uint64 `json:"did_index" validate:"required" example:"1"`
	KeyData  string `json:"key_data" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto EncodeRemoveKeyDTO) Into() helperdomain.EncodeRemoveKeyParamInput {
	var keyData [32]byte
	copy(keyData[:], dto.DIDParam.KeyData)

	return helperdomain.EncodeRemoveKeyParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.RemoveKeyParams{
			DIDIndex:      dto.DIDParam.DIDIndex,
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
	DIDIndex  uint64 `json:"did_index" validate:"required" example:"1"`
	ClaimID   string `json:"claim_id" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Topic     uint8  `json:"topic" validate:"required" example:"1"`
	Issuer    string `json:"issuer" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Signature string `json:"signature" validate:"required,hex_data" example:"0xabcdef"`
	Data      string `json:"data" validate:"required,hex_data" example:"0x1234567890"`
	URI       string `json:"uri" validate:"required" example:"https://example.com/claims/1"`
}

func (dto EncodeAddClaimDTO) Into() helperdomain.EncodeAddClaimParamInput {
	var claimID [32]byte
	copy(claimID[:], dto.DIDParam.ClaimID)

	return helperdomain.EncodeAddClaimParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.StoreClaimParams{
			DIDIndex:  dto.DIDParam.DIDIndex,
			ClaimID:   claimID,
			Topic:     dto.DIDParam.Topic,
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
	DIDIndex uint64 `json:"did_index" validate:"required" example:"1"`
	ClaimID  string `json:"claim_id" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto EncodeRemoveClaimDTO) Into() helperdomain.EncodeRemoveClaimParamInput {
	var claimID [32]byte
	copy(claimID[:], dto.DIDParam.ClaimID)

	return helperdomain.EncodeRemoveClaimParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.RemoveClaimParams{
			DIDIndex: dto.DIDParam.DIDIndex,
			ClaimID:  claimID,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}
