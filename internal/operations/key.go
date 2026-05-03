package operations

import (
	"github.com/Lugriz/memdb/internal/domain"
)

func KeySetHandler(store domain.Persistence, key string, value any) (domain.OperationResult, error) {
	data, ok := value.(string)
	if !ok {
		return domain.OperationResult{
			Type: domain.WRITE_OPERATION,
		}, domain.ErrInvalidValueType
	}

	store.SetKV(key, domain.Value{
		DataType: domain.KEY,
		Data:     data,
	})

	return domain.OperationResult{
		Type: domain.WRITE_OPERATION,
		Write: &domain.WriteOperationResult{
			AffectedKey: true,
		},
	}, nil
}

func KeyGetHandler(store domain.Persistence, key string, _ any) (domain.OperationResult, error) {
	if val, ok := store.GetKV(key); ok {
		return domain.OperationResult{
			Type: domain.READ_OPERATION,
			Read: &domain.ReadOperationResult{
				Value: val.Data,
			},
		}, nil
	}

	return domain.OperationResult{
		Type: domain.READ_OPERATION,
	}, nil
}

func KeyDelHandler(store domain.Persistence, key string, _ any) (domain.OperationResult, error) {
	var writeOp domain.WriteOperationResult

	writeOp.AffectedKey = store.DeleteKV(key)

	return domain.OperationResult{
		Type:  domain.WRITE_OPERATION,
		Write: &writeOp,
	}, nil
}
