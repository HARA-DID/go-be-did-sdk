package helperusecase

import (
	"github.com/meQlause/go-be-did/internal/repository"
)

type HelperUseCase struct {
	repo repository.HelperRepository
}

func New(repo repository.HelperRepository) *HelperUseCase {
	return &HelperUseCase{repo: repo}
}
