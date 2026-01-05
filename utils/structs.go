package backendutils

type Details struct {
	Service uint8 `json:"service" example:"0"`
	TxType  uint8 `json:"tx_type" example:"10"`
}

type Response struct {
	Success  bool   `json:"success" example:"true"`
	Errors   string `json:"errors" example:"No Error Message"`
	Returned any    `json:"returned,omitempty"`
}

type TxHash struct {
	TxHash []string
}
