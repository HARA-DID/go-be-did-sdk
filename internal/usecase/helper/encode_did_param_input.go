package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
)

func (huc *HelperUseCase) EncodeCreateDIDParam(createDIDParam helperdo.EncodeCreateDIDParamInput) (string, error) {
	return huc.repo.EncodeCreateDIDParam(createDIDParam)
}
