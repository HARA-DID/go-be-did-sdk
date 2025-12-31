package accountabstractiondomain

import WalletFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"

// CreateWalletInput represents the input for creating a wallet
// @Description Wallet creation payload with deployer address and optional salt value
// @Model CreateWalletInput
type CreateWalletInput struct {
	PrivKey string                              `json:"priv_key" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Input   WalletFactorySDK.DeployWalletParams `json:"input"`
}
