package didvccase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) IsCredentialValid(ctx context.Context, input didvcdomain.IsCredentialValidInput) (bool, error) {
	result, err := uc.repo.IsCredentialValid(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to check credential validity: %v", err)
		return false, err
	}

	return result, nil
}
