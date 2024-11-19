package app

import (
	"localloop/libs/pkg/web"
	"localloop/services/catalog/internal/config"
	"localloop/services/catalog/internal/domain/catalog"
)

type App struct {
	config         *config.Config
	Server         *web.Server
	CatalogService *catalog.Service
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

func WithCatalogService() Option {
	return func(app *App) error {
		// We'll implement the repository later with sqlc
		// For now, we'll just initialize the service
		app.CatalogService = catalog.NewService(nil, catalog.ServiceConfig{})
		return nil
	}
}

func WithWebServer() Option {
	return func(app *App) error {
		app.Server = web.NewServer()
		// We'll add routes and handlers later
		return nil
	}
}
