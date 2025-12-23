package accountabstractionhandler

type TxCheckResponse struct {
	Success  map[string]bool                 `json:"success"`
	Errors   map[string]string               `json:"errors"`
	Returned map[string]*WalletDeployedEvent `json:"returned"`
}
