package domain

import "errors"

var (
	ErrInvalidDataType  = errors.New("invalid data type")
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidValueType = errors.New("invalid value type")
	ErrInvalidKey       = errors.New("invalid key")
)
