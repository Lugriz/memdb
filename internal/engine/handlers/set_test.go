package handlers_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine/handlers"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	appError "github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/mocks"
	"github.com/Lugriz/memdb/internal/persistence"
)

func TestSetHandler(t *testing.T) {
	tests := []struct {
		Name          string
		Persistence   *mocks.MockPersistence
		ExpectSetCall bool
		Key           string
		Value         persistence.Value
		Result        runtime.Result
		ExpectErr     bool
		Err           error
	}{
		{
			Name: "write key successfully",
			Persistence: &mocks.MockPersistence{
				SpySet: &mocks.Spy{},
				SpyGet: &mocks.Spy{
					Returns: []any{persistence.Value{}, false},
				},
			},
			ExpectSetCall: true,
			Key:           "key1",
			Value: persistence.Value{
				DataType: datatypes.HASH,
				Data: map[string]any{
					"name":     "a name",
					"price":    20.40,
					"isActive": true,
				},
			},
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
				Write: &runtime.WriteResult{
					AffectedKey: true,
				},
			},
		},
		{
			Name: "fails when writing a value that does not match the data type",
			Persistence: &mocks.MockPersistence{
				SpySet: &mocks.Spy{},
				SpyGet: &mocks.Spy{
					Returns: []any{
						persistence.Value{},
						false,
					},
				},
			},
			ExpectSetCall: false,
			Key:           "key1",
			Value: persistence.Value{
				DataType: datatypes.KEY,
				Data: map[string]any{
					"name":     "a name",
					"price":    20.40,
					"isActive": true,
				},
			},
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
			},
			ExpectErr: true,
			Err:       appError.ErrInvalidValueType,
		},
		{
			Name: "fails when rewriting a value with different data type",
			Persistence: &mocks.MockPersistence{
				SpySet: &mocks.Spy{},
				SpyGet: &mocks.Spy{
					Returns: []any{
						persistence.Value{
							DataType: datatypes.KEY,
							Data:     "val",
						},
						true,
					},
				},
			},
			ExpectSetCall: false,
			Key:           "key1",
			Value: persistence.Value{
				DataType: datatypes.HASH,
				Data: map[string]any{
					"name":     "a name",
					"price":    20.40,
					"isActive": true,
				},
			},
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
			},
			ExpectErr: true,
			Err:       appError.ErrInvalidValueType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := handlers.SetHandler(tt.Persistence, tt.Key, tt.Value)

			if tt.ExpectErr && !errors.Is(tt.Err, err) {
				t.Errorf("Got %s err, Want %s", err, tt.Err)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %s result, Want %s", toString(tt.Result), toString(result))
			}

			if tt.Persistence.SpySet.Called != tt.ExpectSetCall {
				t.Errorf("Got %t when called Set, Want %t", tt.Persistence.SpySet.Called, true)
			}
		})
	}
}
