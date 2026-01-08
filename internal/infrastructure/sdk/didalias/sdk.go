package didaliassdk

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	AliasFactorySDK "github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	AliasStorageSDK "github.com/meQlause/alias-root-sdk/pkg/aliasstorage"
)

var (
	aaOnce sync.Once
	aaSDK  *DIDAliasSDK
	aaErr  error
)

type DIDAliasHNS struct {
	AliasFactory string
	AliasStorage string
}

func (d DIDAliasHNS) Validate() error {
	if d.AliasFactory == "" || d.AliasStorage == "" {
		return errors.New("DID Alias factory and storage are required")
	}
	return nil
}

type DIDAliasSDK struct {
	AliasFactory *AliasFactorySDK.AliasFactory
	AliasStorage *AliasStorageSDK.AliasStorage
}

func NewDIDAliasSDK(
	ctx context.Context,
	hns DIDAliasHNS,
	bc *blockchain.Blockchain,
) (*DIDAliasSDK, error) {

	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if bc == nil {
		return nil, errors.New("blockchain cannot be nil")
	}
	if err := hns.Validate(); err != nil {
		return nil, fmt.Errorf("invalid HNS config: %w", err)
	}

	aliasFactory, err := AliasFactorySDK.NewAliasFactoryWithHNS(ctx, hns.AliasFactory, bc)
	if err != nil {
		return nil, fmt.Errorf("init entry point: %w", err)
	}

	aliasStorage, err := AliasStorageSDK.NewAliasStorageWithHNS(ctx, hns.AliasStorage, bc)
	if err != nil {
		return nil, fmt.Errorf("init gas manager: %w", err)
	}

	return &DIDAliasSDK{
		AliasFactory: aliasFactory,
		AliasStorage: aliasStorage,
	}, nil
}

func InitializeDIDAliasSDK(ctx context.Context, hns DIDAliasHNS, bc *blockchain.Blockchain) error {
	aaOnce.Do(func() {
		aaSDK, aaErr = NewDIDAliasSDK(ctx, hns, bc)
	})
	return aaErr
}

func GetDIDAliasSDK() *DIDAliasSDK {
	sdk, err := didAliasSDK()
	if err != nil {
		panic(err)
	}
	return sdk
}
func didAliasSDK() (*DIDAliasSDK, error) {
	if aaSDK == nil {
		return nil, errors.New("AccountAbstractionSDK not initialized, call InitializeAccountAbstractionSDK first")
	}
	return aaSDK, nil
}
