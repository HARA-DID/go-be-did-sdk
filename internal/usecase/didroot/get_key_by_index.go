package didrootusecase

import (
	"context"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (uc *DIDRootUseCase) GetDIDKeyDataByIndex(ctx context.Context, input didrootdomain.GetDIDKeyDataByIndexInput) (utils.Hash, error) {
	result, err := uc.repo.GetDIDKeyDataByIndex(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get DID key by index: %v", err)
		return utils.Hash{}, err
	}

	return result, nil
}
