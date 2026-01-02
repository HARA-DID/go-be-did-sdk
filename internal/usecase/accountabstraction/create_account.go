package accountabstractionusecase

import (
	"context"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"

	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *AccountAbstactionUseCase) CreateWallet(ctx context.Context, input aado.CreateWalletInput) (*aado.TxHash, error) {
	result, err := uc.repo.CreateWallet(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to create account: %v", err)
		return nil, err
	}

	return result, nil
}
