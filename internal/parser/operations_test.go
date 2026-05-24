package parser_test

import (
	"errors"
	"testing"

	"github.com/Lugriz/memdb/internal/datatypes"
	appErrors "github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/parser"
)

func TestParseOperation(t *testing.T) {
	tests := []struct {
		Name      string
		Key       string
		Result    datatypes.Operation
		ExpectErr bool
		Err       error
	}{
		{
			Name:   "returns a operation type with upper case key",
			Key:    "SET",
			Result: datatypes.SET,
		},
		{
			Name:   "returns a operation type with lower case key",
			Key:    "get",
			Result: datatypes.GET,
		},
		{
			Name:   "Should return DEL even with whitespaces",
			Key:    "  DEL  ",
			Result: datatypes.DEL,
		},
		{
			Name:      "returns an error when operation type does not exist",
			Key:       "INVALID",
			Result:    -1,
			ExpectErr: true,
			Err:       appErrors.ErrInvalidOperation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := parser.ParseOperation(tt.Key)

			if tt.ExpectErr && !errors.Is(err, tt.Err) {
				t.Errorf("want %s error, got %s", tt.Err, err)
			}

			if result != tt.Result {
				t.Errorf("want %d, got %d", tt.Result, result)
			}
		})
	}
}
