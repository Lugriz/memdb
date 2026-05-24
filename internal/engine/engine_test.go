package engine_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/datatypes"
	"github.com/Lugriz/memdb/internal/engine"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	appErrors "github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/mocks"
	"github.com/Lugriz/memdb/internal/persistence"
	"github.com/Lugriz/memdb/internal/registry"
)

func mockRegistry() registry.DataTypeRegistry {
	return registry.DataTypeRegistry{
		datatypes.KEY: registry.OperationRegistry{
			datatypes.SET: func(_ persistence.Persistence, _ string, _ any) (runtime.Result, error) {
				return runtime.Result{
					Write: &runtime.WriteResult{
						AffectedKey: true,
					},
				}, nil
			},
		},
	}
}

func TestEngine(t *testing.T) {
	tests := []struct {
		Name        string
		Command     *engine.Command
		Result      runtime.Result
		ExpectedErr bool
		Err         error
	}{
		{
			Name: "invalid key",
			Command: &engine.Command{
				DataType:  datatypes.KEY,
				Key:       "",
				Operation: datatypes.SET,
				Value:     "1",
			},
			ExpectedErr: true,
			Err:         appErrors.ErrInvalidKey,
		},
		{
			Name: "invalid data type",
			Command: &engine.Command{
				DataType:  -1,
				Key:       "test",
				Operation: datatypes.SET,
				Value:     "1",
			},
			ExpectedErr: true,
			Err:         appErrors.ErrInvalidDataType,
		},
		{
			Name: "invalid operation",
			Command: &engine.Command{
				DataType:  datatypes.KEY,
				Key:       "test",
				Operation: -1,
				Value:     "1",
			},
			ExpectedErr: true,
			Err:         appErrors.ErrInvalidOperation,
		},
		{
			Name: "execute command successfully",
			Command: &engine.Command{
				DataType:  datatypes.KEY,
				Key:       "test",
				Operation: datatypes.SET,
				Value:     "1",
			},
			Result: runtime.Result{
				Write: &runtime.WriteResult{
					AffectedKey: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			engine := engine.NewEngine(&mocks.MockPersistence{}, mockRegistry())

			result, err := engine.Execute(tt.Command)

			if tt.ExpectedErr && !errors.Is(tt.Err, err) {
				t.Errorf("Got %s err, Want %s", err, tt.Err)
			}

			if !reflect.DeepEqual(tt.Result, result) {
				t.Errorf("Got %v result, Want %v", result, tt.Result)
			}
		})
	}
}
