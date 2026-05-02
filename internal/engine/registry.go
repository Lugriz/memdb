package engine

import (
	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/operations"
)

var DataTypeRegistryConfig = domain.DataTypeRegistry{
	domain.KEY: domain.OperationRegistry{
		domain.SET: operations.KeySetHandler,
		domain.GET: operations.KeyGetHandler,
		domain.DEL: operations.KeyDelHandler,
	},
}
