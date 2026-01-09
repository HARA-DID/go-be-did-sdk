package didevent

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

type DomainRegisteredEvent struct {
	Node       string   `json:"node"`
	Owner      string   `json:"owner"`
	Domain     string   `json:"domain"`
	ExpiryTime *big.Int `json:"expiry_time"`
}

type DIDSetEvent struct {
	Node string `json:"node"`
	DID  string `json:"did"`
}

type NameServerOwnershipTransferredEvent struct {
	Node          string `json:"node"`
	PreviousOwner string `json:"previous_owner"`
	NewOwner      string `json:"new_owner"`
}

type DIDRootStorageUpdatedEvent struct {
	OldAddress string `json:"old_address"`
	NewAddress string `json:"new_address"`
}

type FactoryContractUpdatedEvent struct {
	OldFactory string `json:"old_factory"`
	NewFactory string `json:"new_factory"`
}

type AliasRevokedEvent struct {
	Node string `json:"node"`
}

type AliasUnrevokedEvent struct {
	Node string `json:"node"`
}

type RegistrationExtendedEvent struct {
	Node          string   `json:"node"`
	NewExpiryTime *big.Int `json:"new_expiry_time"`
}

type RegisterTLDEventsResult struct {
	TLDRegistered *TLDRegisteredEvent `json:"tld_registered,omitempty"`
}

type RegisterDomainEventsResult struct {
	DomainRegistered *DomainRegisteredEvent `json:"domain_registered,omitempty"`
}

type RegisterSubdomainEventsResult struct {
	DomainRegistered *DomainRegisteredEvent `json:"domain_registered,omitempty"`
}

type SetDIDEventsResult struct {
	DIDSet *DIDSetEvent `json:"did_set,omitempty"`
}

type ExtendRegistrationEventsResult struct {
	RegistrationExtended *RegistrationExtendedEvent `json:"registration_extended,omitempty"`
}

type RevokeAliasEventsResult struct {
	AliasRevoked *AliasRevokedEvent `json:"alias_revoked,omitempty"`
}

type UnrevokeAliasEventsResult struct {
	AliasUnrevoked *AliasUnrevokedEvent `json:"alias_unrevoked,omitempty"`
}

type TransferAliasOwnershipEventsResult struct {
	NameServerOwnershipTransferred *NameServerOwnershipTransferredEvent `json:"name_server_ownership_transferred,omitempty"`
}

type SetDIDRootStorageEventsResult struct {
	DIDRootStorageUpdated *DIDRootStorageUpdatedEvent `json:"did_root_storage_updated,omitempty"`
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
		if log.Address != contractAddress || len(log.Topics) == 0 {
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

func DecodeRegisterDomainEvents(ctx context.Context, txHash utils.Hash) (*RegisterDomainEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &RegisterDomainEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DomainRegistered"].ID && len(log.Topics) >= 3 && len(log.Data) >= 64 {
			node := utils.BytesToHash(log.Topics[1].Bytes())
			owner := utils.BytesToAddress(log.Topics[2].Bytes())

			domainOffset := new(big.Int).SetBytes(log.Data[0:32]).Uint64()
			expiryTime := new(big.Int).SetBytes(log.Data[32:64])

			domainLength := new(big.Int).SetBytes(log.Data[domainOffset : domainOffset+32]).Uint64()
			domain := string(log.Data[domainOffset+32 : domainOffset+32+domainLength])

			result.DomainRegistered = &DomainRegisteredEvent{
				Node:       node.Hex(),
				Owner:      owner.Hex(),
				Domain:     domain,
				ExpiryTime: expiryTime,
			}
		}
	}

	return result, nil
}

func DecodeRegisterSubdomainEvents(ctx context.Context, txHash utils.Hash) (*RegisterSubdomainEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &RegisterSubdomainEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DomainRegistered"].ID && len(log.Topics) >= 3 && len(log.Data) >= 64 {
			node := utils.BytesToHash(log.Topics[1].Bytes())
			owner := utils.BytesToAddress(log.Topics[2].Bytes())

			domainOffset := new(big.Int).SetBytes(log.Data[0:32]).Uint64()
			expiryTime := new(big.Int).SetBytes(log.Data[32:64])

			domainLength := new(big.Int).SetBytes(log.Data[domainOffset : domainOffset+32]).Uint64()
			domain := string(log.Data[domainOffset+32 : domainOffset+32+domainLength])

			result.DomainRegistered = &DomainRegisteredEvent{
				Node:       node.Hex(),
				Owner:      owner.Hex(),
				Domain:     domain,
				ExpiryTime: expiryTime,
			}
		}
	}

	return result, nil
}

func DecodeSetDIDEvents(ctx context.Context, txHash utils.Hash) (*SetDIDEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &SetDIDEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DIDSet"].ID && len(log.Topics) >= 3 {
			node := utils.BytesToHash(log.Topics[1].Bytes())
			did := utils.BytesToHash(log.Topics[2].Bytes())

			result.DIDSet = &DIDSetEvent{
				Node: node.Hex(),
				DID:  did.Hex(),
			}
		}
	}

	return result, nil
}

func DecodeExtendRegistrationEvents(ctx context.Context, txHash utils.Hash) (*ExtendRegistrationEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &ExtendRegistrationEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["RegistrationExtended"].ID && len(log.Topics) >= 2 && len(log.Data) >= 32 {
			node := utils.BytesToHash(log.Topics[1].Bytes())
			newExpiryTime := new(big.Int).SetBytes(log.Data[:32])

			result.RegistrationExtended = &RegistrationExtendedEvent{
				Node:          node.Hex(),
				NewExpiryTime: newExpiryTime,
			}
		}
	}

	return result, nil
}

func DecodeRevokeAliasEvents(ctx context.Context, txHash utils.Hash) (*RevokeAliasEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &RevokeAliasEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["AliasRevoked"].ID && len(log.Topics) >= 2 {
			node := utils.BytesToHash(log.Topics[1].Bytes())

			result.AliasRevoked = &AliasRevokedEvent{
				Node: node.Hex(),
			}
		}
	}

	return result, nil
}

func DecodeUnrevokeAliasEvents(ctx context.Context, txHash utils.Hash) (*UnrevokeAliasEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &UnrevokeAliasEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["AliasUnrevoked"].ID && len(log.Topics) >= 2 {
			node := utils.BytesToHash(log.Topics[1].Bytes())

			result.AliasUnrevoked = &AliasUnrevokedEvent{
				Node: node.Hex(),
			}
		}
	}

	return result, nil
}

func DecodeTransferAliasOwnershipEvents(ctx context.Context, txHash utils.Hash) (*TransferAliasOwnershipEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &TransferAliasOwnershipEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["NameServerOwnershipTransferred"].ID && len(log.Topics) >= 4 {
			node := utils.BytesToHash(log.Topics[1].Bytes())
			previousOwner := utils.BytesToAddress(log.Topics[2].Bytes())
			newOwner := utils.BytesToAddress(log.Topics[3].Bytes())

			result.NameServerOwnershipTransferred = &NameServerOwnershipTransferredEvent{
				Node:          node.Hex(),
				PreviousOwner: previousOwner.Hex(),
				NewOwner:      newOwner.Hex(),
			}
		}
	}

	return result, nil
}

func DecodeSetDIDRootStorageEvents(ctx context.Context, txHash utils.Hash) (*SetDIDRootStorageEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didaliassdk.GetDIDAliasSDK().AliasStorage.Address
	events := didaliassdk.GetDIDAliasSDK().AliasStorage.ContractABI.Events
	result := &SetDIDRootStorageEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["DIDRootStorageUpdated"].ID && len(log.Topics) >= 3 {
			oldAddress := utils.BytesToAddress(log.Topics[1].Bytes())
			newAddress := utils.BytesToAddress(log.Topics[2].Bytes())

			result.DIDRootStorageUpdated = &DIDRootStorageUpdatedEvent{
				OldAddress: oldAddress.Hex(),
				NewAddress: newAddress.Hex(),
			}
		}
	}

	return result, nil
}
