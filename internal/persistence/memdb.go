package persistence

import "github.com/Lugriz/memdb/internal/domain"

type Memdb struct {
	kv map[string]string
}

var _ domain.Persistence = &Memdb{}

func (db *Memdb) SetKV(key, value string) {
	db.kv[key] = value
}

func (db *Memdb) GetKV(key string) (string, bool) {
	val, ok := db.kv[key]

	return val, ok
}

func (db *Memdb) DeleteKV(key string) bool {
	if _, ok := db.kv[key]; ok {
		delete(db.kv, key)

		return true
	}

	return false
}

func NewDB() *Memdb {
	return &Memdb{
		kv: make(map[string]string),
	}
}
