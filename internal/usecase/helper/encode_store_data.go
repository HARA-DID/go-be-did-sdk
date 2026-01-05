package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeStoreDataParam(input helperdo.EncodeStoreDataParamInput) (string, error) {
	return huc.repo.EncodeStoreDataParam(input)
}
