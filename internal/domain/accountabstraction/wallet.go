package accountabstractiondomain

type Wallet struct {
	Address string
	Owner   string
	Nonce   int64
}

type WalletInfo struct {
	Address string         `json:"address"`
	Owner   string         `json:"owner"`
	Nonce   int64          `json:"nonce"`
	Balance string         `json:"balance"`
	Data    map[string]any `json:"data"`
}
