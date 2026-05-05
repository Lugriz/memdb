package domain_test

import (
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
