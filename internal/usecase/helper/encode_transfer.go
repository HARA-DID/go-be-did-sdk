package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeTransferDIDOwnerParam(input helperdo.EncodeTransferDIDOwnerParamInput) (string, error) {
	return huc.repo.EncodeTransferDIDOwnerParam(input)
}
