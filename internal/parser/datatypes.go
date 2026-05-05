package parser

import (
	"strings"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/shared/maps"
)

var stringToDataType = maps.Invert(domain.DataTypeStrings)

func ParseDataType(key string) (domain.DataType, error) {
	dt, ok := stringToDataType[strings.ToUpper(strings.TrimSpace(key))]
	if ok {
		return dt, nil
	}

	return -1, domain.ErrInvalidDataType
}
