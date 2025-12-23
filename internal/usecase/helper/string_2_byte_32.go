package helperusecase

import "github.com/meQlause/hara-core-blockchain-lib/utils"

func (huc *HelperUseCase) StringToByte32(str string) [32]byte {
	return utils.StringToByte32(str)
}
