package web

import (
	"context"
	"localloop/libs/pkg/web/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// Server is a general-purpose server struct
type Server struct {
	httpServer *http.Server
	Router     *mux.Router
}

// NewServer creates a new Server instance
func NewServer() *Server {
	router := mux.NewRouter()
	router.Use(middleware.Logging)

	return &Server{
		Router: router,
	}
}

// Run starts the server and handles graceful shutdown
func (s *Server) Run(port string) error {
	if port == "" {
		port = "8080"
	}

	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: s.Router,
	}

	// Start the server in a new goroutine
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Could not listen on %s: %v\n", s.httpServer.Addr, err)
		}
	}()

	log.Printf("Server is running on port %s\n", port)

	// Wait for interrupt signal for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down the server...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v\n", err)
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}
