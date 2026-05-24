package parser

import (
	"strings"

	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/shared/maps"
)

var stringToOperations = maps.Invert(datatypes.OperationStrings)

func ParseOperation(key string) (datatypes.Operation, error) {
	op, ok := stringToOperations[strings.ToUpper(strings.TrimSpace(key))]
	if ok {
		return op, nil
	}

	return -1, errors.ErrInvalidOperation
}
