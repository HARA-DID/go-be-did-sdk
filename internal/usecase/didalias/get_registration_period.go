package didaliasusecase

import (
	"context"
	"math/big"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *DIDAliasUseCase) GetRegistrationPeriod(ctx context.Context, input didaliasdomain.GetRegistrationPeriodInput) (*big.Int, error) {
	result, err := uc.repo.GetRegistrationPeriod(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get registration period: %v", err)
		return nil, err
	}
	return result, nil
}
