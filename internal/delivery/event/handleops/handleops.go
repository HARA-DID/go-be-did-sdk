package didevent

import (
	"context"

	backendutils "github.com/meQlause/go-be-did/utils"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

var Registry = map[uint8]map[uint8]DecodeFunc{
	backendutils.ServiceDIDRoot: {
		backendutils.TypeCreateDID:     decodeCreateDID,
		backendutils.TypeUpdateDID:     decodeUpdateDID,
		backendutils.TypeDeactivateDID: decodeDeactivateDID,
		backendutils.TypeReactivateDID: decodeReactivateDID,
		backendutils.TypeTransferDID:   decodeTransferDID,

		backendutils.TypeStoreData:  decodeStoreData,
		backendutils.TypeDeleteData: decodeDeleteData,

		backendutils.TypeAddKey:    decodeAddKey,
		backendutils.TypeRemoveKey: decodeRemoveKey,

		backendutils.TypeAddClaim:    decodeAddClaim,
		backendutils.TypeRemoveClaim: decodeRemoveClaim,
	}, backendutils.ServiceDIDAlias: {
		backendutils.TypeSetDIDRootStorage:      decodeSetDIDRootStorage,
		backendutils.TypeRegisterTLD:            decodeRegisterTLD,
		backendutils.TypeRegisterDomain:         decodeRegisterDomain,
		backendutils.TypeRegisterSubdomain:      decodeRegisterSubdomain,
		backendutils.TypeSetDID:                 decodeSetDID,
		backendutils.TypeExtendRegistration:     decodeExtendRegistration,
		backendutils.TypeRevokeAlias:            decodeRevokeAlias,
		backendutils.TypeUnrevokeAlias:          decodeUnrevokeAlias,
		backendutils.TypeTransferAliasOwnership: decodeTransferAliasOwnership,
	},
}

type DecodeFunc func(ctx context.Context, txHash utils.Hash) (any, error)

// DID Root
func decodeCreateDID(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeCreateDIDEvents(ctx, txHash)
}

func decodeUpdateDID(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeUpdateDIDEvents(ctx, txHash)
}

func decodeDeactivateDID(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeDeactivateDIDEvents(ctx, txHash)
}

func decodeReactivateDID(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeReactivateDIDEvents(ctx, txHash)
}

func decodeTransferDID(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeTransferDIDEvents(ctx, txHash)
}

func decodeStoreData(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeStoreDataEvents(ctx, txHash)
}

func decodeDeleteData(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeDeleteDataEvents(ctx, txHash)
}

func decodeAddKey(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeAddKeyEvents(ctx, txHash)
}

func decodeRemoveKey(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeRemoveKeyEvents(ctx, txHash)
}

func decodeAddClaim(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeAddClaimEvents(ctx, txHash)
}

func decodeRemoveClaim(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeRemoveClaimEvents(ctx, txHash)
}

// DID Alias

func decodeSetDIDRootStorage(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeSetDIDRootStorageEvents(ctx, txHash)
}

func decodeRegisterTLD(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeRegisterTLDEvents(ctx, txHash)
}

func decodeRegisterDomain(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeRegisterDomainEvents(ctx, txHash)
}

func decodeRegisterSubdomain(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeRegisterSubdomainEvents(ctx, txHash)
}

func decodeSetDID(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeSetDIDEvents(ctx, txHash)
}

func decodeExtendRegistration(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeExtendRegistrationEvents(ctx, txHash)
}

func decodeRevokeAlias(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeRevokeAliasEvents(ctx, txHash)
}

func decodeUnrevokeAlias(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeUnrevokeAliasEvents(ctx, txHash)
}

func decodeTransferAliasOwnership(ctx context.Context, txHash utils.Hash) (any, error) {
	return DecodeTransferAliasOwnershipEvents(ctx, txHash)
}
