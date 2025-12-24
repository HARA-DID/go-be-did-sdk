package helperdomain

import (
	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
)

type StringToByte32Input struct {
	Input string `json:"Input"`
}

type EncodeCreateDIDParamInput struct {
	Address       string                     `json:"Address"`
	DIDParam      rootfactory.CreateDIDParam `json:"DIDParam"`
	KeyIdentifier string                     `json:"KeyIdentifier"`
}
