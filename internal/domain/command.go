package domain

type Command struct {
	DataType      DataType
	Key           string
	OperationType OperationType
	Value         any
}
