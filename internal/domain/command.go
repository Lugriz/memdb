package domain

type Command struct {
	DataType  DataType
	Key       string
	Operation Operation
	Value     any
}
