package internalhelper

import (
	helperdomain "github.com/meQlause/go-be-did/internal/domain/helper"
	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type InternalHelper struct {
}

func NewInternalHelper() (repository.HelperRepository, error) {

	return &InternalHelper{}, nil
}

func (h *InternalHelper) StringToByte32(input helperdomain.StringToByte32Input) [32]byte {
	return utils.StringToByte32(input.Input)
}
