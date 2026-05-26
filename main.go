package main

import (
	"fmt"

	"github.com/Lugriz/memdb/cmd/api"
	"github.com/Lugriz/memdb/internal/engine"
	"github.com/Lugriz/memdb/internal/persistence"
	"github.com/Lugriz/memdb/internal/registry"
)

func main() {
	eng := engine.NewEngine(persistence.NewInMemoryDB(), registry.DataTypeRegistryConfig)

	if err := api.Run(eng); err != nil {
		panic(fmt.Sprintf("fails when running the server: %s", err))
	}
}
