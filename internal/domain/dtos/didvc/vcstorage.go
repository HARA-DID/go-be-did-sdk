package didvcdto

import (
	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// GetIdentityTokenCountDTO represents the request payload for getting identity token count
// @Description Data Transfer Object for retrieving total identity tokens for a DID
type GetIdentityTokenCountDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetIdentityTokenCountDTO) Into() didvcdomain.GetIdentityTokenCountInput {
	return didvcdomain.GetIdentityTokenCountInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// GetCertificateTokenCountDTO represents the request payload for getting certificate token count
// @Description Data Transfer Object for retrieving total certificate tokens for a DID
type GetCertificateTokenCountDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetCertificateTokenCountDTO) Into() didvcdomain.GetCertificateTokenCountInput {
	return didvcdomain.GetCertificateTokenCountInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// GetIdentityTokenIdsDTO represents the request payload for getting identity token IDs with pagination
// @Description Data Transfer Object for retrieving paginated identity token IDs for a DID
type GetIdentityTokenIdsDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Offset  string `query:"offset" validate:"required" example:"0"`
	Limit   string `query:"limit" validate:"required" example:"10"`
}

func (dto *GetIdentityTokenIdsDTO) Into() didvcdomain.GetIdentityTokenIdsInput {
	return didvcdomain.GetIdentityTokenIdsInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
		Offset:  utils.StringToBigInt(dto.Offset),
		Limit:   utils.StringToBigInt(dto.Limit),
	}
}

// GetCertificateTokenIdsDTO represents the request payload for getting certificate token IDs with pagination
// @Description Data Transfer Object for retrieving paginated certificate token IDs for a DID
type GetCertificateTokenIdsDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Offset  string `query:"offset" validate:"required" example:"0"`
	Limit   string `query:"limit" validate:"required" example:"10"`
}

func (dto *GetCertificateTokenIdsDTO) Into() didvcdomain.GetCertificateTokenIdsInput {
	return didvcdomain.GetCertificateTokenIdsInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
		Offset:  utils.StringToBigInt(dto.Offset),
		Limit:   utils.StringToBigInt(dto.Limit),
	}
}

// GetAllIdentityTokenIdsDTO represents the request payload for getting all identity token IDs
// @Description Data Transfer Object for retrieving all identity token IDs for a DID
type GetAllIdentityTokenIdsDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetAllIdentityTokenIdsDTO) Into() didvcdomain.GetAllIdentityTokenIdsInput {
	return didvcdomain.GetAllIdentityTokenIdsInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// GetAllCertificateTokenIdsDTO represents the request payload for getting all certificate token IDs
// @Description Data Transfer Object for retrieving all certificate token IDs for a DID
type GetAllCertificateTokenIdsDTO struct {
	DIDHash string `query:"did_hash" validate:"required,hex32" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
}

func (dto *GetAllCertificateTokenIdsDTO) Into() didvcdomain.GetAllCertificateTokenIdsInput {
	return didvcdomain.GetAllCertificateTokenIdsInput{
		DIDHash: utils.HexToHash(dto.DIDHash),
	}
}

// GetDIDRootStorageDTO represents the request payload for getting DID root storage address
// @Description Data Transfer Object for retrieving the DID root storage contract address
type GetDIDRootStorageDTO struct {
}

func (dto *GetDIDRootStorageDTO) Into() didvcdomain.GetDIDRootStorageInput {
	return didvcdomain.GetDIDRootStorageInput{}
}
