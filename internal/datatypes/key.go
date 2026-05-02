package datatypes

import (
	"github.com/Lugriz/memdb/internal/domain"
)

func KeySetHandler(store domain.Persistence, key string, args any) (domain.OperationResult, error) {
	data, ok := args.(string)
	if !ok {
		return domain.OperationResult{}, domain.ErrInvalidValueType // FIXME: Invalid type
	}

	store.SetKV(key, domain.Value{
		DataType: domain.KEY,
		Data:     data,
	})

	return domain.OperationResult{
		Write: &domain.WriteOperationResult{
			AffectedKey: true,
		},
	}, nil
}

func KeyGetHandler(store domain.Persistence, key string, _ any) (domain.OperationResult, error) {
	if val, ok := store.GetKV(key); ok {
		return domain.OperationResult{
			Read: &domain.ReadOperationResult{
				Value: val.Data,
			},
		}, nil
	}

	return domain.OperationResult{}, nil
}

func KeyDelHandler(store domain.Persistence, key string, _ any) (domain.OperationResult, error) {
	var writeOp domain.WriteOperationResult

	writeOp.AffectedKey = store.DeleteKV(key)

	return domain.OperationResult{
		Write: &writeOp,
	}, nil
}
