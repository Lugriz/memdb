package registry

import (
	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/persistence"
)

type Handler func(persistence persistence.Persistence, key string, value any) (runtime.Result, error)

type OperationRegistry map[datatypes.Operation]Handler

type DataTypeRegistry map[datatypes.DataType]OperationRegistry
