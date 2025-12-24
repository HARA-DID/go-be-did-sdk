package accountabstractionhandler

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	aado "github.com/meQlause/go-be-did/internal/domain/accountabstraction"

	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/go-be-did/pkg/response"
)

type WalletDeployedEvent struct {
	WalletAddress   string `json:"wallet_address"`
	DeployerAddress string `json:"deployer_address"`
	Salt            string `json:"salt"`
}

func (ah *AccountAbstractionHandler) CreateAccount(c *fiber.Ctx) error {
	var input aado.CreateWalletInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := ah.uc.CreateAccount(c.Context(), input)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	txSuccess, txErrors := config.Blockchain().CheckTxs(c.Context(), result.TxHash)

	resp := make(map[string]Response)

	for txHash, ok := range txSuccess {
		eventData, err := ah.decodeWalletDeployedEvent(c.Context(), txHash)
		if err != nil {
			fmt.Printf("Could not decode event for tx %s: %v\n", txHash, err)
			continue
		}
		if eventData != nil {
			resp[txHash.Hex()] = Response{
				Success:  ok,
				Errors:   "No Error Message",
				Returned: eventData,
			}
		}
	}

	for h, err := range txErrors {
		resp[h.Hex()] = Response{
			Success:  false,
			Errors:   err.Error(),
			Returned: nil,
		}
	}

	return response.Success(c, resp)
}

func (ah *AccountAbstractionHandler) decodeWalletDeployedEvent(ctx context.Context, txHash common.Hash) (*WalletDeployedEvent, error) {
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
