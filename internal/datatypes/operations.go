package datatypes

type Operation int

const (
	SET Operation = iota
	GET
	DEL
	ADD
)

var OperationStrings = map[Operation]string{
	SET: "SET",
	GET: "GET",
	DEL: "DEL",
	ADD: "ADD",
}

func (o Operation) String() string {
	op, ok := OperationStrings[o]
	if ok {
		return op
	}

	return ""
}
