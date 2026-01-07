package accountabstractionusecase

import (
	"context"

	aado "github.com/meQlause/go-be-did/internal/domain/entities/accountabstraction"

	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *AccountAbstactionUseCase) IsValidWallet(ctx context.Context, input aado.IsValidWalletInput) (bool, error) {
	result, err := uc.repo.IsValidWallet(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to create account: %v", err)
		return false, err
	}

	return result, nil
}
