package mocks

import (
	"github.com/Lugriz/memdb/internal/persistence"
)

type MockPersistence struct {
	SpySet    *Spy
	SpyGet    *Spy
	SpyDelete *Spy
}

var _ persistence.Persistence = &MockPersistence{}

func (m *MockPersistence) Set(key string, value persistence.Value) {
	m.SpySet.Called = true
}

func (m *MockPersistence) Get(key string) (persistence.Value, bool) {
	r := m.SpyGet.Returns

	return r[0].(persistence.Value), r[1].(bool)
}

func (m *MockPersistence) Delete(key string) bool {
	r := m.SpyDelete.Returns
	return r[0].(bool)
}
