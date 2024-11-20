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

	// Category routes
	router.HandleFunc("/categories", catalogHandler.ListCategories).Methods("GET")
	router.HandleFunc("/categories", catalogHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", catalogHandler.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", catalogHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", catalogHandler.DeleteCategory).Methods("DELETE")

	// Field routes
	router.HandleFunc("/fields", catalogHandler.ListFields).Methods("GET")
	router.HandleFunc("/fields", catalogHandler.CreateField).Methods("POST")
	router.HandleFunc("/fields/{id}", catalogHandler.GetField).Methods("GET")
	router.HandleFunc("/fields/{id}", catalogHandler.UpdateField).Methods("PUT")
	router.HandleFunc("/fields/{id}", catalogHandler.DeleteField).Methods("DELETE")

	// Category-Field assignment
	router.HandleFunc("/categories/{categoryId}/fields", catalogHandler.GetCategoryFields).Methods("GET")
	router.HandleFunc("/categories/{categoryId}/fields/{fieldId}", catalogHandler.AssignFieldToCategory).Methods("POST")

	// Field Type routes
	router.HandleFunc("/field-types", catalogHandler.ListFieldTypes).Methods("GET")
	router.HandleFunc("/field-types", catalogHandler.CreateFieldType).Methods("POST")
	router.HandleFunc("/field-types/{id}", catalogHandler.GetFieldType).Methods("GET")
	router.HandleFunc("/field-types/{id}", catalogHandler.UpdateFieldType).Methods("PUT")
	router.HandleFunc("/field-types/{id}", catalogHandler.DeleteFieldType).Methods("DELETE")

	// Field Type Discriminator routes
	router.HandleFunc("/field-type-discriminators", catalogHandler.ListFieldTypeDiscriminators).Methods("GET")
	router.HandleFunc("/field-type-discriminators", catalogHandler.CreateFieldTypeDiscriminator).Methods("POST")
	router.HandleFunc("/field-type-discriminators/{id}", catalogHandler.GetFieldTypeDiscriminator).Methods("GET")
}
