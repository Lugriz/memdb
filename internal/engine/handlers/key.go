package handlers

import (
	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/persistence"
)

func KeySetHandler(store persistence.Persistence, key string, value any) (runtime.Result, error) {
	data, ok := value.(string)
	if !ok {
		return runtime.Result{
			Type: runtime.WRITE_RESULT,
		}, errors.ErrInvalidValueType
	}

	store.Set(key, persistence.Value{
		DataType: datatypes.KEY,
		Data:     data,
	})

	return runtime.Result{
		Type: runtime.WRITE_RESULT,
		Write: &runtime.WriteResult{
			AffectedKey: true,
		},
	}, nil
}

func KeyGetHandler(store persistence.Persistence, key string, _ any) (runtime.Result, error) {
	if val, ok := store.Get(key); ok {
		return runtime.Result{
			Type: runtime.READ_RESULT,
			Read: &runtime.ReadResult{
				Value: val.Data,
			},
		}, nil
	}

	return runtime.Result{
		Type: runtime.READ_RESULT,
	}, nil
}

func KeyDelHandler(store persistence.Persistence, key string, _ any) (runtime.Result, error) {
	var writeOp runtime.WriteResult

	writeOp.AffectedKey = store.Delete(key)

	return runtime.Result{
		Type:  runtime.WRITE_RESULT,
		Write: &writeOp,
	}, nil
}
