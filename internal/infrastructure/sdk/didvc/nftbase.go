package didvcsdk

import (
	"context"
	"fmt"
	"math/big"

	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/nftbase"
	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
	backendutils "github.com/meQlause/go-be-did/utils"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (s *DIDVCSDK) GetMetadata(
	ctx context.Context,
	input didvcdomain.GetMetadataInput,
) (*nftbase.CredentialMetadata, error) {
	var resp *nftbase.CredentialMetadata
	var err error

	switch input.Options {
	case backendutils.OptionsCertificate:
		resp, err = s.CertificateNFT.GetMetadata(ctx, input.TokenID)
	case backendutils.OptionsIdentity:
		resp, err = s.IdentityNFT.GetMetadata(ctx, input.TokenID)
	default:
		return nil, fmt.Errorf("invalid option: %d", input.Options)
	}

	if err != nil {
		return nil, fmt.Errorf("get metadata failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) IsCredentialValid(
	ctx context.Context,
	input didvcdomain.IsCredentialValidInput,
) (bool, error) {
	var resp bool
	var err error

	switch input.Options {
	case backendutils.OptionsCertificate:
		resp, err = s.CertificateNFT.IsCredentialValid(ctx, input.TokenID)
	case backendutils.OptionsIdentity:
		resp, err = s.IdentityNFT.IsCredentialValid(ctx, input.TokenID)
	default:
		return false, fmt.Errorf("invalid option: %d", input.Options)
	}

	if err != nil {
		return false, fmt.Errorf("check credential validity failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetCredentialsWithMetadata(
	ctx context.Context,
	input didvcdomain.GetCredentialsWithMetadataInput,
) (*nftbase.CredentialsWithMetadataResult, error) {
	var res *nftbase.CredentialsWithMetadataResult
	var err error

	switch input.Options {
	case backendutils.OptionsCertificate:
		res, err = s.CertificateNFT.GetCredentialsWithMetadata(ctx, input.TokenIDs)
	case backendutils.OptionsIdentity:
		res, err = s.IdentityNFT.GetCredentialsWithMetadata(ctx, input.TokenIDs)
	default:
		return nil, fmt.Errorf("invalid option: %d", input.Options)
	}

	if err != nil {
		return nil, fmt.Errorf("get credentials with metadata failed: %w", err)
	}

	return res, nil
}

func (s *DIDVCSDK) GetUnclaimedTokenId(
	ctx context.Context,
	input didvcdomain.GetUnclaimedTokenIdInput,
) (utils.Hash, error) {
	var resp utils.Hash
	var err error

	switch input.Options {
	case backendutils.OptionsCertificate:
		resp, err = s.CertificateNFT.GetUnclaimedTokenId(ctx, input.TokenID)
	case backendutils.OptionsIdentity:
		resp, err = s.IdentityNFT.GetUnclaimedTokenId(ctx, input.TokenID)
	default:
		return utils.Hash{}, fmt.Errorf("invalid option: %d", input.Options)
	}

	if err != nil {
		return utils.Hash{}, fmt.Errorf("get unclaimed token id failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) TotalTokensToBeClaimedByDid(
	ctx context.Context,
	input didvcdomain.TotalTokensToBeClaimedByDidInput,
) (*big.Int, error) {
	var resp *big.Int
	var err error

	switch input.Options {
	case backendutils.OptionsCertificate:
		resp, err = s.CertificateNFT.TotalTokensToBeClaimedByDid(ctx, input.DIDHash)
	case backendutils.OptionsIdentity:
		resp, err = s.IdentityNFT.TotalTokensToBeClaimedByDid(ctx, input.DIDHash)
	default:
		return big.NewInt(0), fmt.Errorf("invalid option: %d", input.Options)
	}

	if err != nil {
		return big.NewInt(0), fmt.Errorf("get total tokens to be claimed failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetToBeClaimedTokensByDid(
	ctx context.Context,
	input didvcdomain.GetToBeClaimedTokensByDidInput,
) ([]*big.Int, error) {
	var resp []*big.Int
	var err error

	switch input.Options {
	case backendutils.OptionsCertificate:
		resp, err = s.CertificateNFT.GetToBeClaimedTokensByDid(ctx, input.DIDHash, input.Offset, input.Limit)
	case backendutils.OptionsIdentity:
		resp, err = s.IdentityNFT.GetToBeClaimedTokensByDid(ctx, input.DIDHash, input.Offset, input.Limit)
	default:
		return nil, fmt.Errorf("invalid option: %d", input.Options)
	}

	if err != nil {
		return nil, fmt.Errorf("get tokens to be claimed by DID failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) IsApprovedForAll(
	ctx context.Context,
	input didvcdomain.IsApprovedForAllInput,
) (bool, error) {
	var resp bool
	var err error

	switch input.Options {
	case backendutils.OptionsCertificate:
		resp, err = s.CertificateNFT.IsApprovedForAll(ctx, input.Owner, input.Operator)
	case backendutils.OptionsIdentity:
		resp, err = s.IdentityNFT.IsApprovedForAll(ctx, input.Owner, input.Operator)
	default:
		return false, fmt.Errorf("invalid option: %d", input.Options)
	}

	if err != nil {
		return false, fmt.Errorf("check approval for all failed: %w", err)
	}

	return resp, nil
}
