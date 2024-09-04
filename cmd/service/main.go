package main

import (
	"context"

	"github.com/0xMik3/link-metrics/internal/config"
	"github.com/0xMik3/link-metrics/internal/infra/adapters/driven/cmux"
	"github.com/0xMik3/link-metrics/internal/infra/adapters/driven/psql"
	"github.com/0xMik3/link-metrics/internal/infra/adapters/driver/rest"
)

func main() {
	var config config.Config

	ctx := context.Background()
	ctx = config.GetEnvs(ctx)

	db, err := psql.Connect(&config)
	if err != nil {
		return
	}
	defer db.Close()

	restServer := rest.NewRestHandler()
	restServer.InitializeRoutes()

	mux := cmux.NewCmuxConfig(config.Port)
	go restServer.Start(mux.HttpListener())
	mux.Start()
}
