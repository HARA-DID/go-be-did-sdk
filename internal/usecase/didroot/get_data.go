package didrootusecase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) GetData(ctx context.Context, input didrootdomain.GetDataInput) (string, error) {
	result, err := uc.repo.GetData(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get data: %v", err)
		return "", err
	}

	return result, nil
}
