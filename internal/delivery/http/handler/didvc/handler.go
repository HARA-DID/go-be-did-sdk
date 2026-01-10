package didvchandler

import (
	didvccase "github.com/meQlause/go-be-did/internal/usecase/didvc"
)

type DIDVCHandler struct {
	uc *didvccase.DIDVCUseCase
}

func NewDIDVCHandler(uc *didvccase.DIDVCUseCase) *DIDVCHandler {
	return &DIDVCHandler{uc: uc}
}
