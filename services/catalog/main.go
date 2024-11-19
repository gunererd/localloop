package main

import (
	"log"

	app "localloop/services/catalog/internal"
	"localloop/services/catalog/internal/config"
)

func main() {
	cfg := config.Load()

	catalogApp, err := app.NewApp(
		cfg,
		app.WithPostgresDatabase(),
		app.WithPostgresCatalogRepository(),
		app.WithCatalogService(),
		app.WithWebServer(),
	)

	if err != nil {
		log.Fatal("Error initializing the app:", err)
	}

	if err := catalogApp.Server.Run(cfg.Port); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
