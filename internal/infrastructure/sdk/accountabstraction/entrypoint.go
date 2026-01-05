package accountabstractionsdk

import (
	"context"
	"fmt"
	"math/big"

	aado "github.com/meQlause/go-be-did/internal/domain/entities/accountabstraction"
	backendutils "github.com/meQlause/go-be-did/utils"

	"github.com/meQlause/account-abstraction-sdk/pkg/entrypoint"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (s *AccountAbstractionSDK) HandleOps(
	ctx context.Context,
	input aado.HandleOpsInput,
	net *network.Network,
) (*backendutils.TxHash, error) {
	relayerWallet := wallet.NewWallet(input.PrivKey)

	latestBlock, err := net.LatestBlock(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error fetching latest block: %w", err)
	}
	latestBlock += 10

	data := utils.Hex2Bytes(input.Data)

	messageHash := net.PackedKeccak256().
		AddAddress(input.Wallet).
		AddAddress(input.Target).
		AddUint256(big.NewInt(0)).
		AddBytes32(utils.Keccak256Hash(data)).
		AddUint256(new(big.Int).SetUint64(latestBlock)).
		AddUint256(new(big.Int).SetUint64(input.Nonce.Uint64())).
		Hash()

	sig, err := relayerWallet.SignEIP191Message(messageHash.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error signing message: %w", err)
	}

	handleOp := entrypoint.HandleOpsParams{
		Wallet: input.Wallet,
		UserOp: entrypoint.UserOp{
			Target:            input.Target,
			Value:             big.NewInt(0),
			Data:              data,
			ClientBlockNumber: new(big.Int).SetUint64(latestBlock),
			UserNonce:         input.Nonce,
			Signature:         utils.Hex2Bytes(sig.Signature),
		},
	}

	hashes, err := s.EntryPoint.HandleOps(
		ctx,
		relayerWallet,
		handleOp,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("handle ops failed: %w", err)
	}

	return &backendutils.TxHash{
		TxHash: hashes,
	}, nil
}

func (s *AccountAbstractionSDK) IsValidWallet(
	ctx context.Context,
	input aado.IsValidWalletInput,
) (bool, error) {
	resp, err := s.EntryPoint.IsValidWallet(
		ctx,
		input.Wallet,
	)
	if err != nil {
		return false, fmt.Errorf("Error : %w", err)
	}

	return resp, nil
}
