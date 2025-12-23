package repository

import (
	"context"

	aa "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
)

type AccountAbstractionRepository interface {
	CreateAccount(ctx context.Context, input aa.CreateAccountInput) (*aa.TxHash, error)
	ExecuteOperation(ctx context.Context, input aa.ExecuteOperationInput) (*aa.TxHash, error)
}
