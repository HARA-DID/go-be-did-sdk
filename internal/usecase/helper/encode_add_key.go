package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
)

func (huc *HelperUseCase) EncodeAddKeyParam(input helperdo.EncodeAddKeyParamInput) (string, error) {
	return huc.repo.EncodeAddKeyParam(input)
}
