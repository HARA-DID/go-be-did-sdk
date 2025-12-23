package accountabstraction

import (
	EntryPointSDK "github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"
	WalletFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"
)

type Account struct {
	Address string
	Owner   string
	Nonce   int64
}

type CreateAccountInput struct {
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

type AccountInfo struct {
	Address string                 `json:"address"`
	Owner   string                 `json:"owner"`
	Nonce   int64                  `json:"nonce"`
	Balance string                 `json:"balance"`
	Data    map[string]interface{} `json:"data"`
}
