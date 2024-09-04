package main

import (
	"context"

	repo "github.com/0xMik3/link-metrics/internal/application/repository/psql"
	s "github.com/0xMik3/link-metrics/internal/application/services/shortener"
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
	psql.Sync_tables(db)
	defer db.Close()

	shortenerRepo := repo.NewShortenerRepository(db)
	shortenerService := s.NewShortenerService(ctx, shortenerRepo)

	restServer := rest.NewRestHandler(shortenerService)
	restServer.InitializeRoutes()

	mux := cmux.NewCmuxConfig(config.Port)
	go restServer.Start(mux.HttpListener())
	mux.Start()
}
