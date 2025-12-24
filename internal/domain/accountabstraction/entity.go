package accountabstractiondomain

import (
	EntryPointSDK "github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"
	WalletFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"
)

type Wallet struct {
	Address string
	Owner   string
	Nonce   int64
}

type CreateWalletInput struct {
	PrivKey string                              `json:"privKey"`
	Input   WalletFactorySDK.DeployWalletParams `json:"input"`
}

type TxHash struct {
	TxHash []string
}

type ExecuteOperationInput struct {
	PrivKey string
	Input   EntryPointSDK.HandleOpsParams
}

type ExecuteOperationOutput struct {
	TxHash string
	Status string
}

type WalletInfo struct {
	Address string         `json:"address"`
	Owner   string         `json:"owner"`
	Nonce   int64          `json:"nonce"`
	Balance string         `json:"balance"`
	Data    map[string]any `json:"data"`
}
