package dto

import (
	WalletSDK "github.com/meQlause/account-abstraction-sdk/pkg/wallet"
	accountabstractiondomain "github.com/meQlause/go-be-did/internal/domain/accountabstraction"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// ValidateUserOpsInputDTO represents the DTO for ValidateUserOps request (used for Swagger documentation and request parsing)
// @Description Validation payload with wallet address and UserOp object
type ValidateUserOpsDTO struct {
	Wallet string    `json:"wallet" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Input  UserOpDTO `json:"input" validate:"required,dive"`
}

// UserOpDTO represents the DTO for UserOp (used for Swagger documentation and request parsing)
// @Description UserOp object containing target address, value, data, client block number, user nonce, and signature
// ValidateUserOpsInputDTO represents the DTO for ValidateUserOps request
type UserOpDTO struct {
	Target            string `json:"target" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Value             string `json:"value" validate:"required,numeric" example:"0"`
	Data              string `json:"data" validate:"required,hex_data" example:"0x..."`
	ClientBlockNumber string `json:"client_block_number" validate:"required,numeric" example:"12345"`
	UserNonce         string `json:"user_nonce" validate:"required,numeric" example:"0"`
	Signature         string `json:"signature" validate:"required,hex_data" example:"0x..."`
}

func (dto *ValidateUserOpsDTO) Into() accountabstractiondomain.ValidateUserOpsInput {
	return accountabstractiondomain.ValidateUserOpsInput{
		Wallet: utils.HexToAddress(dto.Wallet),
		Input: WalletSDK.UserOp{
			Target:            utils.HexToAddress(dto.Input.Target),
			Value:             utils.StringToBigInt(dto.Input.Value),
			Data:              utils.Hex2Bytes(dto.Input.Data),
			ClientBlockNumber: utils.StringToBigInt(dto.Input.ClientBlockNumber),
			UserNonce:         utils.StringToBigInt(dto.Input.UserNonce),
			Signature:         utils.Hex2Bytes(dto.Input.Signature),
		},
	}
}
