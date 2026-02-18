package datatypes

type DataTypeRegistry map[DataType]OperationRegistry

type OperationRegistry map[OperationType]OperationHandler

var DataTypeRegistryConfig = DataTypeRegistry{
	KEY: OperationRegistry{
		SET: KeySetHandler,
		GET: KeyGetHandler,
		DEL: KeyDelHandler,
	},
}
