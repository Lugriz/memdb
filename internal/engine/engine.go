package engine

import (
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/persistence"
	"github.com/Lugriz/memdb/internal/registry"
)

type Engine interface {
	Execute(*Command) (runtime.Result, error)
}

type engine struct {
	persistence      persistence.Persistence
	dataTypeRegistry registry.DataTypeRegistry
}

func (e *engine) Execute(command *Command) (runtime.Result, error) {
	if command.Key == "" {
		return runtime.Result{}, errors.ErrInvalidKey
	}

	opRegistry, ok := e.dataTypeRegistry[command.DataType]
	if !ok {
		return runtime.Result{}, errors.ErrInvalidDataType
	}

	handler, ok := opRegistry[command.Operation]
	if !ok {
		return runtime.Result{}, errors.ErrInvalidOperation
	}

	return handler(e.persistence, command.Key, command.Value)
}

func NewEngine(persistence persistence.Persistence, dataTypeRegistry registry.DataTypeRegistry) *engine {
	return &engine{
		persistence:      persistence,
		dataTypeRegistry: dataTypeRegistry,
	}
}
