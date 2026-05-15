package mocks

import "github.com/Lugriz/memdb/internal/domain"

type MockEngine struct {
	SpyExecute *Spy
}

var _ domain.Engine = &MockEngine{}

func (m *MockEngine) Execute(command *domain.Command) (domain.OperationResult, error) {
	result := m.SpyExecute.Returns[0].(domain.OperationResult)
	err := m.SpyExecute.Returns[1]
	if err == nil {
		return result, nil
	}

	return result, err.(error)
}
