package registry

import (
	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/handlers"
)

var DataTypeRegistryConfig = DataTypeRegistry{
	datatypes.KEY: OperationRegistry{
		datatypes.SET: handlers.KeySetHandler,
		datatypes.GET: handlers.GetHandler,
		datatypes.DEL: handlers.DelHandler,
	},
	datatypes.HASH: OperationRegistry{
		datatypes.SET: handlers.HashSetHandler,
		datatypes.ADD: handlers.HashAddHandler,
	},
}
