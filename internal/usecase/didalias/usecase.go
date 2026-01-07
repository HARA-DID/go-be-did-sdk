package didaliasusecase

import (
	"github.com/meQlause/go-be-did/internal/repository"
)

type DIDAliasUseCase struct {
	repo repository.DIDAliasRepository
}

func New(repo repository.DIDAliasRepository) *DIDAliasUseCase {
	return &DIDAliasUseCase{repo: repo}
}
