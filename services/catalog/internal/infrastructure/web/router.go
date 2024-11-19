package web

import (
	"localloop/libs/pkg/web"
	catalog "localloop/services/catalog/internal/domain"
	"localloop/services/catalog/internal/infrastructure/web/handler"
)

type CatalogManagementServer struct {
	*web.Server    // Embedding the shared Server struct
	catalogService *catalog.Service
}

func NewCatalogManagementServer(catalogService *catalog.Service) *CatalogManagementServer {
	server := &CatalogManagementServer{
		Server:         web.NewServer(),
		catalogService: catalogService,
	}

	server.setupRoutes()
	return server
}

func (s *CatalogManagementServer) setupRoutes() {
	router := s.Router

	catalogHandler := handler.NewCatalogHandler(s.catalogService)

	router.HandleFunc("/categories", catalogHandler.ListCategories).Methods("GET")
	router.HandleFunc("/categories", catalogHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", catalogHandler.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", catalogHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", catalogHandler.DeleteCategory).Methods("DELETE")
}
