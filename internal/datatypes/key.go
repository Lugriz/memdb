package datatypes

import (
	"github.com/Lugriz/memdb/internal/domain"
)

func KeySetHandler(store domain.Persistence, key string, args [][]byte) (domain.OperationResult, error) {
	if len(args) < 1 {
		return domain.OperationResult{}, domain.ErrMissingArgs
	}

	value := string(args[0])

	store.SetKV(key, value)

	return domain.OperationResult{
		Write: &domain.WriteOperationResult{
			AffectedKeys: 1,
		},
	}, nil
}

func KeyGetHandler(store domain.Persistence, key string, _ [][]byte) (domain.OperationResult, error) {
	if val, ok := store.GetKV(key); ok {
		return domain.OperationResult{
			Read: &domain.ReadOperationResult{
				Value: val,
			},
		}, nil
	}

	return domain.OperationResult{}, nil
}

func KeyDelHandler(store domain.Persistence, key string, _ [][]byte) (domain.OperationResult, error) {
	var writeOp domain.WriteOperationResult

	if ok := store.DeleteKV(key); ok {
		writeOp.AffectedKeys = 1

		return domain.OperationResult{
			Write: &writeOp,
		}, nil
	}

	return domain.OperationResult{
		Write: &writeOp,
	}, nil
}
