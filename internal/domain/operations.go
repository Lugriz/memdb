package domain

import (
	"strings"

	"github.com/Lugriz/memdb/internal/shared/maps"
)

type Operation int

const (
	SET Operation = iota
	GET
	DEL
)

var operationStrings = map[Operation]string{
	SET: "SET",
	GET: "GET",
	DEL: "DEL",
}

var stringToOperations = maps.Invert(operationStrings)

func (o Operation) String() string {
	op, ok := operationStrings[o]
	if ok {
		return op
	}

	return ""
}

func ParseOperation(key string) (Operation, error) {
	op, ok := stringToOperations[strings.ToUpper(strings.TrimSpace(key))]
	if ok {
		return op, nil
	}

	return -1, ErrInvalidOperation
}

type OperationType int

const (
	READ_OPERATION OperationType = iota
	WRITE_OPERATION
)

type ReadOperationResult struct {
	Value any
}

type WriteOperationResult struct {
	AffectedKey bool
}

type OperationResult struct {
	Type  OperationType
	Read  *ReadOperationResult
	Write *WriteOperationResult
}

type OperationHandler func(persistence Persistence, key string, value any) (OperationResult, error)
