package didaliasevent

import (
	"context"
	"math/big"

	didaliassdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didalias"

	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type TLDRegisteredEvent struct {
	Node  string `json:"node"`
	Owner string `json:"owner"`
	TLD   string `json:"tld"`
}

type RegisterTLDEventsResult struct {
	TLDRegistered *TLDRegisteredEvent `json:"tld_registered,omitempty"`
}

func DecodeRegisterTLDEvents(ctx context.Context, txHash utils.Hash) (*RegisterTLDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &RegisterTLDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress {
			continue
		}

		if len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["TLDRegistered"].ID && len(log.Topics) >= 3 && len(log.Data) >= 32 {
			node := utils.BytesToHash(log.Topics[1].Bytes())
			owner := utils.BytesToAddress(log.Topics[2].Bytes())

			tldOffset := new(big.Int).SetBytes(log.Data[0:32]).Uint64()
			tldLength := new(big.Int).SetBytes(log.Data[tldOffset : tldOffset+32]).Uint64()
			tld := string(log.Data[tldOffset+32 : tldOffset+32+tldLength])

			result.TLDRegistered = &TLDRegisteredEvent{
				Node:  node.Hex(),
				Owner: owner.Hex(),
				TLD:   tld,
			}
		}
	}

	return result, nil
}
