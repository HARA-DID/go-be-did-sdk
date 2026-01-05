package didrootusecase

import (
	"context"

	"github.com/meQlause/did-root-sdk/pkg/rootstorage"
	"github.com/meQlause/go-be-did/pkg/logger"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

func (uc *DIDRootUseCase) GetKey(ctx context.Context, input didrootdomain.GetKeyInput) (*rootstorage.Key, error) {
	result, err := uc.repo.GetKey(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get key: %v", err)
		return nil, err
	}

	return result, nil
}
