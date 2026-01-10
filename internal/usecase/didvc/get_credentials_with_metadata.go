package didvccase

import (
	"context"

	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/nftbase"
	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetCredentialsWithMetadata(ctx context.Context, input didvcdomain.GetCredentialsWithMetadataInput) (*nftbase.CredentialsWithMetadataResult, error) {
	result, err := uc.repo.GetCredentialsWithMetadata(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get credentials with metadata: %v", err)
		return nil, err
	}

	return result, nil
}
