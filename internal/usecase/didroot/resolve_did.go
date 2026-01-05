package didrootusecase

import (
	"context"

	"github.com/meQlause/did-root-sdk/pkg/rootstorage"
	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) ResolveDID(ctx context.Context, input didrootdomain.ResolveDIDInput) (*rootstorage.DIDDocument, error) {
	result, err := uc.repo.ResolveDID(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to resolve DID: %v", err)
		return nil, err
	}

	return result, nil
}
