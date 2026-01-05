package accountabstractiondto

import (
	"strings"

	accountabstractiondomain "github.com/meQlause/go-be-did/internal/domain/entities/accountabstraction"

	"github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// CreateWalletInputDTO represents the DTO for CreateWallet request (used for Swagger documentation and request parsing)
// @Description Wallet creation payload with deployer address and optional salt value
type CreateWalletInputDTO struct {
	PrivKey string                `json:"priv_key" validate:"required,eth_private_key" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Input   DeployWalletParamsDTO `json:"input" validate:"required"`
}

// DeployWalletParamsDTO represents the DTO for DeployWalletParams (used for Swagger documentation and request parsing)
// @Description Wallet deployment parameters with owners and salt
type DeployWalletParamsDTO struct {
	Owners []string `json:"owners" validate:"required,min=1,dive,eth_address" example:"0x111...,0x222..."`
	Salt   string   `json:"salt" validate:"required,hex32" example:"0xabc123... (32 bytes hex)"`
}

func (dto *CreateWalletInputDTO) Into() accountabstractiondomain.CreateWalletInput {
	owners := make([]utils.Address, len(dto.Input.Owners))
	for i, o := range dto.Input.Owners {
		owners[i] = utils.HexToAddress(o)
	}

	var salt [32]byte
	saltBytes, _ := utils.DecodeString(strings.TrimPrefix(dto.Input.Salt, "0x"))
	copy(salt[:], saltBytes)

	return accountabstractiondomain.CreateWalletInput{
		PrivKey: dto.PrivKey,
		Input: walletfactory.DeployWalletParams{
			Owners: owners,
			Salt:   salt,
		},
	}
}
