package didrootusecase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) GetOriginalKey(ctx context.Context, input didrootdomain.GetOriginalKeyInput) (string, error) {
	result, err := uc.repo.GetOriginalKey(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get original key: %v", err)
		return "", err
	}

	return result, nil
}
