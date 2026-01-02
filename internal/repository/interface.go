package repository

import (
	"context"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
)

type AccountAbstractionRepository interface {
	CreateWallet(ctx context.Context, input aado.CreateWalletInput) (*aado.TxHash, error)
	HandleOps(ctx context.Context, input aado.HandleOpsInput, net *network.Network) (*aado.TxHash, error)
}

type HelperRepository interface {
	StringToHex32(input helperdo.StringToHex32Input) string
	EncodeCreateDIDParam(createDIDParam helperdo.EncodeCreateDIDParamInput) (string, error)
}
