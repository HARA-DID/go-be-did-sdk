package accountabstractionusecase

import (
	"context"

	aado "github.com/meQlause/go-be-did/internal/domain/entities/accountabstraction"
	"github.com/meQlause/go-be-did/pkg/logger"
	backendutils "github.com/meQlause/go-be-did/utils"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
)

func (uc *AccountAbstactionUseCase) HandleOps(ctx context.Context, input aado.HandleOpsInput, net *network.Network) (*backendutils.TxHash, error) {
	result, err := uc.repo.HandleOps(ctx, input, net)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to execute operation: %v", err)
		return nil, err
	}

	return result, nil
}
