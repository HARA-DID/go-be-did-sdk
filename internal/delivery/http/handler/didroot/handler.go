package didroothandler

import (
	didrootusecase "github.com/meQlause/go-be-did/internal/usecase/didroot"
)

type DIDRootHandler struct {
	uc *didrootusecase.DIDRootUseCase
}

func NewDIDRootHandler(uc *didrootusecase.DIDRootUseCase) *DIDRootHandler {
	return &DIDRootHandler{uc: uc}
}
