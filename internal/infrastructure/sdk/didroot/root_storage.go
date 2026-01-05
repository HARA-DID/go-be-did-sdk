package didrootsdk

import (
	"context"
	"fmt"
	"math/big"

	didrootdomain "github.com/meQlause/go-be-did/internal/domain/entities/didroot"

	"github.com/meQlause/did-root-sdk/pkg/rootstorage"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (s *DIDRootSDK) GetData(
	ctx context.Context,
	input didrootdomain.GetDataInput,
) (string, error) {
	resp, err := s.RootStorage.GetData(
		ctx,
		input.Hash,
	)
	if err != nil {
		return "", fmt.Errorf("get data failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) ResolveDID(
	ctx context.Context,
	input didrootdomain.ResolveDIDInput,
) (*rootstorage.DIDDocument, error) {
	resp, err := s.RootStorage.ResolveDID(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return nil, fmt.Errorf("resolve DID failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) VerifyDIDOwnership(
	ctx context.Context,
	input didrootdomain.VerifyDIDOwnershipInput,
) (bool, error) {
	resp, err := s.RootStorage.VerifyDIDOwnership(
		ctx,
		input.DIDHash,
		input.Owner,
	)
	if err != nil {
		return false, fmt.Errorf("verify DID ownership failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) GetKey(
	ctx context.Context,
	input didrootdomain.GetKeyInput,
) (*rootstorage.Key, error) {
	resp, err := s.RootStorage.GetKey(
		ctx,
		input.DIDHash,
		input.KeyDataHashed,
	)
	if err != nil {
		return nil, fmt.Errorf("get key failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) GetKeysByDID(
	ctx context.Context,
	input didrootdomain.GetKeysByDIDInput,
) ([]utils.Hash, error) {
	resp, err := s.RootStorage.GetKeysByDID(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return nil, fmt.Errorf("get keys by DID failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) GetClaim(
	ctx context.Context,
	input didrootdomain.GetClaimInput,
) (*rootstorage.Claim, error) {
	resp, err := s.RootStorage.GetClaim(
		ctx,
		input.DIDHash,
		input.ClaimID,
	)
	if err != nil {
		return nil, fmt.Errorf("get claim failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) GetClaimsByDID(
	ctx context.Context,
	input didrootdomain.GetClaimsByDIDInput,
) ([]utils.Hash, error) {
	resp, err := s.RootStorage.GetClaimsByDID(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return nil, fmt.Errorf("get claims by DID failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) VerifyClaim(
	ctx context.Context,
	input didrootdomain.VerifyClaimInput,
) (bool, error) {
	resp, err := s.RootStorage.VerifyClaim(
		ctx,
		input.DIDHash,
		input.ClaimID,
		input.ToVerify,
	)
	if err != nil {
		return false, fmt.Errorf("verify claim failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) GetDIDKeyDataCount(
	ctx context.Context,
	input didrootdomain.GetDIDKeyDataCountInput,
) (uint64, error) {
	resp, err := s.RootStorage.GetDIDKeyDataCount(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return 0, fmt.Errorf("get DID key count failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) GetDIDKeyDataByIndex(
	ctx context.Context,
	input didrootdomain.GetDIDKeyDataByIndexInput,
) (utils.Hash, error) {
	resp, err := s.RootStorage.GetDIDKeyDataByIndex(
		ctx,
		input.DIDHash,
		input.Index,
	)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("get DID key by index failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) GetOriginalKey(
	ctx context.Context,
	input didrootdomain.GetOriginalKeyInput,
) (string, error) {
	resp, err := s.RootStorage.GetOriginalKey(
		ctx,
		input.KeyCode,
	)
	if err != nil {
		return "", fmt.Errorf("get original key failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) DIDIndexMap(
	ctx context.Context,
	input didrootdomain.DIDIndexMapInput,
) (string, error) {
	resp, err := s.RootStorage.DIDIndexMap(
		ctx,
		input.DIDIndex,
	)
	if err != nil {
		return "", fmt.Errorf("DID index map failed: %w", err)
	}

	return resp, nil
}

func (s *DIDRootSDK) DIDIndexMapReverse(
	ctx context.Context,
	input didrootdomain.DIDIndexMapReverseInput,
) (*big.Int, error) {
	resp, err := s.RootStorage.DIDIndexMapReverse(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return nil, fmt.Errorf("DID index map reverse failed: %w", err)
	}

	return resp, nil
}
