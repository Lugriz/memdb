package errors

import "errors"

var (
	ErrInvalidDataType           = errors.New("invalid data type")
	ErrInvalidOperation          = errors.New("invalid operation")
	ErrInvalidValueType          = errors.New("invalid value type")
	ErrInvalidValueTypeOperation = errors.New("invalid value type operation")
	ErrInvalidKey                = errors.New("invalid key")
	ErrUnknownKey                = errors.New("unknown key")
)
