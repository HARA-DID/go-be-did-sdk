package didrootusecase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) VerifyDIDOwnership(ctx context.Context, input didrootdomain.VerifyDIDOwnershipInput) (bool, error) {
	result, err := uc.repo.VerifyDIDOwnership(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to verify DID ownership: %v", err)
		return false, err
	}

	return result, nil
}
