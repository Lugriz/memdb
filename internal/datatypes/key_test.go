package datatypes

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/mocks"
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
		Args            [][]byte
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
			Args:            [][]byte{[]byte("val 1")},
			Result: domain.OperationResult{
				Write: &domain.WriteOperationResult{
					AffectedKeys: 1,
				},
			},
		},
		{
			Name: "missing args",
			Persistence: &mocks.MockPersistence{
				SpySetKV: &mocks.Spy{},
			},
			ExpectSetKVCall: false,
			Key:             "key1",
			Args:            [][]byte{},
			Result:          domain.OperationResult{},
			ExpectErr:       true,
			Err:             domain.ErrMissingArgs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := KeySetHandler(tt.Persistence, tt.Key, tt.Args)

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
		Args        [][]byte
		Result      domain.OperationResult
	}{
		{
			Name: "returns a value",
			Persistence: &mocks.MockPersistence{
				SpyGetKV: &mocks.Spy{
					Returns: []any{"value1", true},
				},
			},
			Key: "key1",
			Result: domain.OperationResult{
				Read: &domain.ReadOperationResult{
					Value: "value1",
				},
			},
		},
		{
			Name: "not found value",
			Persistence: &mocks.MockPersistence{
				SpyGetKV: &mocks.Spy{
					Returns: []any{"", false},
				},
			},
			Key:    "key1",
			Result: domain.OperationResult{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := KeyGetHandler(tt.Persistence, tt.Key, tt.Args)

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
		Args        [][]byte
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
				Write: &domain.WriteOperationResult{
					AffectedKeys: 1,
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
				Write: &domain.WriteOperationResult{
					AffectedKeys: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := KeyDelHandler(tt.Persistence, tt.Key, tt.Args)

			if err != nil {
				t.Errorf("Got %s err, Want %v", err, nil)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %s result, Want %s", toString(result), toString(tt.Result))
			}
		})
	}
}
