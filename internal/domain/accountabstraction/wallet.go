package accountabstractiondomain

import (
	WalletSDK "github.com/meQlause/account-abstraction-sdk/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

// ValidateUserOpsInput represents the input for validating user operations
// @Description Validation payload with wallet address and UserOp object
// @Model ValidateUserOpsInput
type ValidateUserOpsInput struct {
	Wallet utils.Address    `json:"wallet" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Input  WalletSDK.UserOp `json:"input"`
}

type Wallet struct {
	Address string
	Owner   string
	Nonce   int64
}

type WalletInfo struct {
	Address string         `json:"address"`
	Owner   string         `json:"owner"`
	Nonce   int64          `json:"nonce"`
	Balance string         `json:"balance"`
	Data    map[string]any `json:"data"`
}
