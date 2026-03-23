package parser_test

import (
	"errors"
	"testing"

	"github.com/Lugriz/memdb/internal/domain"
	"github.com/Lugriz/memdb/internal/parser"
)

func TestParser(t *testing.T) {
	tests := []struct {
		Name      string
		Input     string
		Result    *domain.Command
		ExpectErr bool
		Err       error
	}{
		{
			Name:  "returns parsed KEY/SET statement",
			Input: "KEY keyname SET value1",
			Result: &domain.Command{
				DataType:      domain.KEY,
				Key:           "keyname",
				OperationType: domain.SET,
				Args:          []string{"value1"},
			},
		},
		{
			Name:  "returns parsed KEY/GET statement",
			Input: "KEY keyname GET",
			Result: &domain.Command{
				DataType:      domain.KEY,
				Key:           "keyname",
				OperationType: domain.GET,
				Args:          []string{},
			},
		},
		{
			Name:  "accepts multiple spaces",
			Input: "KEY     keyname   SET     value1",
			Result: &domain.Command{
				DataType:      domain.KEY,
				Key:           "keyname",
				OperationType: domain.SET,
				Args:          []string{"value1"},
			},
		},
		{
			Name:  "leading and trailing whitespaces",
			Input: "    KEY keyname SET value1   ",
			Result: &domain.Command{
				DataType:      domain.KEY,
				Key:           "keyname",
				OperationType: domain.SET,
				Args:          []string{"value1"},
			},
		},
		{
			Name:  "using keywords for key/values",
			Input: "KEY SET SET SET",
			Result: &domain.Command{
				DataType:      domain.KEY,
				Key:           "SET",
				OperationType: domain.SET,
				Args:          []string{"SET"},
			},
		},
		{
			Name:      "invalid data type",
			Input:     "INVALIDTYPE keyname GET",
			Result:    nil,
			ExpectErr: true,
			Err:       domain.ErrInvalidDataType,
		},
		{
			Name:      "invalid operation type",
			Input:     "KEY keyname INVALIDGET",
			Result:    nil,
			ExpectErr: true,
			Err:       domain.ErrInvalidOperationType,
		},
		{
			Name:      "missing arguments in statement",
			Input:     "KEY keyname",
			Result:    nil,
			ExpectErr: true,
			Err:       domain.ErrMissingArgs,
		},
		{
			Name:      "empty key",
			Input:     "KEY  SET value1",
			Result:    nil,
			ExpectErr: true,
			Err:       domain.ErrInvalidOperationType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			cmd, err := parser.ParseStatement(tt.Input)

			if tt.ExpectErr && !errors.Is(err, tt.Err) {
				t.Errorf("want %s error, got %s", tt.Err, err)
			}

			if !tt.ExpectErr && !compareCommand(cmd, tt.Result) {
				t.Errorf("want %v, got %v", tt.Result, cmd)
			}
		})
	}
}

func compareCommand(c1, c2 *domain.Command) bool {
	return c1.DataType == c2.DataType &&
		c1.OperationType == c2.OperationType &&
		c1.Key == c2.Key &&
		compareArgs(c1.Args, c2.Args)
}

func compareArgs(args, target []string) bool {
	if len(args) != len(target) {
		return false
	}

	for i := range args {
		if args[i] != target[i] {
			return false
		}
	}

	return true
}
