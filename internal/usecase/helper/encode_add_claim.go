package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
)

func (huc *HelperUseCase) EncodeAddClaimParam(input helperdo.EncodeAddClaimParamInput) (string, error) {
	return huc.repo.EncodeAddClaimParam(input)
}
