package domain

type Value struct {
	DataType DataType
	Data     any
}

type Persistence interface {
	SetKV(key string, value Value)
	GetKV(key string) (Value, bool)
	DeleteKV(key string) bool
}
