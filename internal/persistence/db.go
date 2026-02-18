package persistence

type DB struct {
	kv map[string]string
}

var _ Store = &DB{}

func (db *DB) SetKV(key, value string) {
	db.kv[key] = value
}

func (db *DB) GetKV(key string) (string, bool) {
	val, ok := db.kv[key]

	return val, ok
}

func (db *DB) DeleteKV(key string) bool {
	if _, ok := db.kv[key]; ok {
		delete(db.kv, key)

		return true
	}

	return false
}

func NewDB() *DB {
	return &DB{
		kv: make(map[string]string),
	}
}
