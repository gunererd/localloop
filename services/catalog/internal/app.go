package app

import (
	"localloop/libs/pkg/web"
	"localloop/services/catalog/internal/config"
	"localloop/services/catalog/internal/infrastructure/web/handler"
)

type App struct {
	config *config.Config
	Server *web.Server
}

type Option func(*App) error

func NewApp(cfg *config.Config, opts ...Option) (*App, error) {
	app := &App{
		config: cfg,
	}

	for _, opt := range opts {
		if err := opt(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func WithWebServer() Option {
	return func(app *App) error {
		server := web.NewServer()

		// Initialize handlers
		catalogHandler := handler.NewCatalogHandler()

		// Setup routes
		server.Router.HandleFunc("/categories", catalogHandler.ListCategories).Methods("GET")
		server.Router.HandleFunc("/categories", catalogHandler.CreateCategory).Methods("POST")
		server.Router.HandleFunc("/categories/{id}", catalogHandler.GetCategory).Methods("GET")
		server.Router.HandleFunc("/categories/{id}", catalogHandler.UpdateCategory).Methods("PUT")
		server.Router.HandleFunc("/categories/{id}", catalogHandler.DeleteCategory).Methods("DELETE")

		app.Server = server
		return nil
	}
}
