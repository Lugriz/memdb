package datatypes_test

import (
	"testing"

	"github.com/Lugriz/memdb/internal/datatypes"
)

func TestOperationString(t *testing.T) {
	tests := []struct {
		Name   string
		OpType datatypes.Operation
		Result string
	}{
		{
			Name:   "Should return GET",
			OpType: datatypes.GET,
			Result: "GET",
		},
		{
			Name:   "Should return GET_ALL",
			OpType: datatypes.GET_ALL,
			Result: "GET_ALL",
		},
		{
			Name:   "Should return SET",
			OpType: datatypes.SET,
			Result: "SET",
		},
		{
			Name:   "Should return DEL",
			OpType: datatypes.DEL,
			Result: "DEL",
		},
		{
			Name:   "Should return DEL_ALL",
			OpType: datatypes.DEL_ALL,
			Result: "DEL_ALL",
		},
		{
			Name:   "Should return ADD",
			OpType: datatypes.ADD,
			Result: "ADD",
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
