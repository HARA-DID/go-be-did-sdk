package accountabstractionsdk

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/contract"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	EntryPointSDK "github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"
	GasManagerSDK "github.com/meQlause/account-abstraction-sdk/pkg/gasmanager"
	WalletSDK "github.com/meQlause/account-abstraction-sdk/pkg/wallet"
	WalletFactorySDK "github.com/meQlause/account-abstraction-sdk/pkg/walletfactory"
)

var (
	aaOnce sync.Once
	aaSDK  *AccountAbstractionSDK
	aaErr  error
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
	EntryPoint    *EntryPointSDK.EntryPoint
	GasManager    *GasManagerSDK.GasManager
	Wallet        *WalletSDK.Wallet
	WalletFactory *WalletFactorySDK.WalletFactory
}

func newAccountAbstractionSDK(
	ctx context.Context,
	hns AccountAbstractionHNS,
	bc *blockchain.Blockchain,
) (*AccountAbstractionSDK, error) {

	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if bc == nil {
		return nil, errors.New("blockchain cannot be nil")
	}
	if err := hns.Validate(); err != nil {
		return nil, fmt.Errorf("invalid HNS config: %w", err)
	}

	entryPoint, err := EntryPointSDK.NewEntryPointWithHNS(ctx, hns.EntryPoint, bc)
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

func InitializeAccountAbstractionSDK(ctx context.Context, hns AccountAbstractionHNS, bc *blockchain.Blockchain) error {
	aaOnce.Do(func() {
		aaSDK, aaErr = newAccountAbstractionSDK(ctx, hns, bc)
	})
	return aaErr
}

func ChangeWalletImplementationAddress(newAddress utils.Address, bc *blockchain.Blockchain) {
	contract, _ := contract.NewContract(utils.ContractConfig{
		ABIJSON: aaSDK.Wallet.Contract.ABI,
		Detail: utils.ContractDetail{
			Address:         newAddress.Hex(),
			CallBackend:     aaSDK.Wallet.Contract.Caller,
			TransactBackend: aaSDK.Wallet.Contract.Transact,
			LogBackend:      aaSDK.Wallet.Contract.Filter,
		},
	})

	aaSDK.Wallet = WalletSDK.NewWallet(
		newAddress,
		aaSDK.Wallet.ContractABI,
		bc,
		contract,
	)
}

func GetAccountAbstractionSDK() *AccountAbstractionSDK {
	sdk, err := accountAbstractionSDK()
	if err != nil {
		panic(err)
	}
	return sdk
}

func accountAbstractionSDK() (*AccountAbstractionSDK, error) {
	if aaSDK == nil {
		return nil, errors.New("AccountAbstractionSDK not initialized, call InitializeAccountAbstractionSDK first")
	}
	return aaSDK, nil
}
