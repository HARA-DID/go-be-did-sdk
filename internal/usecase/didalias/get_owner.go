package didaliasusecase

import (
	"context"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *DIDAliasUseCase) GetOwner(ctx context.Context, input didaliasdomain.GetOwnerInput) (string, error) {
	result, err := uc.repo.GetOwner(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get owner: %v", err)
		return "", err
	}
	return result, nil
}
