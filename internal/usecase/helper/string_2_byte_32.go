package helperusecase

	import helperdo "github.com/meQlause/go-be-did/internal/domain/helper"


func (huc *HelperUseCase) StringToByte32(str helperdo.StringToByte32Input) [32]byte {
	return huc.repo.StringToByte32(str)
}
