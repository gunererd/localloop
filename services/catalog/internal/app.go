package app

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"localloop/libs/pkg/db"
	"localloop/services/catalog/internal/config"
	"localloop/services/catalog/internal/domain/catalog"
	"localloop/services/catalog/internal/infrastructure/repository/inmemory"
	"localloop/services/catalog/internal/infrastructure/repository/postgresql"
	"localloop/services/catalog/internal/infrastructure/web"
)

type App struct {
	config         *config.Config
	PostgresDB     *sql.DB
	CatalogRepo    catalog.Repository
	CatalogService *catalog.Service
	Server         *web.CatalogManagementServer
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

// WithPostgresDatabase sets up the PostgreSQL connection and applies it to the App
func WithPostgresDatabase() Option {
	return func(app *App) error {
		log.Println("Connecting to PostgreSQL:", app.config.PostgresURI)
		db, err := db.ConnectPostgreSQL(app.config.PostgresURI)
		if err != nil {
			log.Fatal("Failed to connect to PostgreSQL:", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = db.PingContext(ctx)
		if err != nil {
			log.Fatal("Failed to verify PostgreSQL connection:", err)
		}

		app.PostgresDB = db

		log.Println("Connected to PostgreSQL successfully")
		return nil
	}
}

func WithPostgresCatalogRepository() Option {
	return func(app *App) error {
		if app.PostgresDB == nil {
			return errors.New("PostgresDB is not initialized. Make sure to call WithPostgresDatabase first")
		}

		app.CatalogRepo = postgresql.NewCatalogRepository(app.PostgresDB)
		return nil
	}
}

func WithCatalogService() Option {
	return func(app *App) error {
		if app.CatalogRepo == nil {
			return errors.New("CatalogRepo is not initialized. Make sure to call WithPostgresCatalogRepository first")
		}

		app.CatalogService = catalog.NewService(app.CatalogRepo, catalog.ServiceConfig{})
		return nil
	}
}

func WithWebServer() Option {
	return func(app *App) error {
		if app.CatalogService == nil {
			return errors.New("CatalogService is not initialized. Make sure to call WithCatalogService first")
		}

		app.Server = web.NewCatalogManagementServer(app.CatalogService)
		return nil
	}
}

func WithInMemoryCatalogRepository() Option {
	return func(app *App) error {
		app.CatalogRepo = inmemory.NewCatalogRepository()
		return nil
	}
}
