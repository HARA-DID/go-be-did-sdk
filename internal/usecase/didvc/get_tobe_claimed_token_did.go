package didvccase

import (
	"context"
	"math/big"

	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetToBeClaimedTokensByDid(ctx context.Context, input didvcdomain.GetToBeClaimedTokensByDidInput) ([]*big.Int, error) {
	result, err := uc.repo.GetToBeClaimedTokensByDid(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get tokens to be claimed by DID: %v", err)
		return nil, err
	}

	return result, nil
}
