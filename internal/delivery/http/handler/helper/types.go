package helperhandler

import backendutils "github.com/meQlause/go-be-did/utils"

type HelperResponse struct {
	Data    string               `json:"data" example:"0x..."`
	Target  string               `json:"target" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Nonce   string               `json:"nonce" example:"1"`
	Details backendutils.Details `json:"details"`
}
