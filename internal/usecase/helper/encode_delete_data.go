package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
)

func (huc *HelperUseCase) EncodeDeleteDataParam(input helperdo.EncodeDeleteDataParamInput) (string, error) {
	return huc.repo.EncodeDeleteDataParam(input)
}
