package didvcdomain

import (
	"math/big"

	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type GetMetadataInput struct {
	TokenID *big.Int
	Options uint8
}

type IsCredentialValidInput struct {
	TokenID *big.Int
	Options uint8
}

type GetCredentialsWithMetadataInput struct {
	TokenIDs []*big.Int
	Options  uint8
}

type GetUnclaimedTokenIdInput struct {
	TokenID *big.Int
	Options uint8
}

type CredentialExistsInput struct {
	TokenID *big.Int
	Options uint8
}

type GetCredentialOwnerInput struct {
	TokenID *big.Int
	Options uint8
}

type TotalTokensToBeClaimedByDidInput struct {
	DIDHash utils.Hash
	Options uint8
}

type GetToBeClaimedTokensByDidInput struct {
	DIDHash utils.Hash
	Offset  *big.Int
	Limit   *big.Int
	Options uint8
}

type IsApprovedForAllInput struct {
	Owner    utils.Address
	Operator utils.Address
	Options  uint8
}

type GetFactoryInput struct {
	Options uint8
}
