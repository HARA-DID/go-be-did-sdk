package dto

// HandleOpsInputDTO represents the DTO for HandleOps request (used for Swagger documentation and request parsing)
// @Description HandleOps request payload with private key, wallet address, target address, data, and nonce
type HandleOpsInputDTO struct {
	PrivKey string `json:"priv_key" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef" binding:"required"`
	Data    string `json:"data" example:"0x..." binding:"required"`
	Target  string `json:"target" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb" binding:"required"`
	Nonce   string `json:"nonce" example:"0x0" binding:"required"`
	Wallet  string `json:"wallet" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb" binding:"required"`
}

// UserOpDTO represents the DTO for UserOp (used for Swagger documentation and request parsing)
// @Description UserOp object containing target address, value, data, client block number, user nonce, and signature
type UserOpDTO struct {
	Target            string `json:"target" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Value             string `json:"value" example:"0x0"`
	Data              string `json:"data" example:"0x..."`
	ClientBlockNumber string `json:"client_block_number" example:"0x12345"`
	UserNonce         string `json:"user_nonce" example:"0x0"`
	Signature         string `json:"signature" example:"0x..."`
}

// ValidateUserOpsInputDTO represents the DTO for ValidateUserOps request (used for Swagger documentation and request parsing)
// @Description Validation payload with wallet address and UserOp object
type ValidateUserOpsInputDTO struct {
	Wallet string    `json:"wallet" example:"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"`
	Input  UserOpDTO `json:"input"`
}

// DeployWalletParamsDTO represents the DTO for DeployWalletParams (used for Swagger documentation and request parsing)
// @Description Wallet deployment parameters with owners and salt
type DeployWalletParamsDTO struct {
	Owners []string `json:"owners" example:"0x111...,0x222..."`
	Salt   string   `json:"salt" example:"0xabc... (32 bytes hex)"`
}

// CreateWalletInputDTO represents the DTO for CreateWallet request (used for Swagger documentation and request parsing)
// @Description Wallet creation payload with deployer address and optional salt value
type CreateWalletInputDTO struct {
	PrivKey string                `json:"priv_key" example:"0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"`
	Input   DeployWalletParamsDTO `json:"input"`
}
