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

func TestHashAddHandler(t *testing.T) {
	tests := []struct {
		Name                    string
		Persistence             *mocks.MockPersistence
		ExpectGetCall           bool
		ExpectSetCall           bool
		ExpectSetCallWithParams []any
		Key                     string
		Value                   any
		Result                  runtime.Result
		ExpectErr               bool
		Err                     error
	}{
		{
			Name: "write key successfully",
			Persistence: &mocks.MockPersistence{
				SpySet: &mocks.Spy{},
				SpyGet: &mocks.Spy{
					Returns: []any{persistence.Value{
						DataType: datatypes.HASH,
						Data: map[string]any{
							"name": "a name",
						},
					}, true},
				},
			},
			ExpectGetCall: true,
			ExpectSetCall: true,
			ExpectSetCallWithParams: []any{
				"key1",
				persistence.Value{
					DataType: datatypes.HASH,
					Data: map[string]any{
						"name":     "a name",
						"price":    20.40,
						"isActive": true,
					},
				},
			},
			Key: "key1",
			Value: map[string]any{
				"price":    20.40,
				"isActive": true,
			},
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
				Write: &runtime.WriteResult{
					AffectedKey: true,
				},
			},
		},
		{
			Name: "fails when adding a not map data type value",
			Persistence: &mocks.MockPersistence{
				SpySet: &mocks.Spy{},
				SpyGet: &mocks.Spy{},
			},
			ExpectGetCall: false,
			ExpectSetCall: false,
			Key:           "key1",
			Value:         "invalid_type",
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
			},
			ExpectErr: true,
			Err:       appError.ErrInvalidValueType,
		},
		{
			Name: "fails when adding a value to an unexisting key",
			Persistence: &mocks.MockPersistence{
				SpySet: &mocks.Spy{},
				SpyGet: &mocks.Spy{
					Returns: []any{
						persistence.Value{},
						false,
					},
				},
			},
			ExpectGetCall: true,
			ExpectSetCall: false,
			Key:           "key1",
			Value: map[string]any{
				"name":     "a name",
				"price":    20.40,
				"isActive": true,
			},
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
			},
			ExpectErr: true,
			Err:       appError.ErrUnknownKey,
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
			ExpectGetCall: true,
			ExpectSetCall: false,
			Key:           "key1",
			Value: map[string]any{
				"name":     "a name",
				"price":    20.40,
				"isActive": true,
			},
			Result: runtime.Result{
				Type: runtime.WRITE_RESULT,
			},
			ExpectErr: true,
			Err:       appError.ErrInvalidValueTypeOperation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := handlers.HashAddHandler(tt.Persistence, tt.Key, tt.Value)

			if tt.ExpectErr && !errors.Is(tt.Err, err) {
				t.Errorf("Got %s err, Want %s", err, tt.Err)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %s result, Want %s", toString(tt.Result), toString(result))
			}

			if tt.ExpectGetCall != tt.Persistence.SpyGet.Called {
				t.Errorf("Got %t when called Get, Want %t", tt.Persistence.SpyGet.Called, tt.ExpectGetCall)
			}

			if tt.ExpectSetCall != tt.Persistence.SpySet.Called {
				t.Errorf("Got %t when called Set, Want %t", tt.Persistence.SpySet.Called, tt.ExpectSetCall)
			}

			if tt.ExpectSetCall && !reflect.DeepEqual(tt.ExpectSetCallWithParams, tt.Persistence.SpySet.Params) {
				t.Errorf("Got %s params when called Set, Want %s", toString(tt.Persistence.SpySet.Params), toString(tt.ExpectSetCallWithParams))
			}
		})
	}
}
