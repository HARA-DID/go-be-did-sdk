package accountabstractionsdk

import (
	"context"
	"fmt"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
)

func (s *AccountAbstractionSDK) ExecuteOperation(
	ctx context.Context,
	input aado.ExecuteOperationInput,
) (*aado.TxHash, error) {
	wallet := wallet.NewWallet(input.PrivKey)

	hashes, err := s.EntryPoint.HandleOps(
		ctx,
		wallet,
		input.Input,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("handle ops failed: %w", err)
	}

	return &aado.TxHash{
		TxHash: hashes,
	}, nil
}
