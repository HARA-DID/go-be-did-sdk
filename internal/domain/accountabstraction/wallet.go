package accountabstractiondomain

import (
	WalletSDK "github.com/meQlause/account-abstraction-sdk/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type ValidateUserOpsInput struct {
	Wallet utils.Address
	Input  WalletSDK.UserOp
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
