package repository

import (
	"context"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
)

type AccountAbstractionRepository interface {
	CreateWallet(ctx context.Context, input aado.CreateWalletInput) (*aado.TxHash, error)
	HandleOps(ctx context.Context, input aado.HandleOpsInput, net *network.Network) (*aado.TxHash, error)
}
