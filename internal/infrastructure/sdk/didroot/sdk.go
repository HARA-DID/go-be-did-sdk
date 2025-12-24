package didrootsdk

import (
	"context"
	"errors"
	"fmt"
	"sync"

	rootFactorySDK "github.com/meQlause/did-root-sdk/pkg/rootfactory"
	rootStorageSDK "github.com/meQlause/did-root-sdk/pkg/rootstorage"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"
)

var (
	didRootOnce sync.Once
	didRootSDK  *DIDRootSDK
	didRootErr  error
)

type DIDRootHNS struct {
	RootFactory string
	RootStorage string
}

func (h DIDRootHNS) Validate() error {
	if h.RootFactory == "" {
		return errors.New("root factory HNS is empty")
	}
	if h.RootStorage == "" {
		return errors.New("root storage HNS is empty")
	}
	return nil
}

type DIDRootSDK struct {
	RootFactory *rootFactorySDK.RootFactory
	RootStorage *rootStorageSDK.RootStorage
}

func newDIDRootSDK(
	ctx context.Context,
	hns DIDRootHNS,
	bc *blockchain.Blockchain,
) (*DIDRootSDK, error) {

	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if bc == nil {
		return nil, errors.New("blockchain cannot be nil")
	}
	if err := hns.Validate(); err != nil {
		return nil, fmt.Errorf("invalid HNS config: %w", err)
	}

	rootFactory, err := rootFactorySDK.NewRootFactoryWithHNS(
		ctx,
		hns.RootFactory,
		bc,
	)
	if err != nil {
		return nil, fmt.Errorf("init root factory: %w", err)
	}

	rootStorage, err := rootStorageSDK.NewRootStorageWithHNS(
		ctx,
		hns.RootStorage,
		bc,
	)
	if err != nil {
		return nil, fmt.Errorf("init root storage: %w", err)
	}

	return &DIDRootSDK{
		RootFactory: rootFactory,
		RootStorage: rootStorage,
	}, nil
}

func InitializeDIDRootSDK(ctx context.Context, hns DIDRootHNS, bc *blockchain.Blockchain) error {
	didRootOnce.Do(func() {
		didRootSDK, didRootErr = newDIDRootSDK(ctx, hns, bc)
	})
	return didRootErr
}

func GetDIDRootSDK() *DIDRootSDK {
	sdk, err := rootSDK()
	if err != nil {
		panic(err)
	}
	return sdk
}

func rootSDK() (*DIDRootSDK, error) {
	if didRootSDK == nil {
		return nil, errors.New("DIDRootSDK not initialized, call InitializeDIDRootSDK first")
	}
	return didRootSDK, nil
}
