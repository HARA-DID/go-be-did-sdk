package didvccase

import (
	"context"
	"math/big"

	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetAllCertificateTokenIds(ctx context.Context, input didvcdomain.GetAllCertificateTokenIdsInput) ([]*big.Int, error) {
	result, err := uc.repo.GetAllCertificateTokenIds(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get all certificate token ids: %v", err)
		return nil, err
	}

	return result, nil
}
