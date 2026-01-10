package didvcsdk

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/meQlause/hara-core-blockchain-lib/pkg/blockchain"

	NFTBaseSDK "github.com/meQlause/did-verifiable-credentials-sdk/pkg/nftbase"
	VCFactorySDK "github.com/meQlause/did-verifiable-credentials-sdk/pkg/vcfactory"
	VCStorageSDK "github.com/meQlause/did-verifiable-credentials-sdk/pkg/vcstorage"
)

var (
	vcOnce sync.Once
	vcSDK  *DIDVCSDK
	vcErr  error
)

type DIDVCHNS struct {
	VCFactory      string
	VCStorage      string
	CertificateNFT string
	IdentityNFT    string
}

func (d DIDVCHNS) Validate() error {
	if d.VCFactory == "" || d.VCStorage == "" {
		return errors.New("DID VC factory and storage are required")
	}
	if d.CertificateNFT == "" || d.IdentityNFT == "" {
		return errors.New("DID VC certificate NFT and identity NFT are required")
	}
	return nil
}

type DIDVCSDK struct {
	VCFactory      *VCFactorySDK.VCFactory
	VCStorage      *VCStorageSDK.VCStorage
	CertificateNFT *NFTBaseSDK.NFTBase
	IdentityNFT    *NFTBaseSDK.NFTBase
}

func NewDIDVCSDK(
	ctx context.Context,
	hns DIDVCHNS,
	bc *blockchain.Blockchain,
) (*DIDVCSDK, error) {

	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if bc == nil {
		return nil, errors.New("blockchain cannot be nil")
	}
	if err := hns.Validate(); err != nil {
		return nil, fmt.Errorf("invalid HNS config: %w", err)
	}

	vcFactory, err := VCFactorySDK.NewVCFactoryWithHNS(ctx, hns.VCFactory, bc)
	if err != nil {
		return nil, fmt.Errorf("init vc factory: %w", err)
	}

	vcStorage, err := VCStorageSDK.NewVCStorageWithHNS(ctx, hns.VCStorage, bc)
	if err != nil {
		return nil, fmt.Errorf("init vc storage: %w", err)
	}

	certificateNFT, err := NFTBaseSDK.NewNFTBaseWithHNS(ctx, hns.CertificateNFT, bc)
	if err != nil {
		return nil, fmt.Errorf("init certificate nft: %w", err)
	}

	identityNFT, err := NFTBaseSDK.NewNFTBaseWithHNS(ctx, hns.IdentityNFT, bc)
	if err != nil {
		return nil, fmt.Errorf("init identity nft: %w", err)
	}

	return &DIDVCSDK{
		VCFactory:      vcFactory,
		VCStorage:      vcStorage,
		CertificateNFT: certificateNFT,
		IdentityNFT:    identityNFT,
	}, nil
}

func InitializeDIDVCSDK(ctx context.Context, hns DIDVCHNS, bc *blockchain.Blockchain) error {
	vcOnce.Do(func() {
		vcSDK, vcErr = NewDIDVCSDK(ctx, hns, bc)
	})
	return vcErr
}

func GetDIDVCSDK() *DIDVCSDK {
	sdk, err := didVCSDK()
	if err != nil {
		panic(err)
	}
	return sdk
}

func didVCSDK() (*DIDVCSDK, error) {
	if vcSDK == nil {
		return nil, errors.New("DIDVCSDK not initialized, call InitializeDIDVCSDK first")
	}
	return vcSDK, nil
}
