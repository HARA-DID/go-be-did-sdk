package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeCreateDIDParam(input helperdo.EncodeCreateDIDParamInput) (string, error) {
	return huc.repo.EncodeCreateDIDParam(input)
}
