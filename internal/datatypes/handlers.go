package datatypes

import (
	"github.com/Lugriz/memdb/internal/persistence"
)

func KeySetHandler(store persistence.Store, key string, args [][]byte) (OperationResult, error) {
	if len(args) < 1 {
		return OperationResult{}, ErrMissingArgs
	}

	value := string(args[0])

	store.SetKV(key, value)

	return OperationResult{
		Write: &WriteOperationResult{
			AffectedKeys: 1,
		},
	}, nil
}

func KeyGetHandler(store persistence.Store, key string, _ [][]byte) (OperationResult, error) {
	if val, ok := store.GetKV(key); ok {
		return OperationResult{
			Read: &ReadOperationResult{
				Value: val,
			},
		}, nil
	}

	return OperationResult{}, nil
}

func KeyDelHandler(store persistence.Store, key string, _ [][]byte) (OperationResult, error) {
	if ok := store.DeleteKV(key); ok {
		return OperationResult{
			Write: &WriteOperationResult{
				AffectedKeys: 1,
			},
		}, nil
	}

	return OperationResult{}, nil
}
