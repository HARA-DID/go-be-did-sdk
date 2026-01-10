package didvccase

import (
	"context"

	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/vcstorage"
	"github.com/meQlause/go-be-did/pkg/logger"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (uc *DIDVCUseCase) GetCertificateTokenIds(ctx context.Context, input didvcdomain.GetCertificateTokenIdsInput) (*vcstorage.TokenIdsResult, error) {
	result, err := uc.repo.GetCertificateTokenIds(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get certificate token ids: %v", err)
		return nil, err
	}

	return result, nil
}
