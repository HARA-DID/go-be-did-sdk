package dto

import (
	accountabstraction "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	backendutils "github.com/meQlause/go-be-did/utils"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// HandleOpsInputDTO represents the DTO for HandleOps request (used for Swagger documentation and request parsing)
// @Description HandleOps request payload with private key, wallet address, target address, data, and nonce
type HandleOpsDTO struct {
	Details backendutils.Details `json:"details" validate:"required"`
	PrivKey string               `json:"priv_key" validate:"required,eth_private_key" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Data    string               `json:"data" validate:"required,hex_data" example:"0x..."`
	Target  string               `json:"target" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Nonce   string               `json:"nonce" validate:"required,numeric" example:"0"`
	Wallet  string               `json:"wallet" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

func (dto *HandleOpsDTO) Into() accountabstraction.HandleOpsInput {
	return accountabstraction.HandleOpsInput{
		PrivKey: dto.PrivKey,
		Data:    dto.Data,
		Target:  utils.HexToAddress(dto.Target),
		Nonce:   utils.StringToBigInt(dto.Nonce),
		Wallet:  utils.HexToAddress(dto.Wallet),
	}
}
