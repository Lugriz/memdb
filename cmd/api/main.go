package api

import (
	"github.com/Lugriz/memdb/internal/engine"
	"github.com/gin-gonic/gin"
)

func Router(eng engine.Engine) *gin.Engine {
	router := gin.Default()

	router.POST("/:dataType/:operation", Handler(eng))

	return router
}

func Run(eng engine.Engine) error {
	return Router(eng).Run()
}
