package didvcdto

import (
	"math/big"
	"strconv"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// GetMetadataDTO represents the request payload for retrieving credential metadata
// @Description Data Transfer Object for getting credential metadata by token ID
type GetMetadataDTO struct {
	TokenID string `query:"token_id" validate:"required,hex_bigint" example:"0x1"`
	Options string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *GetMetadataDTO) Into() didvcdomain.GetMetadataInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)
	return didvcdomain.GetMetadataInput{
		TokenID: utils.StringToBigInt(dto.TokenID),
		Options: uint8(options),
	}
}

// IsCredentialValidDTO represents the request payload for checking credential validity
// @Description Data Transfer Object for verifying if a credential is valid and not expired
type IsCredentialValidDTO struct {
	TokenID string `query:"token_id" validate:"required,hex_bigint" example:"0x1"`
	Options string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *IsCredentialValidDTO) Into() didvcdomain.IsCredentialValidInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)
	return didvcdomain.IsCredentialValidInput{
		TokenID: utils.StringToBigInt(dto.TokenID),
		Options: uint8(options),
	}
}

// GetCredentialsWithMetadataDTO represents the request payload for batch retrieving credentials
// @Description Data Transfer Object for getting multiple credentials with their metadata
type GetCredentialsWithMetadataDTO struct {
	TokenIDs string `query:"token_ids" validate:"required" example:"1,2,3"`
	Options  string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *GetCredentialsWithMetadataDTO) Into() didvcdomain.GetCredentialsWithMetadataInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)

	// Parse comma-separated token IDs
	tokenIDStrs := utils.SplitAndTrim(dto.TokenIDs, ",")
	tokenIDs := make([]*big.Int, 0, len(tokenIDStrs))
	for _, idStr := range tokenIDStrs {
		tokenIDs = append(tokenIDs, utils.StringToBigInt(idStr))
	}

	return didvcdomain.GetCredentialsWithMetadataInput{
		TokenIDs: tokenIDs,
		Options:  uint8(options),
	}
}

// GetUnclaimedTokenIdDTO represents the request payload for getting unclaimed token DID
// @Description Data Transfer Object for retrieving the DID hash of an unclaimed credential
type GetUnclaimedTokenIdDTO struct {
	TokenID string `query:"token_id" validate:"required,hex_bigint" example:"0x1"`
	Options string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *GetUnclaimedTokenIdDTO) Into() didvcdomain.GetUnclaimedTokenIdInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)
	return didvcdomain.GetUnclaimedTokenIdInput{
		TokenID: utils.StringToBigInt(dto.TokenID),
		Options: uint8(options),
	}
}

// CredentialExistsDTO represents the request payload for checking credential existence
// @Description Data Transfer Object for verifying if a credential exists
type CredentialExistsDTO struct {
	TokenID string `query:"token_id" validate:"required,hex_bigint" example:"0x1"`
	Options string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *CredentialExistsDTO) Into() didvcdomain.CredentialExistsInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)
	return didvcdomain.CredentialExistsInput{
		TokenID: utils.StringToBigInt(dto.TokenID),
		Options: uint8(options),
	}
}

// TotalTokensToBeClaimedByDidDTO represents the request payload for getting unclaimed count
// @Description Data Transfer Object for retrieving total unclaimed credentials for a DID
type TotalTokensToBeClaimedByDidDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Options string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *TotalTokensToBeClaimedByDidDTO) Into() didvcdomain.TotalTokensToBeClaimedByDidInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)
	return didvcdomain.TotalTokensToBeClaimedByDidInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
		Options: uint8(options),
	}
}

// GetToBeClaimedTokensByDidDTO represents the request payload for getting paginated unclaimed credentials
// @Description Data Transfer Object for retrieving a list of unclaimed credentials with pagination
type GetToBeClaimedTokensByDidDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Offset  string `query:"offset" validate:"required" example:"0"`
	Limit   string `query:"limit" validate:"required" example:"10"`
	Options string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *GetToBeClaimedTokensByDidDTO) Into() didvcdomain.GetToBeClaimedTokensByDidInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)
	return didvcdomain.GetToBeClaimedTokensByDidInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
		Offset:  utils.StringToBigInt(dto.Offset),
		Limit:   utils.StringToBigInt(dto.Limit),
		Options: uint8(options),
	}
}

// IsApprovedForAllDTO represents the request payload for checking operator approval
// @Description Data Transfer Object for verifying if an operator is approved for all credentials
type IsApprovedForAllDTO struct {
	Owner    string `query:"owner" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Operator string `query:"operator" validate:"required,eth_address" example:"0x1234567890abcdef1234567890abcdef12345678"`
	Options  string `query:"options" validate:"required,oneof=0 1" example:"0"`
}

func (dto *IsApprovedForAllDTO) Into() didvcdomain.IsApprovedForAllInput {
	options, _ := strconv.ParseUint(dto.Options, 10, 8)
	return didvcdomain.IsApprovedForAllInput{
		Owner:    utils.HexToAddress(dto.Owner),
		Operator: utils.HexToAddress(dto.Operator),
		Options:  uint8(options),
	}
}
