package accountabstractiondomain

import (
	"math/big"

	EntryPointFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// HandleOpsInput represents the input for handling user operations
// @Description HandleOps request payload with private key, wallet address, target address, data, and nonce
// @Model HandleOpsInput
type HandleOpsInput struct {
	PrivKey string        `json:"priv_key" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef" binding:"required"`
	Data    string        `json:"data" example:"0x..." binding:"required"`
	Target  utils.Address `json:"target" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb" binding:"required"`
	Nonce   big.Int       `json:"nonce" example:"0x0" binding:"required"`
	Wallet  utils.Address `json:"wallet" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb" binding:"required"`
}

type HandleOpsParams = EntryPointFactorySDK.HandleOpsParams
type UserOp = EntryPointFactorySDK.UserOp
