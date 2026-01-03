package validator

import "github.com/go-playground/validator/v10"

func FormatError(err error) map[string]string {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = message(e)
	}

	return errors
}

func message(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "field is required"
	case "uint64":
		return "must be a valid unsigned 64-bit integer"
	case "eth_private_key":
		return "invalid ethereum private key (must be 64 hex characters)"
	case "eth_public_key":
		return "invalid ethereum public key (must be 64 hex characters)"
	case "eth_address":
		return "invalid ethereum address"
	case "hex32":
		return "must be 32 bytes in hex format (64 hex characters with optional 0x prefix)"
	case "hex_bigint":
		return "must be a valid hexadecimal big integer (with optional 0x prefix)"
	case "hexadecimal":
		return "must be valid hexadecimal data"
	case "min":
		return "must have at least " + e.Param() + " item(s)"
	case "dive":
		return "invalid array element"
	default:
		return "invalid value"
	}
}
