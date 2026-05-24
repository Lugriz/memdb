package runtime

type ResultType int

const (
	READ_RESULT ResultType = iota
	WRITE_RESULT
)

type ReadResult struct {
	Value any
}

type WriteResult struct {
	AffectedKey bool
}

type Result struct {
	Type  ResultType
	Read  *ReadResult
	Write *WriteResult
}
