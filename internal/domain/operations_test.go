package domain_test

import (
	"errors"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
)

func TestOperationTypeString(t *testing.T) {
	tests := []struct {
		Name   string
		OpType domain.OperationType
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

func TestParseOperationType(t *testing.T) {
	tests := []struct {
		Name      string
		Key       string
		Result    domain.OperationType
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
			Name:      "returns an error when operation type does not exist",
			Key:       "INVALID",
			Result:    -1,
			ExpectErr: true,
			Err:       domain.ErrInvalidOperationType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := domain.ParseOperationType(tt.Key)

			if tt.ExpectErr && !errors.Is(err, tt.Err) {
				t.Errorf("want %s error, got %s", tt.Err, err)
			}

			if result != tt.Result {
				t.Errorf("want %d, got %d", tt.Result, result)
			}
		})
	}
}
