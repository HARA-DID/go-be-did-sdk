package sdk

import (
	"context"
	"errors"
	"fmt"

	rootFactorySDK "github.com/meQlause/did-root-sdk/pkg/rootfactory"
	rootStorageSDK "github.com/meQlause/did-root-sdk/pkg/rootstorage"
	pkg "github.com/meQlause/hara-core-blockchain-lib/pkg"
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

func NewDIDRootSDK(
	ctx context.Context,
	hns DIDRootHNS,
	bc *pkg.Blockchain,
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
