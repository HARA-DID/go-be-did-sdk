package accountabstractionsdk

import (
	"context"
	"fmt"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
)

func (s *AccountAbstractionSDK) CreateAccount(
	ctx context.Context,
	input aado.CreateWalletInput,
) (*aado.TxHash, error) {
	wallet := wallet.NewWallet(input.PrivKey)

	txHashes, err := s.WalletFactory.DeployWallet(
		ctx,
		wallet,
		input.Input,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("create wallet failed: %w", err)
	}

	return &aado.TxHash{
		TxHash: txHashes,
	}, nil
}
