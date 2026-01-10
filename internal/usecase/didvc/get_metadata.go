package didvccase

import (
	"context"

	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/nftbase"
	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetMetadata(ctx context.Context, input didvcdomain.GetMetadataInput) (*nftbase.CredentialMetadata, error) {
	result, err := uc.repo.GetMetadata(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get metadata: %v", err)
		return nil, err
	}

	return result, nil
}
