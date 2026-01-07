package didaliasusecase

import (
	"context"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (uc *DIDAliasUseCase) GetDID(ctx context.Context, input didaliasdomain.GetDIDInput) (utils.Hash, error) {
	result, err := uc.repo.GetDID(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get DID: %v", err)
		return utils.Hash{}, err
	}
	return result, nil
}
