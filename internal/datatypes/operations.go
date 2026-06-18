package datatypes

type Operation int

const (
	SET Operation = iota
	GET
	GET_ALL
	DEL
	DEL_ALL
	ADD
)

var OperationStrings = map[Operation]string{
	SET:     "SET",
	GET:     "GET",
	GET_ALL: "GET_ALL",
	DEL:     "DEL",
	DEL_ALL: "DEL_ALL",
	ADD:     "ADD",
}

func (o Operation) String() string {
	op, ok := OperationStrings[o]
	if ok {
		return op
	}

	return ""
}
