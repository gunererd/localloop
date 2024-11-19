package app

import (
	"localloop/libs/pkg/web"
	"localloop/services/web/internal/config"
	"localloop/services/web/internal/handler"
	"localloop/services/web/internal/repository"
)

type App struct {
	config *config.Config
	Server *web.Server
}

func NewApp(cfg *config.Config) (*App, error) {
	server := web.NewServer()

	// Initialize repositories
	userRepo := repository.NewUserRepository(cfg.UserServiceURL)

	// Initialize handlers
	authHandler, err := handler.NewAuthHandler(userRepo)
	if err != nil {
		return nil, err
	}

	// Setup routes
	server.Router.HandleFunc("/login", authHandler.ShowLoginPage).Methods("GET")
	server.Router.HandleFunc("/auth/login", authHandler.HandleLogin).Methods("POST")
	server.Router.HandleFunc("/register", authHandler.ShowRegisterPage).Methods("GET")
	server.Router.HandleFunc("/auth/register", authHandler.HandleRegister).Methods("POST")

	return &App{
		config: cfg,
		Server: server,
	}, nil
}
