package helperusecase

import helperdo "github.com/meQlause/go-be-did/internal/domain/helper"

func (huc *HelperUseCase) StringToHex32(input helperdo.StringToHex32Input) string {
	return huc.repo.StringToHex32(input)
}
