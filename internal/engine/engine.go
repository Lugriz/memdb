package engine

import (
	"github.com/Lugriz/memdb/internal/domain"
)

type Engine struct {
	persistence      domain.Persistence
	dataTypeRegistry domain.DataTypeRegistry
}

func (e *Engine) Execute(command *domain.Command) (domain.OperationResult, error) {
	opRegistry, ok := e.dataTypeRegistry[command.DataType]
	if !ok {
		return domain.OperationResult{}, domain.ErrInvalidDataType
	}

	handler, ok := opRegistry[command.Operation]
	if !ok {
		return domain.OperationResult{}, domain.ErrInvalidOperation
	}

	return handler(e.persistence, command.Key, command.Value)
}

func NewEngine(persistence domain.Persistence, dataTypeRegistry domain.DataTypeRegistry) *Engine {
	return &Engine{
		persistence:      persistence,
		dataTypeRegistry: dataTypeRegistry,
	}
}
