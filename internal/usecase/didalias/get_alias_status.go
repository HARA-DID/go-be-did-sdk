package didaliasusecase

import (
	"context"

	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *DIDAliasUseCase) GetAliasStatus(ctx context.Context, input didaliasdomain.GetAliasStatusInput) (*aliasfactory.AliasStatus, error) {
	result, err := uc.repo.GetAliasStatus(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get alias status: %v", err)
		return nil, err
	}
	return result, nil
}
