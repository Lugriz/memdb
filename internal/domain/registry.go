package domain

type DataTypeRegistry map[DataType]OperationRegistry

type OperationRegistry map[OperationType]OperationHandler
