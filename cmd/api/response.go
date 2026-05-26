package api

import (
	"github.com/Lugriz/memdb/internal/engine/runtime"
)

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ReadResponse struct {
	Value any `json:"value"`
}

type WriteResponse struct {
	AffectedKey bool `json:"affected_key"`
}

func toReadResponse(read *runtime.ReadResult) ReadResponse {
	return ReadResponse{
		Value: read.Value,
	}
}

func toWriteResponse(write *runtime.WriteResult) WriteResponse {
	return WriteResponse{
		AffectedKey: write.AffectedKey,
	}
}
