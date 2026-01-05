package didrootusecase

import (
	"context"
	"math/big"

	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) DIDIndexMapReverse(ctx context.Context, input didrootdomain.DIDIndexMapReverseInput) (*big.Int, error) {
	result, err := uc.repo.DIDIndexMapReverse(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get DID index map reverse: %v", err)
		return nil, err
	}

	return result, nil
}
