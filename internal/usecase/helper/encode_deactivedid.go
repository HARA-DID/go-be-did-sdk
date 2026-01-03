package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
)

func (huc *HelperUseCase) EncodeDeactiveDIDParam(input helperdo.EncodeDeactiveDIDParamInput) (string, error) {
	return huc.repo.EncodeDeactiveDIDParam(input)
}
