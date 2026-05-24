package datatypes

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
