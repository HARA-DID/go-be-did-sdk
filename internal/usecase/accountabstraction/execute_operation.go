package accountabstractionusecase

import (
	"context"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *AccountAbstactionUseCase) ExecuteOperation(ctx context.Context, input aado.ExecuteOperationInput) (*aado.TxHash, error) {
	result, err := uc.repo.ExecuteOperation(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to execute operation: %v", err)
		return nil, err
	}

	return result, nil
}
