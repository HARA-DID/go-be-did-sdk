package accountabstractionsdk

import (
	"context"
	"fmt"
	"math/big"

	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/network"
	"github.com/meQlause/hara-core-blockchain-lib/pkg/wallet"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func (s *AccountAbstractionSDK) HandleOps(
	ctx context.Context,
	input aado.HandleOpsInput,
	net *network.Network,
) (*aado.TxHash, error) {
	relayerWallet := wallet.NewWallet(input.PrivKey)

	relayerPubAddress, _ := relayerWallet.GetAddress()
	latestBlock, err := net.LatestBlock(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error fetching latest block: %w", err)
	}

	data := utils.Hex2Bytes(input.Data[2:])
	messageHash := net.PackedKeccak256().
		AddAddress(input.Wallet).
		AddAddress(input.Target).
		AddUint256(big.NewInt(0)).
		AddBytes32(utils.Keccak256Hash(data).Bytes()).
		AddUint256(new(big.Int).SetUint64(latestBlock)).
		AddUint256(new(big.Int).SetUint64(input.Nonce.Uint64())).
		Hash()

	sig, err := relayerWallet.SignEIP191(messageHash.Hex())
	if err != nil {
		return nil, fmt.Errorf("error signing message: %w", err)
	}

	handleOp := aado.HandleOpsParams{
		Wallet: relayerPubAddress,
		UserOp: aado.UserOp{
			Target:            input.Target,
			Value:             big.NewInt(0),
			Data:              data,
			ClientBlockNumber: new(big.Int).SetUint64(latestBlock),
			UserNonce:         &input.Nonce,
			Signature:         utils.Hex2Bytes(sig.Signature[2:]),
		},
	}

	aaSDK.Wallet.Contract.ABI.Methods["validateUserOps"].Inputs.Pack()

	hashes, err := s.EntryPoint.HandleOps(
		ctx,
		relayerWallet,
		handleOp,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("handle ops failed: %w", err)
	}

	return &aado.TxHash{
		TxHash: hashes,
	}, nil
}
