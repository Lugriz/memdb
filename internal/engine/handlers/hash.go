package handlers

import (
	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/persistence"
)

func HashSetHandler(store persistence.Persistence, key string, value any) (runtime.Result, error) {
	return SetHandler(store, key, persistence.Value{
		DataType: datatypes.HASH,
		Data:     value,
	})
}
