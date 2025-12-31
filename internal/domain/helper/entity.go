package helperdomain

import (
	"github.com/meQlause/did-root-sdk/pkg/rootfactory"
)

type StringToByte32Input struct {
	Input string `json:"input" example:"example_string"`
}

type EncodeCreateDIDParamInput struct {
	Address       string                     `json:"address" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	DIDParam      rootfactory.CreateDIDParam `json:"did_param"`
	KeyIdentifier string                     `json:"key_identifier" example:"key1"`
}
