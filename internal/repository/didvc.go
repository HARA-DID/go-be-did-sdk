package repository

import (
	"context"
	"math/big"

	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/nftbase"
	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/vcstorage"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

type DIDVCRepository interface {
	// NFT Base functions
	GetMetadata(ctx context.Context, input didvcdomain.GetMetadataInput) (*nftbase.CredentialMetadata, error)
	IsCredentialValid(ctx context.Context, input didvcdomain.IsCredentialValidInput) (bool, error)
	GetCredentialsWithMetadata(ctx context.Context, input didvcdomain.GetCredentialsWithMetadataInput) (*nftbase.CredentialsWithMetadataResult, error)
	GetUnclaimedTokenId(ctx context.Context, input didvcdomain.GetUnclaimedTokenIdInput) (utils.Hash, error)
	TotalTokensToBeClaimedByDid(ctx context.Context, input didvcdomain.TotalTokensToBeClaimedByDidInput) (*big.Int, error)
	GetToBeClaimedTokensByDid(ctx context.Context, input didvcdomain.GetToBeClaimedTokensByDidInput) ([]*big.Int, error)
	IsApprovedForAll(ctx context.Context, input didvcdomain.IsApprovedForAllInput) (bool, error)

	// VC Storage functions
	GetIdentityTokenCount(ctx context.Context, input didvcdomain.GetIdentityTokenCountInput) (*big.Int, error)
	GetCertificateTokenCount(ctx context.Context, input didvcdomain.GetCertificateTokenCountInput) (*big.Int, error)
	GetIdentityTokenIds(ctx context.Context, input didvcdomain.GetIdentityTokenIdsInput) (*vcstorage.TokenIdsResult, error)
	GetCertificateTokenIds(ctx context.Context, input didvcdomain.GetCertificateTokenIdsInput) (*vcstorage.TokenIdsResult, error)
	GetAllIdentityTokenIds(ctx context.Context, input didvcdomain.GetAllIdentityTokenIdsInput) ([]*big.Int, error)
	GetAllCertificateTokenIds(ctx context.Context, input didvcdomain.GetAllCertificateTokenIdsInput) ([]*big.Int, error)
	GetDIDRootStorage(ctx context.Context, input didvcdomain.GetDIDRootStorageInput) (utils.Address, error)
}
