package didrootevent

import (
	"context"
	"math/big"

	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"

	"github.com/meQlause/go-be-did/internal/config"
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

type DIDUpdatedEvent struct {
	DIDID     *big.Int `json:"did_id"`
	Timestamp *big.Int `json:"timestamp"`
}

type DIDDeactivatedEvent struct {
	DIDID     *big.Int `json:"did_id"`
	Timestamp *big.Int `json:"timestamp"`
}

type DIDReactivatedEvent struct {
	DIDID     *big.Int `json:"did_id"`
	Timestamp *big.Int `json:"timestamp"`
}

type DIDTransferredEvent struct {
	DIDID *big.Int `json:"did_id"`
	From  string   `json:"from"`
	To    string   `json:"to"`
}

type DataChangedEvent struct {
	DIDID *big.Int `json:"did_id"`
	Key   string   `json:"key"`
	Value string   `json:"value"`
}

type DataDeletedEvent struct {
	DIDID *big.Int `json:"did_id"`
	Key   string   `json:"key"`
}

type KeyRemovedEvent struct {
	DIDID   *big.Int `json:"did_id"`
	KeyData string   `json:"key_data"`
}

type ClaimAddedEvent struct {
	DIDID   *big.Int `json:"did_id"`
	ClaimID string   `json:"claim_id"`
	Topic   *big.Int `json:"topic"`
	Issuer  string   `json:"issuer"`
	Data    string   `json:"data"`
}

type ClaimRemovedEvent struct {
	DIDID   *big.Int `json:"did_id"`
	ClaimID string   `json:"claim_id"`
}

type CreateDIDEventsResult struct {
	DIDCreated *DIDCreatedEvent `json:"did_created,omitempty"`
	KeyAdded   *KeyAddedEvent   `json:"key_added,omitempty"`
}

type UpdateDIDEventsResult struct {
	DIDUpdated *DIDUpdatedEvent `json:"did_updated,omitempty"`
}

type DeactivateDIDEventsResult struct {
	DIDDeactivated *DIDDeactivatedEvent `json:"did_deactivated,omitempty"`
}

type ReactivateDIDEventsResult struct {
	DIDReactivated *DIDReactivatedEvent `json:"did_reactivated,omitempty"`
}

type TransferDIDEventsResult struct {
	DIDTransferred *DIDTransferredEvent `json:"did_transferred,omitempty"`
}

type StoreDataEventsResult struct {
	DataChanged *DataChangedEvent `json:"data_changed,omitempty"`
}

type DeleteDataEventsResult struct {
	DataDeleted *DataDeletedEvent `json:"data_deleted,omitempty"`
}

type AddKeyEventsResult struct {
	KeyAdded *KeyAddedEvent `json:"key_added,omitempty"`
}

type RemoveKeyEventsResult struct {
	KeyRemoved *KeyRemovedEvent `json:"key_removed,omitempty"`
}

type AddClaimEventsResult struct {
	ClaimAdded *ClaimAddedEvent `json:"claim_added,omitempty"`
}

type RemoveClaimEventsResult struct {
	ClaimRemoved *ClaimRemovedEvent `json:"claim_removed,omitempty"`
}

func DecodeCreateDIDEvents(ctx context.Context, txHash utils.Hash) (*CreateDIDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &CreateDIDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress {
			continue
		}

		if len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DIDCreated"].ID && len(log.Topics) >= 3 && len(log.Data) >= 32 {
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

		if eventSignature == events["KeyAdded"].ID && len(log.Topics) >= 4 && len(log.Data) >= 32 {
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

func DecodeUpdateDIDEvents(ctx context.Context, txHash utils.Hash) (*UpdateDIDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &UpdateDIDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DIDUpdated"].ID && len(log.Topics) >= 2 && len(log.Data) >= 32 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			timestamp := new(big.Int).SetBytes(log.Data[:32])

			result.DIDUpdated = &DIDUpdatedEvent{
				DIDID:     didID,
				Timestamp: timestamp,
			}
		}
	}

	return result, nil
}

func DecodeDeactivateDIDEvents(ctx context.Context, txHash utils.Hash) (*DeactivateDIDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &DeactivateDIDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DIDDeactivated"].ID && len(log.Topics) >= 2 && len(log.Data) >= 32 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			timestamp := new(big.Int).SetBytes(log.Data[:32])

			result.DIDDeactivated = &DIDDeactivatedEvent{
				DIDID:     didID,
				Timestamp: timestamp,
			}
		}
	}

	return result, nil
}

func DecodeReactivateDIDEvents(ctx context.Context, txHash utils.Hash) (*ReactivateDIDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &ReactivateDIDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DIDReactivated"].ID && len(log.Topics) >= 2 && len(log.Data) >= 32 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			timestamp := new(big.Int).SetBytes(log.Data[:32])

			result.DIDReactivated = &DIDReactivatedEvent{
				DIDID:     didID,
				Timestamp: timestamp,
			}
		}
	}

	return result, nil
}

func DecodeTransferDIDEvents(ctx context.Context, txHash utils.Hash) (*TransferDIDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &TransferDIDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DIDTransferred"].ID && len(log.Topics) >= 4 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			from := utils.BytesToAddress(log.Topics[2].Bytes())
			to := utils.BytesToAddress(log.Topics[3].Bytes())

			result.DIDTransferred = &DIDTransferredEvent{
				DIDID: didID,
				From:  from.Hex(),
				To:    to.Hex(),
			}
		}
	}

	return result, nil
}

func DecodeStoreDataEvents(ctx context.Context, txHash utils.Hash) (*StoreDataEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &StoreDataEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DataChanged"].ID && len(log.Topics) >= 2 && len(log.Data) >= 64 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())

			keyOffset := new(big.Int).SetBytes(log.Data[0:32]).Uint64()
			valueOffset := new(big.Int).SetBytes(log.Data[32:64]).Uint64()

			keyLength := new(big.Int).SetBytes(log.Data[keyOffset : keyOffset+32]).Uint64()
			key := string(log.Data[keyOffset+32 : keyOffset+32+keyLength])

			valueLength := new(big.Int).SetBytes(log.Data[valueOffset : valueOffset+32]).Uint64()
			value := string(log.Data[valueOffset+32 : valueOffset+32+valueLength])

			result.DataChanged = &DataChangedEvent{
				DIDID: didID,
				Key:   key,
				Value: value,
			}
		}
	}

	return result, nil
}

func DecodeDeleteDataEvents(ctx context.Context, txHash utils.Hash) (*DeleteDataEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &DeleteDataEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DataDeleted"].ID && len(log.Topics) >= 2 && len(log.Data) >= 32 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())

			keyOffset := new(big.Int).SetBytes(log.Data[0:32]).Uint64()
			keyLength := new(big.Int).SetBytes(log.Data[keyOffset : keyOffset+32]).Uint64()
			key := string(log.Data[keyOffset+32 : keyOffset+32+keyLength])

			result.DataDeleted = &DataDeletedEvent{
				DIDID: didID,
				Key:   key,
			}
		}
	}

	return result, nil
}

func DecodeAddKeyEvents(ctx context.Context, txHash utils.Hash) (*AddKeyEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &AddKeyEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["KeyAdded"].ID && len(log.Topics) >= 4 && len(log.Data) >= 32 {
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

func DecodeRemoveKeyEvents(ctx context.Context, txHash utils.Hash) (*RemoveKeyEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &RemoveKeyEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["KeyRemoved"].ID && len(log.Topics) >= 3 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			keyData := utils.BytesToHash(log.Topics[2].Bytes())

			result.KeyRemoved = &KeyRemovedEvent{
				DIDID:   didID,
				KeyData: keyData.Hex(),
			}
		}
	}

	return result, nil
}

func DecodeAddClaimEvents(ctx context.Context, txHash utils.Hash) (*AddClaimEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &AddClaimEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["ClaimAdded"].ID && len(log.Topics) >= 3 && len(log.Data) >= 96 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			claimID := utils.BytesToHash(log.Topics[2].Bytes())

			topic := new(big.Int).SetBytes(log.Data[0:32])
			issuer := utils.BytesToAddress(log.Data[44:64])

			dataOffset := new(big.Int).SetBytes(log.Data[64:96]).Uint64()
			dataLength := new(big.Int).SetBytes(log.Data[dataOffset : dataOffset+32]).Uint64()
			claimData := utils.Bytes2Hex(log.Data[dataOffset+32 : dataOffset+32+dataLength])

			result.ClaimAdded = &ClaimAddedEvent{
				DIDID:   didID,
				ClaimID: claimID.Hex(),
				Topic:   topic,
				Issuer:  issuer.Hex(),
				Data:    claimData,
			}
		}
	}

	return result, nil
}

func DecodeRemoveClaimEvents(ctx context.Context, txHash utils.Hash) (*RemoveClaimEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didrootsdk.GetDIDRootSDK().RootFactory.Address
	events := didrootsdk.GetDIDRootSDK().RootFactory.ContractABI.Events
	result := &RemoveClaimEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["ClaimRemoved"].ID && len(log.Topics) >= 3 {
			didID := new(big.Int).SetBytes(log.Topics[1].Bytes())
			claimID := utils.BytesToHash(log.Topics[2].Bytes())

			result.ClaimRemoved = &ClaimRemovedEvent{
				DIDID:   didID,
				ClaimID: claimID.Hex(),
			}
		}
	}

	return result, nil
}
