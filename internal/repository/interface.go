package repository

import (
	"context"

	"github.com/meQlause/go-be-did/internal/domain/accountabstraction"
)

type AccountAbstractionRepository interface {
	CreateAccount(ctx context.Context, input *accountabstraction.CreateAccountInput) (*accountabstraction.TxHash, error)
	ExecuteOperation(ctx context.Context, input *accountabstraction.ExecuteOperationInput) (*accountabstraction.TxHash, error)
}
