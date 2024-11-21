package web

import (
	"localloop/libs/pkg/web"
	catalog "localloop/services/catalog/internal/domain"
	h "localloop/services/catalog/internal/infrastructure/web/handler"
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
	ch := h.NewCatalogHandler(s.catalogService)

	// Category routes
	router.HandleFunc("/categories", ch.ListCategories).Methods("GET")
	router.HandleFunc("/categories", h.HandleRequest[h.CreateCategoryRequest](ch.CreateCategory)).Methods("POST")
	router.HandleFunc("/categories/{id}", h.HandleRequest[struct{}](ch.GetCategory)).Methods("GET")
	router.HandleFunc("/categories/{id}", ch.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", ch.DeleteCategory).Methods("DELETE")

	// Field routes
	router.HandleFunc("/fields", ch.ListFields).Methods("GET")
	router.HandleFunc("/fields", h.HandleRequest[h.CreateFieldRequest](ch.CreateField)).Methods("POST")
	router.HandleFunc("/fields/{id}", h.HandleRequest[struct{}](ch.GetField)).Methods("GET")
	router.HandleFunc("/fields/{id}", h.HandleRequest[h.UpdateFieldRequest](ch.UpdateField)).Methods("PUT")
	router.HandleFunc("/fields/{id}", h.HandleRequest[struct{}](ch.DeleteField)).Methods("DELETE")

	// Category-Field assignment
	router.HandleFunc("/categories/{categoryId}/fields",
		h.HandleRequest[h.GetCategoryFieldsRequest](ch.GetCategoryFields)).Methods("GET")
	router.HandleFunc("/categories/{categoryId}/fields/{fieldId}",
		h.HandleRequest[h.AssignFieldToCategoryRequest](ch.AssignFieldToCategory)).Methods("POST")

	// Field Type routes
	router.HandleFunc("/field-types", h.HandleRequest[struct{}](ch.ListFieldTypes)).Methods("GET")
	router.HandleFunc("/field-types", h.HandleRequest[h.CreateFieldTypeRequest](ch.CreateFieldType)).Methods("POST")
	router.HandleFunc("/field-types/{id}", h.HandleRequest[struct{}](ch.GetFieldType)).Methods("GET")
	router.HandleFunc("/field-types/{id}", h.HandleRequest[h.UpdateFieldTypeRequest](ch.UpdateFieldType)).Methods("PUT")
	router.HandleFunc("/field-types/{id}", h.HandleRequest[struct{}](ch.DeleteFieldType)).Methods("DELETE")

	// Field Type Discriminator routes
	router.HandleFunc("/field-type-discriminators", h.HandleRequest[struct{}](ch.ListFieldTypeDiscriminators)).Methods("GET")
	router.HandleFunc("/field-type-discriminators", h.HandleRequest[h.CreateFieldTypeDiscriminatorRequest](ch.CreateFieldTypeDiscriminator)).Methods("POST")
	router.HandleFunc("/field-type-discriminators/{id}", h.HandleRequest[struct{}](ch.GetFieldTypeDiscriminator)).Methods("GET")
	router.HandleFunc("/field-type-discriminators/{id}", h.HandleRequest[h.UpdateFieldTypeDiscriminatorRequest](ch.UpdateFieldTypeDiscriminator)).Methods("PUT")
	router.HandleFunc("/field-type-discriminators/{id}", h.HandleRequest[struct{}](ch.DeleteFieldTypeDiscriminator)).Methods("DELETE")
}
