package main

import (
	"log"

	"localloop/services/frontend/internal/app"
	"localloop/services/frontend/internal/config"
)

func main() {
	cfg := config.Load()

	frontendApp, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal("Error initializing the app:", err)
	}

	if err := frontendApp.Server.Run(cfg.Port); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
