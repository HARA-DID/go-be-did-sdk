package helperhandler

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/meQlause/go-be-did/internal/validator"
	"github.com/meQlause/go-be-did/pkg/response"
	backendutils "github.com/meQlause/go-be-did/utils"
	"github.com/meQlause/hara-core-blockchain-lib/utils"

	helperdto "github.com/meQlause/go-be-did/internal/domain/dtos/helper"
	accountabstractionsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/accountabstraction"
	didaliassdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didalias"
	didrootsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didroot"
	didvcsdk "github.com/meQlause/go-be-did/internal/infrastructure/sdk/didvc"
	helperuc "github.com/meQlause/go-be-did/internal/usecase/helper"
)

type HelperHandler struct {
	uc *helperuc.HelperUseCase
}

func NewHelperHandler(uc *helperuc.HelperUseCase) *HelperHandler {
	return &HelperHandler{uc: uc}
}

// StringToByte32 godoc
// @Summary      Convert String to Byte32
// @Description  Converts a string input to a 32-byte array (byte32). This is a utility endpoint commonly used for encoding data that needs to fit into a single bytes32 storage slot in smart contracts.
// @Description
// @Description  ## Response Structure
// @Description  Success responses (HTTP 200) contain:
// @Description  - `success` (boolean): Always true for HTTP 200
// @Description  - `data` (string): The hex-encoded 32-byte array result
// @Description  - `meta` (object): Contains timestamp and API version
// @Description
// @Description  ## Common Use Cases
// @Description  - Encoding strings for smart contract storage
// @Description  - Preparing data for keccak256 hashing
// @Description  - Converting identifiers to fixed-size byte arrays
// @Description
// @Description  ## Important Notes
// @Description  - Input strings longer than 32 bytes will be truncated
// @Description  - Shorter strings will be padded with zeros
// @Description  - The result is returned as a hex string with 0x prefix
// @Tags         helper General
// @Accept       json
// @Produce      json
// @Param        request body helperdto.StringToHex32DTO true "String input to convert to byte32"
// @Success      200 {object} response.Response{data=string} "Successfully converted string to byte32"
// @Failure      400 {object} response.Response "Invalid request body - malformed JSON or missing required fields"
// @Failure      500 {object} response.Response "Internal server error - conversion failed"
// @Router       /helper/string-2-hex32 [post]
func (hh *HelperHandler) StringToHex32(c *fiber.Ctx) error {
	var input helperdto.StringToHex32DTO
	if err := c.BodyParser(&input); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := validator.Validate.Struct(&input); err != nil {
		validationErrors := validator.FormatError(err)
		return response.Error(c, fiber.StatusBadRequest, validationErrors)
	}

	stringToHex32Input := input.Into()
	resp := hh.uc.StringToHex32(stringToHex32Input)
	return response.Success(c, resp)
}

func (hh *HelperHandler) parseAndValidate(c *fiber.Ctx, input any) any {
	if err := c.BodyParser(&input); err != nil {
		return err.Error()
	}

	if err := validator.Validate.Struct(input); err != nil {
		validationErrors := validator.FormatError(err)
		return validationErrors
	}

	return nil
}

func (hh *HelperHandler) buildHelperResponse(address utils.Address, encodedData string, details response.Details) (HelperResponse, error) {

	key := new(big.Int).SetUint64(0)

	var target string
	var err error

	switch details.Service {
	case backendutils.ServiceDIDRoot:
		target = didrootsdk.GetDIDRootSDK().RootFactory.Address.Hex()

	case backendutils.ServiceDIDAlias:
		target = didaliassdk.GetDIDAliasSDK().AliasFactory.Address.Hex()

	case backendutils.ServiceDIDVC:
		target = didvcsdk.GetDIDVCSDK().VCFactory.Address.Hex()

	default:
		return HelperResponse{}, fmt.Errorf("invalid service type: %d", details.Service)
	}

	nonceBI, err := accountabstractionsdk.GetAccountAbstractionSDK().EntryPoint.GetNonce(
		context.Background(),
		address,
		key,
	)
	if err != nil {
		return HelperResponse{}, err
	}

	nonce := nonceBI.Uint64()
	nonceStr := strconv.FormatUint(nonce, 10)

	return HelperResponse{
		Data:    encodedData,
		Target:  target,
		Nonce:   nonceStr,
		Details: details,
	}, nil
}
