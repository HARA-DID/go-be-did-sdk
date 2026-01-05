package didrootusecase

import (
	"context"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *DIDRootUseCase) GetDIDKeyDataCount(ctx context.Context, input didrootdomain.GetDIDKeyDataCountInput) (uint64, error) {
	result, err := uc.repo.GetDIDKeyDataCount(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get DID key count: %v", err)
		return 0, err
	}

	return result, nil
}
