package repository

import (
	"context"

	aa "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
)

type AccountAbstractionRepository interface {
	CreateWallet(ctx context.Context, input aa.CreateWalletInput) (*aa.TxHash, error)
	HandleOps(ctx context.Context, input aa.HandleOpsInput, net *network.Network) (*aa.TxHash, error)
}

type HelperRepository interface {
	StringToByte32(input helperdo.StringToByte32Input) [32]byte
	EncodeCreateDIDParam(createDIDParam helperdo.EncodeCreateDIDParamInput) (string, error)
}
