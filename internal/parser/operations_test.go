package parser_test

import (
	"errors"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/parser"
)

func TestParseOperation(t *testing.T) {
	tests := []struct {
		Name      string
		Key       string
		Result    domain.Operation
		ExpectErr bool
		Err       error
	}{
		{
			Name:   "returns a operation type with upper case key",
			Key:    "SET",
			Result: domain.SET,
		},
		{
			Name:   "returns a operation type with lower case key",
			Key:    "get",
			Result: domain.GET,
		},
		{
			Name:   "Should return DEL even with whitespaces",
			Key:    "  DEL  ",
			Result: domain.DEL,
		},
		{
			Name:      "returns an error when operation type does not exist",
			Key:       "INVALID",
			Result:    -1,
			ExpectErr: true,
			Err:       domain.ErrInvalidOperation,
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
