package persistence

type Store interface {
	SetKV(key, value string)
	GetKV(key string) (string, bool)
	DeleteKV(key string) bool
}
