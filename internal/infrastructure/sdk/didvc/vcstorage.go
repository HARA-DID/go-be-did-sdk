package didvcsdk

import (
	"context"
	"fmt"
	"math/big"

	"github.com/meQlause/did-verifiable-credentials-sdk/pkg/vcstorage"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didvcdomain "github.com/meQlause/go-be-did/internal/domain/entities/didvc"
)

func (s *DIDVCSDK) GetIdentityTokenCount(
	ctx context.Context,
	input didvcdomain.GetIdentityTokenCountInput,
) (*big.Int, error) {
	resp, err := s.VCStorage.GetIdentityTokenCount(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("get identity token count failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetCertificateTokenCount(
	ctx context.Context,
	input didvcdomain.GetCertificateTokenCountInput,
) (*big.Int, error) {
	resp, err := s.VCStorage.GetCertificateTokenCount(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("get certificate token count failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetIdentityTokenIds(
	ctx context.Context,
	input didvcdomain.GetIdentityTokenIdsInput,
) (*vcstorage.TokenIdsResult, error) {
	resp, err := s.VCStorage.GetIdentityTokenIds(
		ctx,
		input.DIDHash,
		input.Offset,
		input.Limit,
	)
	if err != nil {
		return nil, fmt.Errorf("get identity token ids failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetCertificateTokenIds(
	ctx context.Context,
	input didvcdomain.GetCertificateTokenIdsInput,
) (*vcstorage.TokenIdsResult, error) {
	resp, err := s.VCStorage.GetCertificateTokenIds(
		ctx,
		input.DIDHash,
		input.Offset,
		input.Limit,
	)
	if err != nil {
		return nil, fmt.Errorf("get certificate token ids failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetAllIdentityTokenIds(
	ctx context.Context,
	input didvcdomain.GetAllIdentityTokenIdsInput,
) ([]*big.Int, error) {
	resp, err := s.VCStorage.GetAllIdentityTokenIds(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return nil, fmt.Errorf("get all identity token ids failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetAllCertificateTokenIds(
	ctx context.Context,
	input didvcdomain.GetAllCertificateTokenIdsInput,
) ([]*big.Int, error) {
	resp, err := s.VCStorage.GetAllCertificateTokenIds(
		ctx,
		input.DIDHash,
	)
	if err != nil {
		return nil, fmt.Errorf("get all certificate token ids failed: %w", err)
	}

	return resp, nil
}

func (s *DIDVCSDK) GetDIDRootStorage(
	ctx context.Context,
	input didvcdomain.GetDIDRootStorageInput,
) (utils.Address, error) {
	resp, err := s.VCStorage.GetDIDRootStorage(ctx)
	if err != nil {
		return utils.Address{}, fmt.Errorf("get DID root storage failed: %w", err)
	}

	return resp, nil
}
