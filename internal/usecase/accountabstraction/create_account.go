package accountabstraction

import (
	"context"

	aa "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *UseCase) CreateAccount(ctx context.Context, input aa.CreateAccountInput) (*aa.TxHash, error) {
	result, err := uc.repo.CreateAccount(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to create account: %v", err)
		return nil, err
	}

	return result, nil
}
