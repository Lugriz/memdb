package domain

import (
	"strings"

	"github.com/Lugriz/memdb/internal/shared/maps"
)

type OperationType int

const (
	SET OperationType = iota
	GET
	DEL
)

var operationStrings = map[OperationType]string{
	SET: "SET",
	GET: "GET",
	DEL: "DEL",
}

var stringToOperations = maps.Invert(operationStrings)

func (o OperationType) String() string {
	op, ok := operationStrings[o]
	if ok {
		return op
	}

	return ""
}

func ParseOperationType(key string) (OperationType, error) {
	op, ok := stringToOperations[strings.ToUpper(key)]
	if ok {
		return op, nil
	}

	return -1, ErrInvalidOperationType
}

type ReadOperationResult struct {
	Value any
}

type WriteOperationResult struct {
	AffectedKey bool
}

type OperationResult struct {
	Read  *ReadOperationResult
	Write *WriteOperationResult
}

type OperationHandler func(persistence Persistence, key string, args []string) (OperationResult, error)
