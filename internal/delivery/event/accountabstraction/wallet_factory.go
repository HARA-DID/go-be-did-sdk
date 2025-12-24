package accountabstractionevent

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/meQlause/go-be-did/internal/config"
)

type WalletDeployedEvent struct {
	WalletAddress   string `json:"wallet_address"`
	DeployerAddress string `json:"deployer_address"`
	Salt            string `json:"salt"`
}

func DecodeWalletDeployedEvent(ctx context.Context, txHash common.Hash) (*WalletDeployedEvent, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		fmt.Println("Error getting receipt:", err.Error())
		return nil, err
	}

	contractAddress := common.HexToAddress("0x17ad613d07e9DdEeBE5D9C903E137142d49A294B")

	for _, log := range receipt.Logs {
		if log.Address != contractAddress {
			continue
		}

		if len(log.Topics) >= 4 {
			walletAddress := common.BytesToAddress(log.Topics[1].Bytes())
			deployerAddress := common.BytesToAddress(log.Topics[2].Bytes())
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
