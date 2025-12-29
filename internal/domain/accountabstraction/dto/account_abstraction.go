package dto

type CreateWalletRequest struct {
	PrivKey string                    `json:"privKey" example:"0xabc123..."`
	Input   DeployWalletParamsRequest `json:"input"`
}

type DeployWalletParamsRequest struct {
	Owners []string `json:"owners" example:"0x111...,0x222..."`
	Salt   string   `json:"salt" example:"0xaaaaaaaa...(32 bytes hex)"`
}
