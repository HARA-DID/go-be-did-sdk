package didrootusecase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) DIDIndexMap(ctx context.Context, input didrootdomain.DIDIndexMapInput) (string, error) {
	result, err := uc.repo.DIDIndexMap(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get DID index map: %v", err)
		return "", err
	}

	return result, nil
}
