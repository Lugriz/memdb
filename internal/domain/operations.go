package domain

type Operation int

const (
	SET Operation = iota
	GET
	DEL
)

var OperationStrings = map[Operation]string{
	SET: "SET",
	GET: "GET",
	DEL: "DEL",
}

func (o Operation) String() string {
	op, ok := OperationStrings[o]
	if ok {
		return op
	}

	return ""
}

type OperationType int

const (
	READ_OPERATION OperationType = iota
	WRITE_OPERATION
)

type ReadOperationResult struct {
	Value any
}

type WriteOperationResult struct {
	AffectedKey bool
}

type OperationResult struct {
	Type  OperationType
	Read  *ReadOperationResult
	Write *WriteOperationResult
}

type OperationHandler func(persistence Persistence, key string, value any) (OperationResult, error)
