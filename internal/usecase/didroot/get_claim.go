package didrootusecase

import (
	"context"

	"github.com/meQlause/did-root-sdk/pkg/rootstorage"
	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) GetClaim(ctx context.Context, input didrootdomain.GetClaimInput) (*rootstorage.Claim, error) {
	result, err := uc.repo.GetClaim(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get claim: %v", err)
		return nil, err
	}

	return result, nil
}
