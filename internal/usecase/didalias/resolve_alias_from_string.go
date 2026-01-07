package didaliasusecase

import (
	"context"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (uc *DIDAliasUseCase) ResolveFromString(ctx context.Context, input didaliasdomain.ResolveFromStringInput) (utils.Hash, error) {
	result, err := uc.repo.ResolveFromString(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to resolve name '%s': %v", input.Name, err)
		return utils.Hash{}, err
	}
	return result, nil
}
