package maps_test

import (
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/shared/maps"
)

func TestInvert(t *testing.T) {
	tests := []struct {
		Name   string
		Input  map[string]string
		Result map[string]string
	}{
		{
			Name: "returns inverted map",
			Input: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			Result: map[string]string{
				"value1": "key1",
				"value2": "key2",
				"value3": "key3",
			},
		},
		{
			Name:   "returns empty map",
			Input:  map[string]string{},
			Result: map[string]string{},
		},
		{
			Name:   "returns nil map",
			Input:  nil,
			Result: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result := maps.Invert(tt.Input)

			if !reflect.DeepEqual(result, tt.Result) {
				t.Errorf("want %v, got %v", tt.Result, result)
			}
		})
	}
}
