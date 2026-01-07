package didaliasusecase

import (
	"context"

	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *DIDAliasUseCase) GetAliasStatusFromString(ctx context.Context, input didaliasdomain.GetAliasStatusFromStringInput) (*aliasfactory.AliasStatus, error) {
	result, err := uc.repo.GetAliasStatusFromString(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get alias status for name '%s': %v", input.Name, err)
		return nil, err
	}
	return result, nil
}
