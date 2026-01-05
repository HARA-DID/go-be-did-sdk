package accountabstractiondomain

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type HandleOpsInput struct {
	PrivKey string
	Data    string
	Target  utils.Address
	Nonce   *big.Int
	Wallet  utils.Address
}

type IsValidWalletInput struct {
	Wallet utils.Address
}
