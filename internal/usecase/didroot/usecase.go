package didrootusecase

import (
	"github.com/meQlause/go-be-did/internal/repository"
)

type DIDRootUseCase struct {
	repo repository.DIDRootRepository
}

func New(repo repository.DIDRootRepository) *DIDRootUseCase {
	return &DIDRootUseCase{repo: repo}
}
