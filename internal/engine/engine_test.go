package engine_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/engine"
	"github.com/Lugriz/memdb/internal/mocks"
)

func mockRegistry() domain.DataTypeRegistry {
	return domain.DataTypeRegistry{
		domain.KEY: domain.OperationRegistry{
			domain.SET: func(_ domain.Persistence, _ string, _ any) (domain.OperationResult, error) {
				return domain.OperationResult{
					Write: &domain.WriteOperationResult{
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
		Command     *domain.Command
		Result      domain.OperationResult
		ExpectedErr bool
		Err         error
	}{
		{
			Name: "invalid key",
			Command: &domain.Command{
				DataType:  domain.KEY,
				Key:       "",
				Operation: domain.SET,
				Value:     "1",
			},
			ExpectedErr: true,
			Err:         domain.ErrInvalidKey,
		},
		{
			Name: "invalid data type",
			Command: &domain.Command{
				DataType:  -1,
				Key:       "test",
				Operation: domain.SET,
				Value:     "1",
			},
			ExpectedErr: true,
			Err:         domain.ErrInvalidDataType,
		},
		{
			Name: "invalid operation",
			Command: &domain.Command{
				DataType:  domain.KEY,
				Key:       "test",
				Operation: -1,
				Value:     "1",
			},
			ExpectedErr: true,
			Err:         domain.ErrInvalidOperation,
		},
		{
			Name: "execute command successfully",
			Command: &domain.Command{
				DataType:  domain.KEY,
				Key:       "test",
				Operation: domain.SET,
				Value:     "1",
			},
			Result: domain.OperationResult{
				Write: &domain.WriteOperationResult{
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
