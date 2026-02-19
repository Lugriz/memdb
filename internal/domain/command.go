package domain

type Command struct {
	DataType      DataType
	Key           string
	OperationType OperationType
	Args          [][]byte
}
