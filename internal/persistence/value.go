package persistence

import "github.com/Lugriz/memdb/internal/datatypes"

type Value struct {
	DataType datatypes.DataType
	Data     any
}
