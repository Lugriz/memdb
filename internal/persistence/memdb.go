package persistence

import (
	"sync"

	"github.com/Lugriz/memdb/internal/domain"
)

type Memdb struct {
	kv sync.Map
}

var _ domain.Persistence = &Memdb{}

func (db *Memdb) SetKV(key string, value domain.Value) {
	db.kv.Store(key, value)
}

func (db *Memdb) GetKV(key string) (domain.Value, bool) {
	val, ok := db.kv.Load(key)
	if ok {
		return val.(domain.Value), true
	}

	return domain.Value{}, false
}

func (db *Memdb) DeleteKV(key string) bool {
	if _, ok := db.kv.Load(key); ok {
		db.kv.Delete(key)
		return true
	}

	return false
}

func NewMemDB() *Memdb {
	return &Memdb{
		kv: sync.Map{},
	}
}
