package datatypes

type DataType int

const (
	KEY DataType = iota
)

var dataTypeStrings = map[DataType]string{
	KEY: "KEY",
}

func (c DataType) String() string {
	cmd, ok := dataTypeStrings[c]
	if ok {
		return cmd
	}

	return ""
}
