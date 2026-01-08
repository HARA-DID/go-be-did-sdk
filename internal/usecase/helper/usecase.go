package helperusecase

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"
	"github.com/meQlause/go-be-did/internal/repository"
)

type HelperUseCase struct {
	repo repository.HelperRepository
}

func New(repo repository.HelperRepository) *HelperUseCase {
	return &HelperUseCase{repo: repo}
}

func (huc *HelperUseCase) StringToHex32(input helperdo.StringToHex32Input) string {
	return huc.repo.StringToHex32(input)
}
