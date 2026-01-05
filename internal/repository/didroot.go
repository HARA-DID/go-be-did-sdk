package repository

import (
	"context"
	"math/big"

	"github.com/meQlause/did-root-sdk/pkg/rootstorage"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"
)

type DIDRootRepository interface {
	ResolveDID(ctx context.Context, input didrootdomain.ResolveDIDInput) (*rootstorage.DIDDocument, error)
	VerifyDIDOwnership(ctx context.Context, input didrootdomain.VerifyDIDOwnershipInput) (bool, error)
	GetKey(ctx context.Context, input didrootdomain.GetKeyInput) (*rootstorage.Key, error)
	GetKeysByDID(ctx context.Context, input didrootdomain.GetKeysByDIDInput) ([]utils.Hash, error)
	GetClaim(ctx context.Context, input didrootdomain.GetClaimInput) (*rootstorage.Claim, error)
	GetClaimsByDID(ctx context.Context, input didrootdomain.GetClaimsByDIDInput) ([]utils.Hash, error)
	VerifyClaim(ctx context.Context, input didrootdomain.VerifyClaimInput) (bool, error)
	GetData(ctx context.Context, input didrootdomain.GetDataInput) (string, error)
	GetDIDKeyDataCount(ctx context.Context, input didrootdomain.GetDIDKeyDataCountInput) (uint64, error)
	GetDIDKeyDataByIndex(ctx context.Context, input didrootdomain.GetDIDKeyDataByIndexInput) (utils.Hash, error)
	GetOriginalKey(ctx context.Context, input didrootdomain.GetOriginalKeyInput) (string, error)
	DIDIndexMap(ctx context.Context, input didrootdomain.DIDIndexMapInput) (string, error)
	DIDIndexMapReverse(ctx context.Context, input didrootdomain.DIDIndexMapReverseInput) (*big.Int, error)
}
