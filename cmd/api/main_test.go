package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/cmd/api"
	"github.com/Lugriz/memdb/internal/engine"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/errors"
	"github.com/Lugriz/memdb/internal/mocks"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func toMap(body []byte) map[string]any {
	if body == nil {
		return nil
	}
	var st map[string]any

	json.Unmarshal(body, &st)

	return st
}

func mapToString(strct any) string {
	b, _ := json.Marshal(strct)

	return string(b)
}

func TestRouter(t *testing.T) {
	tests := []struct {
		Name                 string
		Engine               engine.Engine
		Path                 string
		Body                 []byte
		ExpectedStatusCode   int
		ExpectedBodyResponse map[string]any
	}{
		{
			Name:               "key: invalid data type",
			Path:               "/invalid/get",
			Engine:             &mocks.MockEngine{},
			ExpectedStatusCode: 400,
			ExpectedBodyResponse: map[string]any{
				"type":    api.InvalidDataTypeError,
				"message": errors.ErrInvalidDataType.Error(),
			},
		},
		{
			Name:               "key: invalid operation",
			Path:               "/key/invalid",
			Engine:             &mocks.MockEngine{},
			ExpectedStatusCode: 400,
			ExpectedBodyResponse: map[string]any{
				"type":    api.InvalidOperationError,
				"message": errors.ErrInvalidOperation.Error(),
			},
		},
		{
			Name:               "key: invalid request body format",
			Path:               "/key/get",
			Engine:             &mocks.MockEngine{},
			Body:               []byte("invalid json"),
			ExpectedStatusCode: 400,
			ExpectedBodyResponse: map[string]any{
				"type":    api.InvalidSyntaxError,
				"message": "invalid body format",
			},
		},
		{
			Name: "key: invalid key",
			Path: "/key/get",
			Body: []byte(`{"key":""}`),
			Engine: &mocks.MockEngine{
				SpyExecute: &mocks.Spy{
					Returns: []any{runtime.Result{}, errors.ErrInvalidKey},
				},
			},
			ExpectedStatusCode: 400,
			ExpectedBodyResponse: map[string]any{
				"type":    api.InvalidKeyError,
				"message": errors.ErrInvalidKey.Error(),
			},
		},
		{
			Name: "key: invalid value type",
			Path: "/key/set",
			Body: []byte(`{"key": "key"}`),
			Engine: &mocks.MockEngine{
				SpyExecute: &mocks.Spy{
					Returns: []any{runtime.Result{}, errors.ErrInvalidValueType},
				},
			},
			ExpectedStatusCode: 400,
			ExpectedBodyResponse: map[string]any{
				"type":    api.InvalidValueTypeError,
				"message": errors.ErrInvalidValueType.Error(),
			},
		},
		{
			Name: "key: get unexisting key",
			Path: "/key/get",
			Body: []byte(`{"key": "key1"}`),
			Engine: &mocks.MockEngine{
				SpyExecute: &mocks.Spy{
					Returns: []any{runtime.Result{
						Type: runtime.READ_RESULT,
					}, nil},
				},
			},
			ExpectedStatusCode:   204,
			ExpectedBodyResponse: nil,
		},
		{
			Name: "key: get existing key",
			Path: "/key/get",
			Body: []byte(`{"key": "key1"}`),
			Engine: &mocks.MockEngine{
				SpyExecute: &mocks.Spy{
					Returns: []any{runtime.Result{
						Type: runtime.READ_RESULT,
						Read: &runtime.ReadResult{
							Value: "value 1",
						},
					}, nil},
				},
			},
			ExpectedStatusCode: 200,
			ExpectedBodyResponse: map[string]any{
				"value": "value 1",
			},
		},
		{
			Name: "key: set a key",
			Path: "/key/set",
			Body: []byte(`{"key": "key1", "value": "value 1"}`),
			Engine: &mocks.MockEngine{
				SpyExecute: &mocks.Spy{
					Returns: []any{runtime.Result{
						Type: runtime.WRITE_RESULT,
						Write: &runtime.WriteResult{
							AffectedKey: true,
						},
					}, nil},
				},
			},
			ExpectedStatusCode: 200,
			ExpectedBodyResponse: map[string]any{
				"affected_key": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			router := api.Router(tt.Engine)
			writer := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", tt.Path, bytes.NewReader(tt.Body))
			router.ServeHTTP(writer, request)

			if tt.ExpectedStatusCode != writer.Code {
				t.Errorf("want %d status code, got %d", tt.ExpectedStatusCode, writer.Code)
			}

			if !reflect.DeepEqual(tt.ExpectedBodyResponse, toMap(writer.Body.Bytes())) {
				t.Errorf("want %v response, got %v", mapToString(tt.ExpectedBodyResponse), writer.Body.String())
			}
		})
	}
}
