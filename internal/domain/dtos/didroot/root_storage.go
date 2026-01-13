package didrootdto

import (
	"strconv"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// GetDataDTO represents the request payload for retrieving data by hash
// @Description Data Transfer Object for getting data from the blockchain by its hash
type GetDataDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Key     string `query:"key" validate:"required" example:"email"`
}

func (dto *GetDataDTO) Into() didrootdomain.GetDataInput {
	didHash := utils.HexToHash(dto.DIDHash)
	data := append(didHash[:], []byte(dto.Key)...)
	keyCodeBytes := utils.Keccak256(data)

	var keyCode utils.Hash
	copy(keyCode[:], keyCodeBytes)
	return didrootdomain.GetDataInput{
		Hash: keyCode,
	}
}

// ResolveDIDDTO represents the request payload for resolving a DID
// @Description Data Transfer Object for resolving a Decentralized Identifier to its document
type ResolveDIDDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *ResolveDIDDTO) Into() didrootdomain.ResolveDIDInput {
	return didrootdomain.ResolveDIDInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// VerifyDIDOwnershipDTO represents the request payload for verifying DID ownership
// @Description Data Transfer Object for verifying if an address owns a specific DID
type VerifyDIDOwnershipDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Owner   string `query:"user_owner" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

// Into converts VerifyDIDOwnershipDTO to domain input model
func (dto *VerifyDIDOwnershipDTO) Into() didrootdomain.VerifyDIDOwnershipInput {
	return didrootdomain.VerifyDIDOwnershipInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
		Owner:   utils.HexToAddress(dto.Owner),
	}
}

// GetKeyDTO represents the request payload for retrieving a specific key
// @Description Data Transfer Object for getting a cryptographic key by DID and key hash
type GetKeyDTO struct {
	DIDHash       string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	KeyDataHashed string `query:"key_data_hashed" validate:"required,hex32" example:"0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"`
}

func (dto *GetKeyDTO) Into() didrootdomain.GetKeyInput {
	return didrootdomain.GetKeyInput{
		DIDHash:       utils.HexToHash(dto.DIDHash),
		KeyDataHashed: utils.HexToHash(dto.KeyDataHashed),
	}
}

// GetKeysByDIDDTO represents the request payload for retrieving all keys for a DID
// @Description Data Transfer Object for getting all cryptographic keys associated with a DID
type GetKeysByDIDDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetKeysByDIDDTO) Into() didrootdomain.GetKeysByDIDInput {
	return didrootdomain.GetKeysByDIDInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// GetClaimDTO represents the request payload for retrieving a specific claim
// @Description Data Transfer Object for getting a verifiable claim by DID and claim ID
type GetClaimDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	ClaimID string `query:"claim_hash_id" validate:"required,hex32" example:"0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"`
}

func (dto *GetClaimDTO) Into() didrootdomain.GetClaimInput {
	return didrootdomain.GetClaimInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
		ClaimID: utils.HexToHash(dto.ClaimID),
	}
}

// GetClaimsByDIDDTO represents the request payload for retrieving all claims for a DID
// @Description Data Transfer Object for getting all verifiable claims associated with a DID
type GetClaimsByDIDDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetClaimsByDIDDTO) Into() didrootdomain.GetClaimsByDIDInput {
	return didrootdomain.GetClaimsByDIDInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// VerifyClaimDTO represents the request payload for verifying a claim
// @Description Data Transfer Object for verifying if an address has authority over a specific claim
type VerifyClaimDTO struct {
	DIDHash  string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	ClaimID  string `query:"claim_hash_id" validate:"required,hex32" example:"0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"`
	ToVerify string `query:"address_to_verify" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

func (dto *VerifyClaimDTO) Into() didrootdomain.VerifyClaimInput {
	return didrootdomain.VerifyClaimInput{
		DIDHash:  utils.HexToHash(dto.DIDHash),
		ClaimID:  utils.HexToHash(dto.ClaimID),
		ToVerify: utils.HexToAddress(dto.ToVerify),
	}
}

// GetDIDKeyDataCountDTO represents the request payload for getting key count
// @Description Data Transfer Object for retrieving the total number of keys for a DID
type GetDIDKeyDataCountDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetDIDKeyDataCountDTO) Into() didrootdomain.GetDIDKeyDataCountInput {
	return didrootdomain.GetDIDKeyDataCountInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// GetDIDKeyDataByIndexDTO represents the request payload for retrieving a key by index
// @Description Data Transfer Object for getting a specific key using its index position
type GetDIDKeyDataByIndexDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Index   string `query:"index" validate:"required,uint64" example:"0"`
}

func (dto *GetDIDKeyDataByIndexDTO) Into() didrootdomain.GetDIDKeyDataByIndexInput {
	index, _ := strconv.ParseUint(dto.Index, 10, 64)
	return didrootdomain.GetDIDKeyDataByIndexInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
		Index:   index,
	}
}

// GetOriginalKeyDTO represents the request payload for retrieving original key data
// @Description Data Transfer Object for getting the original key using a key code identifier
type GetOriginalKeyDTO struct {
	KeyCode string `query:"key_code" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetOriginalKeyDTO) Into() didrootdomain.GetOriginalKeyInput {
	return didrootdomain.GetOriginalKeyInput{
		KeyCode: utils.HexToHash(dto.KeyCode),
	}
}

// DIDIndexMapDTO represents the request payload for mapping DID index to hash
// @Description Data Transfer Object for retrieving a DID hash using its index
type DIDIndexMapDTO struct {
	DIDIndex string `query:"did_index" validate:"required" example:"1"`
}

func (dto *DIDIndexMapDTO) Into() didrootdomain.DIDIndexMapInput {
	return didrootdomain.DIDIndexMapInput{
		DIDIndex: utils.StringToBigInt(dto.DIDIndex),
	}
}

// DIDIndexMapReverseDTO represents the request payload for reverse mapping DID hash to index
// @Description Data Transfer Object for retrieving a DID index using its hash
type DIDIndexMapReverseDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *DIDIndexMapReverseDTO) Into() didrootdomain.DIDIndexMapReverseInput {
	return didrootdomain.DIDIndexMapReverseInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}
