package registry

import (
	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/handlers"
)

var DataTypeRegistryConfig = DataTypeRegistry{
	datatypes.KEY: OperationRegistry{
		datatypes.SET: handlers.KeySetHandler,
		datatypes.GET: handlers.KeyGetHandler,
		datatypes.DEL: handlers.KeyDelHandler,
	},
}
