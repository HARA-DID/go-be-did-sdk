package helperhandler

type EncodeCreateDIDParamResponse struct {
	Data   string `json:"data" example:"0x..."`
	Target string `json:"target" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Nonce  uint64 `json:"nonce" example:"1"`
}
