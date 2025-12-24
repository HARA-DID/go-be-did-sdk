package repository

import (
	"context"

	aa "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	helperdomain "github.com/meQlause/go-be-did/internal/domain/helper"
)

type AccountAbstractionRepository interface {
	CreateWallet(ctx context.Context, input aa.CreateWalletInput) (*aa.TxHash, error)
	ExecuteOperation(ctx context.Context, input aa.ExecuteOperationInput) (*aa.TxHash, error)
}

type HelperRepository interface {
	StringToByte32(input helperdomain.StringToByte32Input) [32]byte
}
