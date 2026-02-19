package domain

type OperationType int

const (
	SET OperationType = iota
	GET
	DEL
)

var operationStrings = map[OperationType]string{
	SET: "SET",
	GET: "GET",
	DEL: "DEL",
}

func (o OperationType) String() string {
	op, ok := operationStrings[o]
	if ok {
		return op
	}

	return ""
}

type ReadOperationResult struct {
	Value any
}

type WriteOperationResult struct {
	AffectedKeys int
}

type OperationResult struct {
	Read  *ReadOperationResult
	Write *WriteOperationResult
}

type OperationHandler func(persistence Persistence, key string, args [][]byte) (OperationResult, error)
