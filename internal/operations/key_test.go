package operations_test

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/mocks"
	"github.com/Lugriz/memdb/internal/operations"
)

func toString(s any) string {
	r, _ := json.Marshal(s)
	return string(r)
}

func TestKeySetHandler(t *testing.T) {
	tests := []struct {
		Name            string
		Persistence     *mocks.MockPersistence
		ExpectSetKVCall bool
		Key             string
		Value           any
		Result          domain.OperationResult
		ExpectErr       bool
		Err             error
	}{
		{
			Name: "successfull",
			Persistence: &mocks.MockPersistence{
				SpySetKV: &mocks.Spy{},
			},
			ExpectSetKVCall: true,
			Key:             "key1",
			Value:           "val 1",
			Result: domain.OperationResult{
				Type: domain.WRITE_OPERATION,
				Write: &domain.WriteOperationResult{
					AffectedKey: true,
				},
			},
		},
		{
			Name: "error when invalid value type",
			Persistence: &mocks.MockPersistence{
				SpySetKV: &mocks.Spy{},
			},
			ExpectSetKVCall: false,
			Key:             "key1",
			Value:           struct{}{},
			Result: domain.OperationResult{
				Type: domain.WRITE_OPERATION,
			},
			ExpectErr: true,
			Err:       domain.ErrInvalidValueType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := operations.KeySetHandler(tt.Persistence, tt.Key, tt.Value)

			if tt.ExpectErr && !errors.Is(tt.Err, err) {
				t.Errorf("Got %s err, Want %s", err, tt.Err)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %s result, Want %s", toString(tt.Result), toString(result))
			}

			if tt.Persistence.SpySetKV.Called != tt.ExpectSetKVCall {
				t.Errorf("Got %t when called SetKV, Want %t", tt.Persistence.SpySetKV.Called, true)
			}
		})
	}

}

func TestKeyGetHandler(t *testing.T) {
	tests := []struct {
		Name        string
		Persistence *mocks.MockPersistence
		Key         string
		Value       any
		Result      domain.OperationResult
	}{
		{
			Name: "returns a value",
			Persistence: &mocks.MockPersistence{
				SpyGetKV: &mocks.Spy{
					Returns: []any{
						domain.Value{
							DataType: domain.KEY,
							Data:     "value1",
						},
						true,
					},
				},
			},
			Key: "key1",
			Result: domain.OperationResult{
				Type: domain.READ_OPERATION,
				Read: &domain.ReadOperationResult{
					Value: "value1",
				},
			},
		},
		{
			Name: "not found value",
			Persistence: &mocks.MockPersistence{
				SpyGetKV: &mocks.Spy{
					Returns: []any{domain.Value{}, false},
				},
			},
			Key: "key1",
			Result: domain.OperationResult{
				Type: domain.READ_OPERATION,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := operations.KeyGetHandler(tt.Persistence, tt.Key, tt.Value)

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
		Result      domain.OperationResult
	}{
		{
			Name: "Delete the key",
			Persistence: &mocks.MockPersistence{
				SpyDeleteKV: &mocks.Spy{
					Returns: []any{true},
				},
			},
			Key: "key1",
			Result: domain.OperationResult{
				Type: domain.WRITE_OPERATION,
				Write: &domain.WriteOperationResult{
					AffectedKey: true,
				},
			},
		},
		{
			Name: "Unexisting key",
			Persistence: &mocks.MockPersistence{
				SpyDeleteKV: &mocks.Spy{
					Returns: []any{false},
				},
			},
			Key: "key1",
			Result: domain.OperationResult{
				Type: domain.WRITE_OPERATION,
				Write: &domain.WriteOperationResult{
					AffectedKey: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := operations.KeyDelHandler(tt.Persistence, tt.Key, tt.Value)

			if err != nil {
				t.Errorf("Got %s err, Want %v", err, nil)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %s result, Want %s", toString(result), toString(tt.Result))
			}
		})
	}
}
