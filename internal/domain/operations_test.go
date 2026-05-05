package domain_test

import (
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
