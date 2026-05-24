package engine

import "github.com/Lugriz/memdb/internal/datatypes"

type Command struct {
	DataType  datatypes.DataType
	Key       string
	Operation datatypes.Operation
	Value     any
}
