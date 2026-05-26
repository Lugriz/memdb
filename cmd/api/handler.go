package api

import (
	"net/http"

	"github.com/Lugriz/memdb/internal/engine"
	"github.com/Lugriz/memdb/internal/engine/runtime"
	"github.com/Lugriz/memdb/internal/parser"
	"github.com/gin-gonic/gin"
)

func Handler(eng engine.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dataType, err := parser.ParseDataType(ctx.Param("dataType"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{
				Type:    InvalidDataTypeError,
				Message: err.Error(),
			})
			return
		}

		operation, err := parser.ParseOperation(ctx.Param("operation"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{
				Type:    InvalidOperationError,
				Message: err.Error(),
			})
			return
		}

		var body RequestBody

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{
				Type:    InvalidSyntaxError,
				Message: "invalid body format",
			})
			return
		}

		command := &engine.Command{
			DataType:  dataType,
			Operation: operation,
			Key:       body.Key,
			Value:     body.Value,
		}

		result, err := eng.Execute(command)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{
				Type:    TypeFromError(err),
				Message: err.Error(),
			})
			return
		}

		if result.Type == runtime.READ_RESULT {
			if result.Read == nil {
				ctx.AbortWithStatus(http.StatusNoContent)
				return
			}

			ctx.AbortWithStatusJSON(http.StatusOK, toReadResponse(result.Read))
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, toWriteResponse(result.Write))
	}
}
