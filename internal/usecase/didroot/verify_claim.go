package didrootusecase

import (
	"context"

	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) VerifyClaim(ctx context.Context, input didrootdomain.VerifyClaimInput) (bool, error) {
	result, err := uc.repo.VerifyClaim(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to verify claim: %v", err)
		return false, err
	}

	return result, nil
}
