package main

import (
	"log"

	app "localloop/services/user/internal"
	"localloop/services/user/internal/config"
)

func main() {
	cfg := config.Load()

	userManagementApp, err := app.NewApp(
		cfg,
		app.WithInMemoryUserRepository(),
		app.WithUserService(),
		app.WithWebServer(),
	)

	if err != nil {
		log.Fatal("Error initializing the app:", err)
	}

	if err := userManagementApp.Server.Run(cfg.Port); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
