package repository

import (
	"context"

	aa "github.com/meQlause/go-be-did/internal/domain/accountabstraction"
	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
)

type AccountAbstractionRepository interface {
	CreateWallet(ctx context.Context, input aa.CreateWalletInput) (*aa.TxHash, error)
	ExecuteOperation(ctx context.Context, input aa.ExecuteOperationInput) (*aa.TxHash, error)
}

type HelperRepository interface {
	StringToByte32(input helperdo.StringToByte32Input) [32]byte
	EncodeCreateDIDParam(createDIDParam helperdo.EncodeCreateDIDParamInput) string
}
