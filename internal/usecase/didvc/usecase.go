package didvccase

import (
	"context"

	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/go-be-did/pkg/logger"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

type DIDVCUseCase struct {
	repo repository.DIDVCRepository
}

func New(repo repository.DIDVCRepository) *DIDVCUseCase {
	return &DIDVCUseCase{repo: repo}
}

func (uc *DIDVCUseCase) GetDIDRootStorage(ctx context.Context, input didvcdomain.GetDIDRootStorageInput) (utils.Address, error) {
	result, err := uc.repo.GetDIDRootStorage(ctx, input)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get DID root storage: %v", err)
		return utils.Address{}, err
	}

	return result, nil
}
