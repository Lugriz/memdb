package domain

import "github.com/Lugriz/memdb/internal/shared/maps"

type DataType int

const (
	KEY DataType = iota
)

var dataTypeStrings = map[DataType]string{
	KEY: "KEY",
}

var stringToDataType = maps.Invert(dataTypeStrings)

func (c DataType) String() string {
	cmd, ok := dataTypeStrings[c]
	if ok {
		return cmd
	}

	return ""
}

func ParseDataType(key string) (DataType, error) {
	dt, ok := stringToDataType[key]
	if ok {
		return dt, nil
	}

	return 0, ErrInvalidDataType
}
