package persistence

import (
	"sync"
)

type InMemoryDB struct {
	kv sync.Map
}

var _ Persistence = &InMemoryDB{}

func (db *InMemoryDB) Set(key string, value Value) {
	db.kv.Store(key, value)
}

func (db *InMemoryDB) Get(key string) (Value, bool) {
	val, ok := db.kv.Load(key)
	if ok {
		return val.(Value), true
	}

	return Value{}, false
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
