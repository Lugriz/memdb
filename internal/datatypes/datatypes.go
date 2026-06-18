package datatypes

type DataType int

const (
	KEY DataType = iota
	HASH
)

var DataTypeStrings = map[DataType]string{
	KEY:  "KEY",
	HASH: "HASH",
}

func (c DataType) String() string {
	cmd, ok := DataTypeStrings[c]
	if ok {
		return cmd
	}

	return ""
}
