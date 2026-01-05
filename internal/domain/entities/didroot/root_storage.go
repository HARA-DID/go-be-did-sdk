package didrootdomain

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type GetDataInput struct {
	Hash utils.Hash
}

type ResolveDIDInput struct {
	DIDHash utils.Hash
}

type VerifyDIDOwnershipInput struct {
	DIDHash utils.Hash
	Owner   utils.Address
}

type GetKeyInput struct {
	DIDHash       utils.Hash
	KeyDataHashed utils.Hash
}

type GetKeysByDIDInput struct {
	DIDHash utils.Hash
}

type GetClaimInput struct {
	DIDHash utils.Hash
	ClaimID utils.Hash
}

type GetClaimsByDIDInput struct {
	DIDHash utils.Hash
}

type VerifyClaimInput struct {
	DIDHash  utils.Hash
	ClaimID  utils.Hash
	ToVerify utils.Address
}

type GetDIDKeyDataCountInput struct {
	DIDHash utils.Hash
}

type GetDIDKeyDataByIndexInput struct {
	DIDHash utils.Hash
	Index   uint64
}

type GetOriginalKeyInput struct {
	KeyCode utils.Hash
}

type DIDIndexMapInput struct {
	DIDIndex *big.Int
}

type DIDIndexMapReverseInput struct {
	DIDHash utils.Hash
}
