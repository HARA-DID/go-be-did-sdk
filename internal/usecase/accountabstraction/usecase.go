package accountabstractionusecase

import (
	"github.com/meQlause/go-be-did/internal/repository"
)

type AccountAbstactionUseCase struct {
	repo repository.AccountAbstractionRepository
}

func New(repo repository.AccountAbstractionRepository) *AccountAbstactionUseCase {
	return &AccountAbstactionUseCase{repo: repo}
}
