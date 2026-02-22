package mocks

import "github.com/Lugriz/memdb/internal/domain"

type MockPersistence struct {
	SpySetKV    *Spy
	SpyGetKV    *Spy
	SpyDeleteKV *Spy
}

var _ domain.Persistence = &MockPersistence{}

func (m *MockPersistence) SetKV(key string, value string) {
	m.SpySetKV.Called = true
}

func (m *MockPersistence) GetKV(key string) (string, bool) {
	r := m.SpyGetKV.Returns

	return r[0].(string), r[1].(bool)
}

func (m *MockPersistence) DeleteKV(key string) bool {
	r := m.SpyDeleteKV.Returns
	return r[0].(bool)
}
