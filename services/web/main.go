package main

import (
	"log"

	"localloop/services/web/internal/app"
	"localloop/services/web/internal/config"
)

func main() {
	cfg := config.Load()

	webApp, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal("Error initializing the app:", err)
	}

	if err := webApp.Server.Run(cfg.Port); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
