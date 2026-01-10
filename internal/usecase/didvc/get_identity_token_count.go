package didvccase

import (
	"context"
	"math/big"

	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetIdentityTokenCount(ctx context.Context, input didvcdomain.GetIdentityTokenCountInput) (*big.Int, error) {
	result, err := uc.repo.GetIdentityTokenCount(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get identity token count: %v", err)
		return big.NewInt(0), err
	}

	return result, nil
}
