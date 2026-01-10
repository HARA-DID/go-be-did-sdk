package didvccase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) IsApprovedForAll(ctx context.Context, input didvcdomain.IsApprovedForAllInput) (bool, error) {
	result, err := uc.repo.IsApprovedForAll(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to check approval for all: %v", err)
		return false, err
	}

	return result, nil
}
