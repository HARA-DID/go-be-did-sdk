package didrootevent

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
	},
}

type DecodeFunc func(ctx context.Context, txHash utils.Hash) (any, error)

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
