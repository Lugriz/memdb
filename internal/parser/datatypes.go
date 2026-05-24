package parser

import (
	"strings"

	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/shared/maps"
)

var stringToDataType = maps.Invert(datatypes.DataTypeStrings)

func ParseDataType(key string) (datatypes.DataType, error) {
	dt, ok := stringToDataType[strings.ToUpper(strings.TrimSpace(key))]
	if ok {
		return dt, nil
	}

	return -1, errors.ErrInvalidDataType
}
