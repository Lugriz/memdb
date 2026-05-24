package mocks

import (
	"github.com/Lugriz/memdb/internal/engine"
	"github.com/Lugriz/memdb/internal/engine/runtime"
)

type MockEngine struct {
	SpyExecute *Spy
}

var _ engine.Engine = &MockEngine{}

func (m *MockEngine) Execute(command *engine.Command) (runtime.Result, error) {
	result := m.SpyExecute.Returns[0].(runtime.Result)
	err := m.SpyExecute.Returns[1]
	if err == nil {
		return result, nil
	}

	return result, err.(error)
}
