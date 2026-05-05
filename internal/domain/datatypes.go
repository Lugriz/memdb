package domain

type DataType int

const (
	KEY DataType = iota
)

var DataTypeStrings = map[DataType]string{
	KEY: "KEY",
}

func (c DataType) String() string {
	cmd, ok := DataTypeStrings[c]
	if ok {
		return cmd
	}

	return ""
}
