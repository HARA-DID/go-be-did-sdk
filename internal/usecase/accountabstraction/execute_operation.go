package accountabstraction

import (
	"context"

	"github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *UseCase) ExecuteOperation(ctx context.Context, input *accountabstraction.ExecuteOperationInput) (*accountabstraction.TxHash, error) {
	result, err := uc.repo.ExecuteOperation(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to execute operation: %v", err)
		return nil, err
	}

	return result, nil
}
