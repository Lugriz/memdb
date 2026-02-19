package domain

type Persistence interface {
	SetKV(key, value string)
	GetKV(key string) (string, bool)
	DeleteKV(key string) bool
}
