package mocks

import "github.com/Lugriz/memdb/internal/domain"

type MockPersistence struct {
	SpySet    *Spy
	SpyGet    *Spy
	SpyDelete *Spy
}

var _ domain.Persistence = &MockPersistence{}

func (m *MockPersistence) Set(key string, value domain.Value) {
	m.SpySet.Called = true
}

func (m *MockPersistence) Get(key string) (domain.Value, bool) {
	r := m.SpyGet.Returns

	return r[0].(domain.Value), r[1].(bool)
}

func (m *MockPersistence) Delete(key string) bool {
	r := m.SpyDelete.Returns
	return r[0].(bool)
}
