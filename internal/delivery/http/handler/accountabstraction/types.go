package accountabstractionhandler

import aaevent "github.com/meQlause/go-be-did/internal/delivery/event/accountabstraction"

type Response struct {
	Success  bool                         `json:"success"`
	Errors   string                       `json:"errors"`
	Returned *aaevent.WalletDeployedEvent `json:"returned"`
}
