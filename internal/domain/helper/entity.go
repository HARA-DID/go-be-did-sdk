package helperdomain

import (
	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
)

type StringToHex32Input struct {
	Input string
}

type EncodeCreateDIDParamInput struct {
	Address       string
	DIDParam      rootfactory.CreateDIDParam
	KeyIdentifier string
}
