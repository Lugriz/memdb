package parser_test

import (
	"errors"
	"testing"

	"github.com/Lugriz/memdb/internal/datatypes"
	appErrors "github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/parser"
)

func TestParseDataType(t *testing.T) {
	tests := []struct {
		Name      string
		Key       string
		Result    datatypes.DataType
		ExpectErr bool
		Err       error
	}{
		{
			Name:   "returns a data type with upper case key",
			Key:    "KEY",
			Result: datatypes.KEY,
		},
		{
			Name:   "returns a data type with lower case key",
			Key:    "key",
			Result: datatypes.KEY,
		},
		{
			Name:   "returns a data type even with whitespaces",
			Key:    "  KEY  ",
			Result: datatypes.KEY,
		},
		{
			Name:      "returns an error when data type does not exist",
			Key:       "INVALID",
			Result:    -1,
			ExpectErr: true,
			Err:       appErrors.ErrInvalidDataType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := parser.ParseDataType(tt.Key)

			if tt.ExpectErr && !errors.Is(err, tt.Err) {
				t.Errorf("want %s error, got %s", tt.Err, err)
			}

			if result != tt.Result {
				t.Errorf("want %d, got %d", tt.Result, result)
			}
		})
	}
}
