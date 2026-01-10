package didvccase

import (
	"context"
	"math/big"

	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) TotalTokensToBeClaimedByDid(ctx context.Context, input didvcdomain.TotalTokensToBeClaimedByDidInput) (*big.Int, error) {
	result, err := uc.repo.TotalTokensToBeClaimedByDid(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get total tokens to be claimed: %v", err)
		return big.NewInt(0), err
	}

	return result, nil
}
