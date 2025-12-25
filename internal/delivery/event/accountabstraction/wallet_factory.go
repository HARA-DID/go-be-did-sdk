package accountabstractionevent

import (
	"context"

	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type WalletDeployedEvent struct {
	WalletAddress   string `json:"wallet_address"`
	DeployerAddress string `json:"deployer_address"`
	Salt            string `json:"salt"`
}

func DecodeWalletDeployedEvent(ctx context.Context, txHash utils.Hash) (*WalletDeployedEvent, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := utils.HexToAddress("0x17ad613d07e9DdEeBE5D9C903E137142d49A294B")

	for _, log := range receipt.Logs {
		if log.Address != contractAddress {
			continue
		}

		if len(log.Topics) >= 4 {
			walletAddress := utils.BytesToAddress(log.Topics[1].Bytes())
			deployerAddress := utils.BytesToAddress(log.Topics[2].Bytes())
			salt := log.Topics[3].Hex()

			return &WalletDeployedEvent{
				WalletAddress:   walletAddress.Hex(),
				DeployerAddress: deployerAddress.Hex(),
				Salt:            salt,
			}, nil
		}
	}

	return nil, nil
}
