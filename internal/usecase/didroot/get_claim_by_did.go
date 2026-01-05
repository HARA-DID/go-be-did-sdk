package didrootusecase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) GetClaimsByDID(ctx context.Context, input didrootdomain.GetClaimsByDIDInput) ([]utils.Hash, error) {
	result, err := uc.repo.GetClaimsByDID(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get claims by DID: %v", err)
		return nil, err
	}

	return result, nil
}
