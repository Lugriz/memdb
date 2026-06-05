package handlers_test

import (
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/handlers"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/mocks"
	"github.com/Lugriz/memdb/internal/persistence"
)

func TestKeyGetHandler(t *testing.T) {
	tests := []struct {
		Name        string
		Persistence *mocks.MockPersistence
		Key         string
		Value       any
		Result      runtime.Result
	}{
		{
			Name: "returns a value",
			Persistence: &mocks.MockPersistence{
				SpyGet: &mocks.Spy{
					Returns: []any{
						persistence.Value{
							DataType: datatypes.KEY,
							Data:     "value1",
						},
						true,
					},
				},
			},
			Key: "key1",
			Result: runtime.Result{
				Type: runtime.READ_RESULT,
				Read: &runtime.ReadResult{
					Value: "value1",
				},
			},
		},
		{
			Name: "not found value",
			Persistence: &mocks.MockPersistence{
				SpyGet: &mocks.Spy{
					Returns: []any{persistence.Value{}, false},
				},
			},
			Key: "key1",
			Result: runtime.Result{
				Type: runtime.READ_RESULT,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := handlers.KeyGetHandler(tt.Persistence, tt.Key, tt.Value)

			if err != nil {
				t.Errorf("Got %s err, Want %v", err, nil)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %s result, Want %s", toString(result), toString(tt.Result))
			}
		})
	}
}

func TestKeyDelHandler(t *testing.T) {
	tests := []struct {
		Name        string
		Persistence *mocks.MockPersistence
		Key         string
		Value       any
		Result      runtime.Result
	}{
		{
			Name: "Delete the key",
			Persistence: &mocks.MockPersistence{
				SpyDelete: &mocks.Spy{
					Returns: []any{true},
				},
			},
			Key: "key1",
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
				Write: &runtime.WriteResult{
					AffectedKey: true,
				},
			},
		},
		{
			Name: "Unexisting key",
			Persistence: &mocks.MockPersistence{
				SpyDelete: &mocks.Spy{
					Returns: []any{false},
				},
			},
			Key: "key1",
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
				Write: &runtime.WriteResult{
					AffectedKey: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := handlers.KeyDelHandler(tt.Persistence, tt.Key, tt.Value)

			if err != nil {
				t.Errorf("Got %s err, Want %v", err, nil)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %s result, Want %s", toString(result), toString(tt.Result))
			}
		})
	}
}
