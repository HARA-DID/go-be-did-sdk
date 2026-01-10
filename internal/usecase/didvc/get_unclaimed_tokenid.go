package didvccase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetUnclaimedTokenId(ctx context.Context, input didvcdomain.GetUnclaimedTokenIdInput) (utils.Hash, error) {
	result, err := uc.repo.GetUnclaimedTokenId(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get unclaimed token id: %v", err)
		return utils.Hash{}, err
	}

	return result, nil
}
