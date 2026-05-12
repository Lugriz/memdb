package persistence

import (
	"sync"

	"github.com/Lugriz/memdb/internal/domain"
)

type InMemoryDB struct {
	kv sync.Map
}

var _ domain.Persistence = &InMemoryDB{}

func (db *InMemoryDB) Set(key string, value domain.Value) {
	db.kv.Store(key, value)
}

func (db *InMemoryDB) Get(key string) (domain.Value, bool) {
	val, ok := db.kv.Load(key)
	if ok {
		return val.(domain.Value), true
	}

	return domain.Value{}, false
}

func (db *InMemoryDB) Delete(key string) bool {
	if _, ok := db.kv.Load(key); ok {
		db.kv.Delete(key)
		return true
	}

	return false
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		kv: sync.Map{},
	}
}
