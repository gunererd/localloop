package app

import (
	"context"
	"errors"
	"log"
	"time"

	"localloop/libs/pkg/db"
	"localloop/services/users/internal/config"
	"localloop/services/users/internal/domain/user"
	"localloop/services/users/internal/infrastructure/repository"
	"localloop/services/users/internal/infrastructure/web"

	"go.mongodb.org/mongo-driver/mongo"
)

// App holds all components and configuration for the users service
type App struct {
	config      *config.Config
	MongoClient *mongo.Client
	UserRepo    user.Repository
	UserService *user.Service
	Server      *web.UserManagementServer
}

// Option is a function that applies a configuration to the App
type Option func(*App) error

// NewApp creates a new App instance with default values and applies options
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

// WithMongoDatabase sets up the MongoDB connection and applies it to the App
func WithMongoDatabase() Option {
	return func(app *App) error {
		log.Println(app.config.MongoURI)
		client, err := db.ConnectMongoDB(app.config.MongoURI)
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt to list the databases to verify authentication
		_, err = client.ListDatabaseNames(ctx, map[string]interface{}{})
		if err != nil {
			log.Fatal("Failed to authenticate with MongoDB:", err)
		}

		app.MongoClient = client

		log.Println("Authenticated and connected to MongoDB successfully")
		return nil
	}
}

func WithMongoUserRepositroy() Option {
	return func(app *App) error {

		if app.MongoClient == nil {
			return errors.New("MongoClient is not initialized. Make sure to call WithMongoDatabase before WithUserService")
		}

		database := app.MongoClient.Database("users")
		app.UserRepo = repository.NewMongoRepository(database)
		return nil
	}
}

func WithInMemoryUserRepository() Option {
	return func(app *App) error {
		app.UserRepo = repository.NewInMemoryRepository()
		return nil
	}
}

// WithUserService initializes the user repository and service and applies them to the App
func WithUserService() Option {
	return func(app *App) error {
		app.UserService = user.NewService(app.UserRepo, user.ServiceConfig{
			JWTSecret:            app.config.JWTSecret,
			JWTExpirationMinutes: app.config.JWTExpirationMinutes,
			SaltLength:           app.config.SaltLength,
		})
		return nil
	}
}

func WithWebServer() Option {
	return func(app *App) error {
		app.Server = web.NewUserManagementServer(app.UserService)
		return nil
	}
}
