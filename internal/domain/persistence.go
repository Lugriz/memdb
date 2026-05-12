package domain

type Value struct {
	DataType DataType
	Data     any
}

type Persistence interface {
	Set(key string, value Value)
	Get(key string) (Value, bool)
	Delete(key string) bool
}
