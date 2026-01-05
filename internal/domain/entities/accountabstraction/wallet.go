package accountabstractiondomain

import (
	WalletSDK "github.com/meQlause/account-abstraction-sdk/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type ValidateUserOpsInput struct {
	Wallet utils.Address
	Input  WalletSDK.UserOp
}
