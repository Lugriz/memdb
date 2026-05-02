package domain

type Persistence interface {
	SetKV(key string, value Value)
	GetKV(key string) (Value, bool)
	DeleteKV(key string) bool
}
