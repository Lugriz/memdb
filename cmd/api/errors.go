package api

import "github.com/Lugriz/memdb/internal/errors"

const (
	InvalidDataTypeError  = "INVALID_DATA_TYPE"
	InvalidValueTypeError = "INVALID_VALUE_TYPE"
	InvalidOperationError = "INVALID_OPERATION"
	InvalidSyntaxError    = "INVALID_SYNTAX"
	InvalidKeyError       = "INVALID_KEY"
	UnknownError          = "UNKNOWN"
)

var errorTypes = map[error]string{
	errors.ErrInvalidDataType:  InvalidDataTypeError,
	errors.ErrInvalidValueType: InvalidValueTypeError,
	errors.ErrInvalidOperation: InvalidOperationError,
	errors.ErrInvalidKey:       InvalidKeyError,
}

func TypeFromError(err error) string {
	if t, ok := errorTypes[err]; ok {
		return t
	}

	return UnknownError
}
