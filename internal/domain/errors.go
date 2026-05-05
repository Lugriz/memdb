package domain

import "errors"

var (
	ErrMissingArgs      = errors.New("missing arguments")
	ErrInvalidDataType  = errors.New("invalid data type")
	ErrInvalidOperation = errors.New("invalid operation")
	ErrInvalidValueType = errors.New("invalid value type")
)
