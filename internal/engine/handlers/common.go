package handlers

import (
	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/persistence"
)

func SetHandler(store persistence.Persistence, key string, value persistence.Value) (runtime.Result, error) {
	if !isValidValue(value) {
		return runtime.Result{
			Type: runtime.WRITE_RESULT,
		}, errors.ErrInvalidValueType
	}

	if val, ok := store.Get(key); ok && val.DataType != value.DataType {
		return runtime.Result{
			Type: runtime.WRITE_RESULT,
		}, errors.ErrInvalidValueType
	}

	store.Set(key, value)

	return runtime.Result{
		Type: runtime.WRITE_RESULT,
		Write: &runtime.WriteResult{
			AffectedKey: true,
		},
	}, nil
}

func isValidValue(value persistence.Value) bool {
	ok := false

	switch value.DataType {
	case datatypes.KEY:
		_, ok = value.Data.(string)
	case datatypes.HASH:
		_, ok = value.Data.(map[string]any)
	}

	return ok
}
