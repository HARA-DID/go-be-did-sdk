package didaliashandler

import (
	didaliasusecase "github.com/meQlause/go-be-did/internal/usecase/didalias"
)

type DIDAliasHandler struct {
	uc *didaliasusecase.DIDAliasUseCase
}

func NewDIDAliasHandler(uc *didaliasusecase.DIDAliasUseCase) *DIDAliasHandler {
	return &DIDAliasHandler{uc: uc}
}
