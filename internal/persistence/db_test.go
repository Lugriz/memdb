package persistence_test

import (
	"reflect"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/persistence"
)

func TestGet(t *testing.T) {
	tests := []struct {
		Name          string
		Key           string
		SetupFunc     func(*persistence.InMemoryDB)
		ExpectedValue domain.Value
		ExpectedBool  bool
	}{
		{
			Name: "return existing value",
			Key:  "key",
			SetupFunc: func(db *persistence.InMemoryDB) {
				db.Set("key", domain.Value{
					DataType: domain.KEY,
					Data:     "value 1",
				})
			},
			ExpectedValue: domain.Value{
				DataType: domain.KEY,
				Data:     "value 1",
			},
			ExpectedBool: true,
		},
		{
			Name:          "return unexisting value",
			Key:           "key",
			SetupFunc:     nil,
			ExpectedValue: domain.Value{},
			ExpectedBool:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			db := persistence.NewInMemoryDB()

			if tt.SetupFunc != nil {
				tt.SetupFunc(db)
			}

			value, ok := db.Get(tt.Key)

			if !reflect.DeepEqual(tt.ExpectedValue, value) {
				t.Errorf("Got %v, Want %v", value, tt.ExpectedValue)
			}

			if tt.ExpectedBool != ok {
				t.Errorf("Got %t err, Want %t", ok, tt.ExpectedBool)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		Name         string
		Key          string
		SetupFunc    func(*persistence.InMemoryDB)
		ExpectedBool bool
	}{
		{
			Name:         "returns false when unexisting key",
			Key:          "key",
			SetupFunc:    nil,
			ExpectedBool: false,
		},
		{
			Name: "returns true when deleting a key",
			Key:  "key",
			SetupFunc: func(db *persistence.InMemoryDB) {
				db.Set("key", domain.Value{
					DataType: domain.KEY,
					Data:     "val",
				})
			},
			ExpectedBool: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			db := persistence.NewInMemoryDB()

			if tt.SetupFunc != nil {
				tt.SetupFunc(db)
			}

			ok := db.Delete(tt.Key)

			if tt.ExpectedBool != ok {
				t.Errorf("Got %t err, Want %t", ok, tt.ExpectedBool)
			}
		})
	}
}
