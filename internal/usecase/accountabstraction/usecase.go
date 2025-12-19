package accountabstraction

import (
	"github.com/meQlause/go-be-did/internal/repository"
)

type UseCase struct {
	repo repository.AccountAbstractionRepository
}

func New(repo repository.AccountAbstractionRepository) *UseCase {
	return &UseCase{repo: repo}
}
