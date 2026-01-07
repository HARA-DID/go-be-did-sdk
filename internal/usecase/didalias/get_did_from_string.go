package didaliasusecase

import (
	"context"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (uc *DIDAliasUseCase) GetDIDFromString(ctx context.Context, input didaliasdomain.GetDIDFromStringInput) (utils.Hash, error) {
	result, err := uc.repo.GetDIDFromString(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get DID for name '%s': %v", input.Name, err)
		return utils.Hash{}, err
	}
	return result, nil
}
