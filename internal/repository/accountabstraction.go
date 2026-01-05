package repository

import (
	"context"
	
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
	
	aado "github.com/meQlause/go-be-did/internal/domain/entities/accountabstraction"
	backendutils "github.com/meQlause/go-be-did/utils"
)

type AccountAbstractionRepository interface {
	CreateWallet(ctx context.Context, input aado.CreateWalletInput) (*backendutils.TxHash, error)
	HandleOps(ctx context.Context, input aado.HandleOpsInput, net *network.Network) (*backendutils.TxHash, error)
}
