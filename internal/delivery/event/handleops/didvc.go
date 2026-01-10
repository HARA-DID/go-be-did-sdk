package didevent

import (
	"context"
	"math/big"

	"github.com/meQlause/go-be-did/internal/config"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	didvcsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didvc"
)

// Event structs
type IdentityIssuedEvent struct {
	DID       string   `json:"did"`
	Option    uint8    `json:"option"`
	TokenId   *big.Int `json:"token_id"`
	Recipient string   `json:"recipient"`
}

type CertificateIssuedEvent struct {
	DID       string   `json:"did"`
	Option    uint8    `json:"option"`
	TokenId   *big.Int `json:"token_id"`
	Recipient string   `json:"recipient"`
}

type MetadataUpdatedEvent struct {
	Option  uint8    `json:"option"`
	TokenId *big.Int `json:"token_id"`
}

type CredentialRevokedEvent struct {
	Option  uint8    `json:"option"`
	TokenId *big.Int `json:"token_id"`
}

type CredentialClaimedEvent struct {
	NFTAddress string   `json:"nft_address"`
	TokenId    *big.Int `json:"token_id"`
}

type CredentialBurnedEvent struct {
	DID        string   `json:"did"`
	NFTAddress string   `json:"nft_address"`
	TokenId    *big.Int `json:"token_id"`
}

// Result structs
type IssueCredentialEventsResult struct {
	IdentityIssued    *IdentityIssuedEvent    `json:"identity_issued,omitempty"`
	CertificateIssued *CertificateIssuedEvent `json:"certificate_issued,omitempty"`
}

type UpdateMetadataEventsResult struct {
	MetadataUpdated *MetadataUpdatedEvent `json:"metadata_updated,omitempty"`
}

type RevokeCredentialEventsResult struct {
	CredentialRevoked *CredentialRevokedEvent `json:"credential_revoked,omitempty"`
}

type ClaimCredentialEventsResult struct {
	CredentialClaimed *CredentialClaimedEvent `json:"credential_claimed,omitempty"`
}

type BurnCredentialEventsResult struct {
	CredentialBurned *CredentialBurnedEvent `json:"credential_burned,omitempty"`
}

func DecodeIssueCredentialEvents(ctx context.Context, txHash utils.Hash) (*IssueCredentialEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didvcsdk.GetDIDVCSDK().VCFactory.Address
	events := didvcsdk.GetDIDVCSDK().VCFactory.ContractABI.Events
	result := &IssueCredentialEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress {
			continue
		}

		if len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["IdentityIssued"].ID && len(log.Topics) >= 3 && len(log.Data) >= 64 {
			did := utils.BytesToHash(log.Topics[1].Bytes())
			recipient := utils.BytesToHash(log.Topics[2].Bytes())

			option := uint8(new(big.Int).SetBytes(log.Data[0:32]).Uint64())
			tokenId := new(big.Int).SetBytes(log.Data[32:64])

			result.IdentityIssued = &IdentityIssuedEvent{
				DID:       did.Hex(),
				Option:    option,
				TokenId:   tokenId,
				Recipient: recipient.Hex(),
			}
			continue
		}

		if eventSignature == events["CertificateIssued"].ID && len(log.Topics) >= 3 && len(log.Data) >= 64 {
			did := utils.BytesToHash(log.Topics[1].Bytes())
			recipient := utils.BytesToHash(log.Topics[2].Bytes())

			option := uint8(new(big.Int).SetBytes(log.Data[0:32]).Uint64())
			tokenId := new(big.Int).SetBytes(log.Data[32:64])

			result.CertificateIssued = &CertificateIssuedEvent{
				DID:       did.Hex(),
				Option:    option,
				TokenId:   tokenId,
				Recipient: recipient.Hex(),
			}
		}
	}

	return result, nil
}

func DecodeUpdateMetadataEvents(ctx context.Context, txHash utils.Hash) (*UpdateMetadataEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didvcsdk.GetDIDVCSDK().VCFactory.Address
	events := didvcsdk.GetDIDVCSDK().VCFactory.ContractABI.Events
	result := &UpdateMetadataEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["MetadataUpdated"].ID && len(log.Topics) >= 2 && len(log.Data) >= 32 {
			tokenId := new(big.Int).SetBytes(log.Topics[1].Bytes())
			option := uint8(new(big.Int).SetBytes(log.Data[0:32]).Uint64())

			result.MetadataUpdated = &MetadataUpdatedEvent{
				Option:  option,
				TokenId: tokenId,
			}
		}
	}

	return result, nil
}

func DecodeRevokeCredentialEvents(ctx context.Context, txHash utils.Hash) (*RevokeCredentialEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didvcsdk.GetDIDVCSDK().VCFactory.Address
	events := didvcsdk.GetDIDVCSDK().VCFactory.ContractABI.Events
	result := &RevokeCredentialEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		// CredentialRevoked(Options option, uint256 indexed tokenId)
		if eventSignature == events["CredentialRevoked"].ID && len(log.Topics) >= 2 && len(log.Data) >= 32 {
			tokenId := new(big.Int).SetBytes(log.Topics[1].Bytes())
			option := uint8(new(big.Int).SetBytes(log.Data[0:32]).Uint64())

			result.CredentialRevoked = &CredentialRevokedEvent{
				Option:  option,
				TokenId: tokenId,
			}
		}
	}

	return result, nil
}

func DecodeClaimCredentialEvents(ctx context.Context, txHash utils.Hash) (*ClaimCredentialEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didvcsdk.GetDIDVCSDK().VCFactory.Address
	events := didvcsdk.GetDIDVCSDK().VCFactory.ContractABI.Events
	result := &ClaimCredentialEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["CredentialClaimed"].ID && len(log.Topics) >= 3 {
			nftAddress := utils.BytesToAddress(log.Topics[1].Bytes())
			tokenId := new(big.Int).SetBytes(log.Topics[2].Bytes())

			result.CredentialClaimed = &CredentialClaimedEvent{
				NFTAddress: nftAddress.Hex(),
				TokenId:    tokenId,
			}
		}
	}

	return result, nil
}

func DecodeBurnCredentialEvents(ctx context.Context, txHash utils.Hash) (*BurnCredentialEventsResult, error) {
	receipt, err := config.Blockchain().Network.GetPrimaryClient().Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	contractAddress := didvcsdk.GetDIDVCSDK().VCFactory.Address
	events := didvcsdk.GetDIDVCSDK().VCFactory.ContractABI.Events
	result := &BurnCredentialEventsResult{}

	for _, log := range receipt.Logs {
		if log.Address != contractAddress || len(log.Topics) == 0 {
			continue
		}

		eventSignature := log.Topics[0]

		if eventSignature == events["CredentialBurned"].ID && len(log.Topics) >= 4 {
			did := utils.BytesToHash(log.Topics[1].Bytes())
			nftAddress := utils.BytesToAddress(log.Topics[2].Bytes())
			tokenId := new(big.Int).SetBytes(log.Topics[3].Bytes())

			result.CredentialBurned = &CredentialBurnedEvent{
				DID:        did.Hex(),
				NFTAddress: nftAddress.Hex(),
				TokenId:    tokenId,
			}
		}
	}

	return result, nil
}
