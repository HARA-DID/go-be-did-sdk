package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeRemoveKeyParam(input helperdo.EncodeRemoveKeyParamInput) (string, error) {
	return huc.repo.EncodeRemoveKeyParam(input)
}
