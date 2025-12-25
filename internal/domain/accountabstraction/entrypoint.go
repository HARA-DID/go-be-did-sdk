package accountabstractiondomain

import (
	"math/big"

	EntryPointFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type HandleOpsInput struct {
	PrivKey string        `json:"PrivKey"`
	Data    string        `json:"Data"`
	Target  utils.Address `json:"Target"`
	Nonce   big.Int       `json:"Nonce"`
	Wallet  utils.Address `json:"Wallet"`
}

type HandleOpsParams = EntryPointFactorySDK.HandleOpsParams
type UserOp = EntryPointFactorySDK.UserOp
