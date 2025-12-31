package didrootevent

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/meQlause/go-be-did/internal/config"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type DIDCreatedEvent struct {
	DIDID     *big.Int `json:"did_id"`
	Owner     string   `json:"owner"`
	Timestamp *big.Int `json:"timestamp"`
}

type KeyAddedEvent struct {
	DIDID   *big.Int `json:"did_id"`
	KeyType uint8    `json:"key_type"`
	Purpose uint8    `json:"purpose"`
	Owner   string   `json:"owner"`
}

type DIDEventsResult struct {
	DIDCreated *DIDCreatedEvent `json:"did_created,omitempty"`
	KeyAdded   *KeyAddedEvent   `json:"key_added,omitempty"`
}

var (
	didCreatedEventSignature = crypto.Keccak256Hash([]byte("DIDCreated(uint256,address,uint256)"))
	keyAddedEventSignature = crypto.Keccak256Hash([]byte("KeyAdded(uint256,uint8,uint8,address)"))
)

func DecodeCreateDIDEvents(ctx context.Context, txHash utils.Hash) (*DIDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	result := &DIDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress {
			continue
		}

		if len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == didCreatedEventSignature && len(log.Topics) >= 3 && len(log.Data) >= 32 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			owner := utils.BytesToAddress(log.Topics[2].Bytes())
			timestamp := new(big.Int).SetBytes(log.Data[:32])

			result.DIDCreated = &DIDCreatedEvent{
				DIDID:     didID,
				Owner:     owner.Hex(),
				Timestamp: timestamp,
			}
			continue
		}

		if eventSignature == keyAddedEventSignature && len(log.Topics) >= 4 && len(log.Data) >= 32 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())

			keyType := uint8(log.Topics[2].Bytes()[31]) 
			purpose := uint8(log.Topics[3].Bytes()[31])

			owner := utils.BytesToAddress(log.Data[12:32])

			result.KeyAdded = &KeyAddedEvent{
				DIDID:   didID,
				KeyType: keyType,
				Purpose: purpose,
				Owner:   owner.Hex(),
			}
		}
	}

	return result, nil
}
