package helperhandler

type EncodeCreateDIDParamResponse struct {
	Data   string `json:"Data"`
	Target string `json:"Target"`
	Nonce  uint64 `json:"Nonce"`
}
