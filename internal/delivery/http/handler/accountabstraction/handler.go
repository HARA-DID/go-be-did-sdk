package accountabstractionhandler

import aauc "github.com/meQlause/go-be-did/internal/usecase/accountabstraction"

type AccountAbstractionHandler struct {
	uc *aauc.AccountAbstactionUseCase
}

func NewAccountAbstractionHandler(uc *aauc.AccountAbstactionUseCase) *AccountAbstractionHandler {
	return &AccountAbstractionHandler{uc: uc}
}
