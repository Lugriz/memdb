package parser

import (
	"strings"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/shared/maps"
)

var stringToOperations = maps.Invert(domain.OperationStrings)

func ParseOperation(key string) (domain.Operation, error) {
	op, ok := stringToOperations[strings.ToUpper(strings.TrimSpace(key))]
	if ok {
		return op, nil
	}

	return -1, domain.ErrInvalidOperation
}
