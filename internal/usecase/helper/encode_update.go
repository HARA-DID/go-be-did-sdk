package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeUpdateDIDParam(input helperdo.EncodeUpdateDIDParamInput) (string, error) {
	return huc.repo.EncodeUpdateDIDParam(input)
}
