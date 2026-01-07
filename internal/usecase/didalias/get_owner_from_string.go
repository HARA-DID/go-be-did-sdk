package didaliasusecase

import (
	"context"

	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	"github.com/meQlause/go-be-did/pkg/logger"
)

func (uc *DIDAliasUseCase) GetOwnerFromString(ctx context.Context, input didaliasdomain.GetOwnerFromStringInput) (string, error) {
	result, err := uc.repo.GetOwnerFromString(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get owner for name '%s': %v", input.Name, err)
		return "", err
	}
	return result, nil
}
