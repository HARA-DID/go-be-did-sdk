package didvccase

import (
	"context"
	"math/big"

	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetCertificateTokenCount(ctx context.Context, input didvcdomain.GetCertificateTokenCountInput) (*big.Int, error) {
	result, err := uc.repo.GetCertificateTokenCount(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get certificate token count: %v", err)
		return big.NewInt(0), err
	}

	return result, nil
}
