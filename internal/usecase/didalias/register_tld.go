package didaliasusecase

import (
	"context"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	backendutils "github.com/meQlause/go-be-did/utils"

	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *DIDAliasUseCase) RegisterTLD(ctx context.Context, input didaliasdomain.RegisterTLDInput) (*backendutils.TxHash, error) {
	result, err := uc.repo.RegisterTLD(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to create account: %v", err)
		return nil, err
	}

	return result, nil
}
