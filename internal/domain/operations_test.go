package domain_test

import (
	"errors"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
)

func TestOperationString(t *testing.T) {
	tests := []struct {
		Name   string
		OpType domain.Operation
		Result string
	}{
		{
			Name:   "Should return GET",
			OpType: domain.GET,
			Result: "GET",
		},
		{
			Name:   "Should return SET",
			OpType: domain.SET,
			Result: "SET",
		},
		{
			Name:   "Should return DEL",
			OpType: domain.DEL,
			Result: "DEL",
		},
		{
			Name:   "Should return an empty string when not existing operation",
			OpType: -1,
			Result: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result := tt.OpType.String()

			if result != tt.Result {
				t.Errorf("want %s, got %s", tt.Result, result)
			}
		})
	}
}

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
			result, err := domain.ParseOperation(tt.Key)

			if tt.ExpectErr && !errors.Is(err, tt.Err) {
				t.Errorf("want %s error, got %s", tt.Err, err)
			}

			if result != tt.Result {
				t.Errorf("want %d, got %d", tt.Result, result)
			}
		})
	}
}
