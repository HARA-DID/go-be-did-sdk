package didrootusecase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) GetKeysByDID(ctx context.Context, input didrootdomain.GetKeysByDIDInput) ([]utils.Hash, error) {
	result, err := uc.repo.GetKeysByDID(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get keys by DID: %v", err)
		return nil, err
	}

	return result, nil
}
