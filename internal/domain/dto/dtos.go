package dto

import (
	"strings"

	WalletSDK "github.com/meQlause/account-abstraction-sdk/pkg/wallet"
	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	accountabstractiondomain "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	helperdomain "github.com/meQlause/go-be-did/internal/domain/helper"

	"github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// HandleOpsInputDTO represents the DTO for HandleOps request (used for Swagger documentation and request parsing)
// @Description HandleOps request payload with private key, wallet address, target address, data, and nonce
type HandleOpsDTO struct {
	PrivKey string `json:"priv_key" validate:"required,eth_private_key" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Data    string `json:"data" validate:"required,hex_data" example:"0x..."`
	Target  string `json:"target" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Nonce   string `json:"nonce" validate:"required,numeric" example:"0"`
	Wallet  string `json:"wallet" validate:"required,eth_address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
}

func (dto *HandleOpsDTO) Into() aado.HandleOpsInput {
	return aado.HandleOpsInput{
		PrivKey: dto.PrivKey,
		Data:    dto.Data,
		Target:  utils.HexToAddress(dto.Target),
		Nonce:   utils.StringToBigInt(dto.Nonce),
		Wallet:  utils.HexToAddress(dto.Wallet),
	}
}

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

type StringToHex32DTO struct {
	Input string `json:"input" validate:"required" example:"example_string"`
}

func (dto *StringToHex32DTO) Into() helperdomain.StringToHex32Input {
	return helperdomain.StringToHex32Input{
		Input: dto.Input,
	}
}

type EncodeCreateDIDDTO struct {
	Address       string       `json:"address" validate:"required" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      CreateDIDDTO `json:"did_param"`
	KeyIdentifier string       `json:"key_identifier" example:"key1"`
}

type CreateDIDDTO struct {
	DID string `json:"did" example:"did:example:123456789"`
}

func (dto *EncodeCreateDIDDTO) Into() helperdomain.EncodeCreateDIDParamInput {
	return helperdomain.EncodeCreateDIDParamInput{
		Address: dto.Address,
		DIDParam: rootfactory.CreateDIDParam{
			DID: dto.DIDParam.DID,
		},
		KeyIdentifier: dto.KeyIdentifier,
	}
}
