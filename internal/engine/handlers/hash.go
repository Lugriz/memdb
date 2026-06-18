package handlers

import (
	"maps"

	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/persistence"
)

func HashSetHandler(store persistence.Persistence, key string, value any) (runtime.Result, error) {
	return SetHandler(store, key, persistence.Value{
		DataType: datatypes.HASH,
		Data:     value,
	})
}

func HashAddHandler(store persistence.Persistence, key string, value any) (runtime.Result, error) {
	data, ok := value.(map[string]any)
	if !ok {
		return runtime.Result{
			Type: runtime.WRITE_RESULT,
		}, errors.ErrInvalidValueType
	}

	val, ok := store.Get(key)
	if !ok {
		return runtime.Result{
			Type: runtime.WRITE_RESULT,
		}, errors.ErrUnknownKey
	}

	if val.DataType != datatypes.HASH {
		return runtime.Result{
			Type: runtime.WRITE_RESULT,
		}, errors.ErrInvalidValueTypeOperation
	}

	hash, _ := val.Data.(map[string]any)

	maps.Copy(hash, data)

	store.Set(key, persistence.Value{
		DataType: datatypes.HASH,
		Data:     hash,
	})

	return runtime.Result{
		Type: runtime.WRITE_RESULT,
		Write: &runtime.WriteResult{
			AffectedKey: true,
		},
	}, nil
}
