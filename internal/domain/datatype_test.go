package domain_test

import (
	"errors"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
)

func TestDataTypeString(t *testing.T) {
	tests := []struct {
		Name     string
		DataType domain.DataType
		Result   string
	}{
		{
			Name:     "Should return KEY",
			DataType: domain.KEY,
			Result:   "KEY",
		},
		{
			Name:     "Should return an empty string when not existing data type",
			DataType: -1,
			Result:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result := tt.DataType.String()

			if result != tt.Result {
				t.Errorf("want %s, got %s", tt.Result, result)
			}
		})
	}
}

func TestParseDataType(t *testing.T) {
	tests := []struct {
		Name      string
		Key       string
		Result    domain.DataType
		ExpectErr bool
		Err       error
	}{
		{
			Name:   "returns a operation type with upper case key",
			Key:    "KEY",
			Result: domain.KEY,
		},
		{
			Name:   "returns a operation type with lower case key",
			Key:    "key",
			Result: domain.KEY,
		},
		{
			Name:      "returns an error when data type does not exist",
			Key:       "INVALID",
			Result:    -1,
			ExpectErr: true,
			Err:       domain.ErrInvalidDataType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := domain.ParseDataType(tt.Key)

			if tt.ExpectErr && !errors.Is(err, tt.Err) {
				t.Errorf("want %s error, got %s", tt.Err, err)
			}

			if result != tt.Result {
				t.Errorf("want %d, got %d", tt.Result, result)
			}
		})
	}
}
