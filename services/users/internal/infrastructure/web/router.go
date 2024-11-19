package web

import (
	"localloop/libs/pkg/web"
	"localloop/services/users/internal/domain/user"
	"localloop/services/users/internal/infrastructure/web/handler"
)

type UserManagementServer struct {
	*web.Server // Embedding the shared Server struct
	userService *user.Service
}

// NewUserManagementServer creates a new instance of UserManagementServer
func NewUserManagementServer(userService *user.Service) *UserManagementServer {
	server := &UserManagementServer{
		Server:      web.NewServer(),
		userService: userService,
	}

	server.setupRoutes() // Set up the routes
	return server
}

// setupRoutes initializes all the routes for the UserManagementServer
func (s *UserManagementServer) setupRoutes() {
	router := s.Router

	// Initialize handler and authenticator
	userHandler := handler.NewUserHandler(s.userService)
	authenticator := NewAuthenticator(s.userService)

	// Public routes
	router.HandleFunc("/users/register", userHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", userHandler.LoginUser).Methods("POST")

	// Authenticated routes
	authRouter := router.PathPrefix("/users").Subrouter()
	authRouter.Use(authenticator.AuthMiddleware)
	authRouter.HandleFunc("/me", userHandler.GetCurrentUser).Methods("GET")
	authRouter.HandleFunc("/me", userHandler.UpdateCurrentUser).Methods("PUT")
}
