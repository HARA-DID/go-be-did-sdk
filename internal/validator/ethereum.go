package validator

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-playground/validator/v10"
	"github.com/meQlause/hara-core-blockchain-lib/utils"
)

func registerEthereum(v *validator.Validate) {
	v.RegisterValidation("eth_private_key", ethPrivateKey)
	v.RegisterValidation("eth_public_key", ethPublicKey)
	v.RegisterValidation("eth_address", ethAddress)
	v.RegisterValidation("hex32", hex32Bytes)
	v.RegisterValidation("hex_bigint", hexBigInt)
	v.RegisterValidation("hex_data", hexData)
	v.RegisterValidation("uint64", uint64String)
}

func ethPrivateKey(fl validator.FieldLevel) bool {
	key := strings.TrimPrefix(fl.Field().String(), "0x")

	if len(key) != 64 {
		return false
	}

	bytes, err := utils.DecodeString(key)
	if err != nil {
		return false
	}

	_, err = crypto.ToECDSA(bytes)
	return err == nil
}

func uint64String(fl validator.FieldLevel) bool {
	value := strings.TrimSpace(fl.Field().String())

	if value == "" {
		return false
	}

	_, err := strconv.ParseUint(value, 10, 64)
	return err == nil
}

func ethPublicKey(fl validator.FieldLevel) bool {
	key := strings.TrimPrefix(fl.Field().String(), "0x")

	bytes, err := utils.DecodeString(key)
	return err == nil && len(bytes) == 64
}

func ethAddress(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	return common.IsHexAddress(address)
}

func hex32Bytes(fl validator.FieldLevel) bool {
	value := strings.TrimPrefix(fl.Field().String(), "0x")

	// Must be exactly 64 hex characters (32 bytes)
	if len(value) != 64 {
		return false
	}

	// Must be valid hex
	_, err := utils.DecodeString(value)
	return err == nil
}

func hexBigInt(fl validator.FieldLevel) bool {
	value := strings.TrimPrefix(fl.Field().String(), "0x")

	// Must be valid hex
	if _, err := utils.DecodeString(value); err != nil {
		return false
	}

	// Try to parse as big.Int
	bigInt := new(big.Int)
	_, ok := bigInt.SetString(value, 16)
	return ok
}

func hexData(fl validator.FieldLevel) bool {
	value := strings.TrimPrefix(fl.Field().String(), "0x")

	// Empty hex is valid (represents empty bytes)
	if value == "" {
		return true
	}

	// Must be even length (each byte is 2 hex chars)
	if len(value)%2 != 0 {
		return false
	}

	// Must be valid hex
	_, err := utils.DecodeString(value)
	return err == nil
}
