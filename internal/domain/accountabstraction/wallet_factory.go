package accountabstractiondomain

import WalletFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"

type CreateWalletInput struct {
	PrivKey string                              `json:"privKey" example:"0xabc123..."`
	Input   WalletFactorySDK.DeployWalletParams `json:"input"`
}
