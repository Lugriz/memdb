package domain

import "errors"

var (
	ErrMissingArgs          = errors.New("missing arguments")
	ErrInvalidDataType      = errors.New("invalid data type")
	ErrInvalidOperationType = errors.New("invalid operation type")
	ErrInternal             = errors.New("internal error")
)
