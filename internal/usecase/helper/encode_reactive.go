package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
)

func (huc *HelperUseCase) EncodeReactiveDIDParam(input helperdo.EncodeReactiveDIDParamInput) (string, error) {
	return huc.repo.EncodeReactiveDIDParam(input)
}
