package sdk

import (
	"context"
	"errors"
	"fmt"

	entryPointSDK "github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"
	GasManagerSDK "github.com/meQlause/account-abstraction-sdk/pkg/gasmanager"
	WalletSDK "github.com/meQlause/account-abstraction-sdk/pkg/wallet"
	WalletFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"
	"github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
)

type AccountAbstractionHNS struct {
	EntryPoint    string
	GasManager    string
	Wallet        string
	WalletFactory string
}

func (h AccountAbstractionHNS) Validate() error {
	if h.EntryPoint == "" {
		return errors.New("entrypoint HNS is empty")
	}
	if h.GasManager == "" {
		return errors.New("gas manager HNS is empty")
	}
	if h.Wallet == "" {
		return errors.New("wallet HNS is empty")
	}
	if h.WalletFactory == "" {
		return errors.New("wallet factory HNS is empty")
	}
	return nil
}

type AccountAbstractionSDK struct {
	EntryPoint    *entryPointSDK.EntryPoint
	GasManager    *GasManagerSDK.GasManager
	Wallet        *WalletSDK.Wallet
	WalletFactory *WalletFactorySDK.WalletFactory
}

func (s *AccountAbstractionSDK) CreateAccount(
	ctx context.Context,
	input accountabstraction.CreateAccountInput,
) (*accountabstraction.TxHash, error) {
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

	return &accountabstraction.TxHash{
		TxHash: txHashes,
	}, nil
}

func (s *AccountAbstractionSDK) ExecuteOperation(
	ctx context.Context,
	input accountabstraction.ExecuteOperationInput,
) (*accountabstraction.TxHash, error) {
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

	return &accountabstraction.TxHash{
		TxHash: hashes,
	}, nil
}

func NewAccountAbstractionSDK(
	ctx context.Context,
	hns AccountAbstractionHNS,
	bc *blockchain.Blockchain,
) (repository.AccountAbstractionRepository, error) {

	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if bc == nil {
		return nil, errors.New("blockchain cannot be nil")
	}
	if err := hns.Validate(); err != nil {
		return nil, fmt.Errorf("invalid HNS config: %w", err)
	}

	entryPoint, err := entryPointSDK.NewEntryPointWithHNS(ctx, hns.EntryPoint, bc)
	if err != nil {
		return nil, fmt.Errorf("init entry point: %w", err)
	}

	gasManager, err := GasManagerSDK.NewGasManagerWithHNS(ctx, hns.GasManager, bc)
	if err != nil {
		return nil, fmt.Errorf("init gas manager: %w", err)
	}

	wallet, err := WalletSDK.NewWalletWithHNS(ctx, hns.Wallet, bc)
	if err != nil {
		return nil, fmt.Errorf("init wallet: %w", err)
	}

	walletFactory, err := WalletFactorySDK.NewWalletFactoryWithHNS(ctx, hns.WalletFactory, bc)
	if err != nil {
		return nil, fmt.Errorf("init wallet factory: %w", err)
	}

	return &AccountAbstractionSDK{
		EntryPoint:    entryPoint,
		GasManager:    gasManager,
		Wallet:        wallet,
		WalletFactory: walletFactory,
	}, nil
}
