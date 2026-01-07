package didaliasusecase

import (
	"context"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (uc *DIDAliasUseCase) Resolve(ctx context.Context, input didaliasdomain.ResolveInput) (utils.Hash, error) {
	result, err := uc.repo.Resolve(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to resolve node: %v", err)
		return utils.Hash{}, err
	}
	return result, nil
}
