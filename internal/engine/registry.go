package engine

import (
	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/domain"
)

var DataTypeRegistryConfig = domain.DataTypeRegistry{
	domain.KEY: domain.OperationRegistry{
		domain.SET: datatypes.KeySetHandler,
		domain.GET: datatypes.KeyGetHandler,
		domain.DEL: datatypes.KeyDelHandler,
	},
}
