package domain

type DataTypeRegistry map[DataType]OperationRegistry

type OperationRegistry map[Operation]OperationHandler
