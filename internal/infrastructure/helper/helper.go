package internalhelper

import (
	helperdo "github.com/meQlause/go-be-did/internal/domain/entities/helper"

	"github.com/meQlause/go-be-did/internal/repository"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type InternalHelper struct {
}

func NewInternalHelper() (repository.HelperRepository, error) {
	return &InternalHelper{}, nil
}

func (h *InternalHelper) StringToHex32(input helperdo.StringToHex32Input) string {
	return utils.StringToHex32(input.Input)
}
