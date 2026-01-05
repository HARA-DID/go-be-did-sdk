package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeRemoveClaimParam(input helperdo.EncodeRemoveClaimParamInput) (string, error) {
	return huc.repo.EncodeRemoveClaimParam(input)
}
