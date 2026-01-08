package didaliassdk

import (
	"context"
	"fmt"
	"math/big"

	"github.com/meQlause/alias-root-sdk/pkg/aliasfactory"
	didaliasdomain "github.com/meQlause/go-be-did/internal/domain/entities/didalias"
	backendutils "github.com/meQlause/go-be-did/utils"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (s *DIDAliasSDK) Resolve(
	ctx context.Context,
	input didaliasdomain.ResolveInput,
) (utils.Hash, error) {
	resp, err := s.AliasFactory.Resolve(ctx, input.Node)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("resolve failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) ResolveFromString(
	ctx context.Context,
	input didaliasdomain.ResolveFromStringInput,
) (utils.Hash, error) {
	resp, err := s.AliasFactory.ResolveFromString(ctx, input.Name)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("resolve from string failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) GetAliasStatus(
	ctx context.Context,
	input didaliasdomain.GetAliasStatusInput,
) (*aliasfactory.AliasStatus, error) {
	expired, isRevoked, isValid, err := s.AliasFactory.GetAliasStatus(ctx, input.Node)
	if err != nil {
		return nil, fmt.Errorf("get alias status failed: %w", err)
	}
	return &aliasfactory.AliasStatus{
		Expired:   expired,
		IsRevoked: isRevoked,
		IsValid:   isValid,
	}, nil
}

func (s *DIDAliasSDK) GetAliasStatusFromString(
	ctx context.Context,
	input didaliasdomain.GetAliasStatusFromStringInput,
) (*aliasfactory.AliasStatus, error) {
	expired, isRevoked, isValid, err := s.AliasFactory.GetAliasStatusFromString(ctx, input.Name)
	if err != nil {
		return nil, fmt.Errorf("get alias status from string failed: %w", err)
	}
	return &aliasfactory.AliasStatus{
		Expired:   expired,
		IsRevoked: isRevoked,
		IsValid:   isValid,
	}, nil
}

func (s *DIDAliasSDK) GetOwner(
	ctx context.Context,
	input didaliasdomain.GetOwnerInput,
) (string, error) {
	resp, err := s.AliasFactory.GetOwner(ctx, input.Node)
	if err != nil {
		return "", fmt.Errorf("get owner failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) GetOwnerFromString(
	ctx context.Context,
	input didaliasdomain.GetOwnerFromStringInput,
) (string, error) {
	resp, err := s.AliasFactory.GetOwnerFromString(ctx, input.Name)
	if err != nil {
		return "", fmt.Errorf("get owner from string failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) GetDID(
	ctx context.Context,
	input didaliasdomain.GetDIDInput,
) (utils.Hash, error) {
	resp, err := s.AliasFactory.GetDID(ctx, input.Node)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("get DID failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) GetDIDFromString(
	ctx context.Context,
	input didaliasdomain.GetDIDFromStringInput,
) (utils.Hash, error) {
	resp, err := s.AliasFactory.GetDIDFromString(ctx, input.Name)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("get DID from string failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) Namehash(
	ctx context.Context,
	input didaliasdomain.NamehashInput,
) (utils.Hash, error) {
	resp, err := s.AliasFactory.Namehash(ctx, input.Name)
	if err != nil {
		return utils.Hash{}, fmt.Errorf("namehash failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) GetRegistrationPeriod(
	ctx context.Context,
	input didaliasdomain.GetRegistrationPeriodInput,
) (*big.Int, error) {
	resp, err := s.AliasFactory.GetRegistrationPeriod(ctx, uint8(input.Period))
	if err != nil {
		return nil, fmt.Errorf("get registration period failed: %w", err)
	}
	return resp, nil
}

func (s *DIDAliasSDK) RegisterTLD(ctx context.Context, input didaliasdomain.RegisterTLDInput) (*backendutils.TxHash, error) {
	wallet := wallet.NewWallet(input.PrivKey)
	txHashes, err := s.AliasFactory.RegisterTLD(
		ctx,
		wallet,
		input.Input,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("register tld failed: %w", err)
	}

	return &backendutils.TxHash{
		TxHash: txHashes,
	}, nil
}
