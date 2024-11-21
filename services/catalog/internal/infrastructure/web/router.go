package web

import (
	"localloop/libs/pkg/web"
	bh "localloop/libs/pkg/web/handler"
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
	router.HandleFunc("/categories", bh.HandleRequest(ch.Category.List)).Methods("GET")
	router.HandleFunc("/categories", bh.HandleRequest(ch.Category.Create)).Methods("POST")
	router.HandleFunc("/categories/{id}", bh.HandleRequest(ch.Category.Get)).Methods("GET")
	router.HandleFunc("/categories/{id}", bh.HandleRequest(ch.Category.Update)).Methods("PUT")
	router.HandleFunc("/categories/{id}", bh.HandleRequest(ch.Category.Delete)).Methods("DELETE")

	// Field routes
	router.HandleFunc("/fields", bh.HandleRequest(ch.Field.List)).Methods("GET")
	router.HandleFunc("/fields", bh.HandleRequest(ch.Field.Create)).Methods("POST")
	router.HandleFunc("/fields/{id}", bh.HandleRequest(ch.Field.Get)).Methods("GET")
	router.HandleFunc("/fields/{id}", bh.HandleRequest(ch.Field.Update)).Methods("PUT")
	router.HandleFunc("/fields/{id}", bh.HandleRequest(ch.Field.Delete)).Methods("DELETE")

	// Category-Field assignment
	router.HandleFunc("/categories/{categoryId}/fields",
		bh.HandleRequest(ch.GetCategoryFields)).Methods("GET")
	router.HandleFunc("/categories/{categoryId}/fields/{fieldId}",
		bh.HandleRequest(ch.AssignFieldToCategory)).Methods("POST")

	// Field Type routes
	router.HandleFunc("/field-types", bh.HandleRequest(ch.FieldType.List)).Methods("GET")
	router.HandleFunc("/field-types", bh.HandleRequest(ch.FieldType.Create)).Methods("POST")
	router.HandleFunc("/field-types/{id}", bh.HandleRequest(ch.FieldType.Get)).Methods("GET")
	router.HandleFunc("/field-types/{id}", bh.HandleRequest(ch.FieldType.Update)).Methods("PUT")
	router.HandleFunc("/field-types/{id}", bh.HandleRequest(ch.FieldType.Delete)).Methods("DELETE")

	// Field Type Discriminator routes
	router.HandleFunc("/field-type-discriminators", bh.HandleRequest(ch.Discriminator.List)).Methods("GET")
	router.HandleFunc("/field-type-discriminators", bh.HandleRequest(ch.Discriminator.Create)).Methods("POST")
	router.HandleFunc("/field-type-discriminators/{id}", bh.HandleRequest(ch.Discriminator.Get)).Methods("GET")
	router.HandleFunc("/field-type-discriminators/{id}", bh.HandleRequest(ch.Discriminator.Update)).Methods("PUT")
	router.HandleFunc("/field-type-discriminators/{id}", bh.HandleRequest(ch.Discriminator.Delete)).Methods("DELETE")
}
