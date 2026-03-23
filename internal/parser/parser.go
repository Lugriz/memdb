package parser

import (
	"strings"

	"github.com/Lugriz/memdb/internal/domain"
)

const minLengthTokens = 3

func ParseStatement(statement string) (*domain.Command, error) {
	tokens := strings.Fields(statement)

	if len(tokens) < minLengthTokens {
		return nil, domain.ErrMissingArgs
	}

	dataType, err := domain.ParseDataType(tokens[0])
	if err != nil {
		return nil, domain.ErrInvalidDataType
	}

	operationType, err := domain.ParseOperationType(tokens[2])
	if err != nil {
		return nil, domain.ErrInvalidOperationType
	}

	key := string(tokens[1])
	args := tokens[3:]

	return &domain.Command{
		DataType:      dataType,
		Key:           key,
		OperationType: operationType,
		Args:          args,
	}, nil

}
