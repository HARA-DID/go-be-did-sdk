package helperhandler

import (
	"context"
	"math/big"

	"github.com/gofiber/fiber/v2"

	helperdo "github.com/meQlause/go-be-did/internal/domain/helper"
	accountabstractionsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	helperuc "github.com/meQlause/go-be-did/internal/usecase/helper"

	"github.com/meQlause/go-be-did/pkg/response"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

type HelperHandler struct {
	uc *helperuc.HelperUseCase
}

func NewHelperHandler(uc *helperuc.HelperUseCase) *HelperHandler {
	return &HelperHandler{uc: uc}
}

func (hh *HelperHandler) StringToByte32(c *fiber.Ctx) error {
	var input helperdo.StringToByte32Input
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	resp := hh.uc.StringToByte32(input)
	return response.Success(c, resp)
}

func (hh *HelperHandler) EncodeCreateDIDParam(c *fiber.Ctx) error {
	var input helperdo.EncodeCreateDIDParamInput
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	decodedData, err := hh.uc.EncodeCreateDIDParam(input)
	if err != nil {
		response.Error(c, fiber.StatusInternalServerError, "Can Not Get Nonce")
	}

	walletAddress := utils.HexToAddress(input.Address)

	key := new(big.Int).SetInt64(0)
	nonce, err := accountabstractionsdk.GetAccountAbstractionSDK().EntryPoint.GetNonce(
		context.Background(),
		walletAddress,
		key,
	)
	if err != nil {
		response.Error(c, fiber.StatusInternalServerError, "Can Not Get Nonce")
	}

	nonceValue := nonce.Uint64() & ((1 << 64) - 1)

	resp := EncodeCreateDIDParamResponse{
		Data:   decodedData,
		Target: didrootsdk.GetDIDRootSDK().RootFactory.Address.Hex(),
		Nonce:  nonceValue,
	}

	return response.Success(c, resp)
}
