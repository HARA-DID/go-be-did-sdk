package didvcdomain

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type GetIdentityTokenCountInput struct {
	DIDHash utils.Hash
}

type GetCertificateTokenCountInput struct {
	DIDHash utils.Hash
}

type GetIdentityTokenIdsInput struct {
	DIDHash utils.Hash
	Offset  *big.Int
	Limit   *big.Int
}

type GetCertificateTokenIdsInput struct {
	DIDHash utils.Hash
	Offset  *big.Int
	Limit   *big.Int
}

type GetAllIdentityTokenIdsInput struct {
	DIDHash utils.Hash
}

type GetAllCertificateTokenIdsInput struct {
	DIDHash utils.Hash
}

type GetDIDRootStorageInput struct {
}
